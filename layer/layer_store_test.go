package layer

import (
	"reflect"
	"sort"
	"strings"
	"testing"

	"github.com/docker/docker/daemon/graphdriver"
	"gotest.tools/v3/skip"
)

func TestLayerStore_prune(t *testing.T) {
	s, _, cleanup := newTestStore(t)
	defer cleanup()

	lStore, ok := s.(*layerStore)
	if !ok {
		t.Fatalf("Unexpected store implementation %s", s)
	}

	// Start new transaction with cacheID that is not committed or canceled.
	_, err := lStore.store.StartTransaction("test-prune")
	if err != nil {
		t.Fatal(err)
	}

	txData, err := lStore.store.ListExistingTransactions()
	if err != nil {
		t.Fatal(err)
	}
	cacheIDs := lStore.prune(txData)
	if len(cacheIDs) != 1 {
		t.Errorf("1 directory was expected to be removed, but got %s", cacheIDs)
	}
}

func createGraphDriverLayer(t *testing.T, lStore *layerStore, name string) {
	err := lStore.Driver().Create(name, "", nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestLayerStore_unreferencedDriverLayers(t *testing.T) {
	s, _, cleanup := newTestStore(t)
	defer cleanup()
	skip.If(t, s.DriverName() != "aufs" && s.DriverName() != "overlay2" && s.DriverName() != "vfs",
		"aufs, overlay2, and vfs drivers are supported only, got %s", s.DriverName())

	lStore, ok := s.(*layerStore)
	if !ok {
		t.Fatalf("Unexpected store implementation %s", s)
	}

	// Leaked layers carry the IDs this store generates.
	leaked1 := strings.Repeat("a", 64)
	leaked2 := strings.Repeat("b", 64) + "-init"
	createGraphDriverLayer(t, lStore, leaked1)
	createGraphDriverLayer(t, lStore, leaked2)
	// A layer created by another graphdriver consumer, such as the BuildKit
	// snapshotter, uses its own ID scheme and must not be reclaimed.
	foreign := "vr7rqg7xk2f4o1r4o4x2f5a3z"
	createGraphDriverLayer(t, lStore, foreign)

	_, err := lStore.CreateRWLayer("test-container-1", "", &CreateRWLayerOpts{
		InitFunc: func(root string) error {
			return nil
		}},
	)
	if err != nil {
		t.Fatal(err)
	}
	// Create RW (container) layer without init func - this will not create an -init layer.
	_, err = lStore.CreateRWLayer("test-container-2", "", nil)
	if err != nil {
		t.Fatal(err)
	}

	// 2 leaked layers + 1 foreign layer + 2 containers + init layer for one of container.
	if ids, _ := lStore.driver.(graphdriver.InspectableDriver).List(); len(ids) != 6 {
		t.Fatalf("Unexpected graphdriver storage state: %s", ids)
	}

	txData, err := lStore.store.ListExistingTransactions()
	if err != nil {
		t.Fatal(err)
	}
	if len(txData) != 0 {
		t.Errorf("No existing transactions were expected, but got %s", txData)
	}

	ids, err := lStore.findUnreferencedDriverLayers()
	if err != nil {
		t.Fatal(err)
	}
	sort.Strings(ids)
	if !reflect.DeepEqual(ids, []string{leaked1, leaked2}) {
		t.Errorf("Leaked layers are incorrectly detected: %s", ids)
	}

	n := lStore.deleteUnreferencedDriverLayers(ids)
	if n != 2 {
		t.Errorf("2 layers were expected to be deleted, but got %d", n)
	}

	for _, cacheID := range ids {
		if lStore.Driver().Exists(cacheID) {
			t.Errorf("Layer %s was not deleted", cacheID)
		}
	}
	if !lStore.Driver().Exists(foreign) {
		t.Errorf("Foreign layer %s was deleted", foreign)
	}
}
