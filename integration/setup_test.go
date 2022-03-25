package integration

import (
	"context"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func bundle(bundle string) error {
	log.Printf("---> Making bundle: %s", bundle)

	ctx := context.Background()
	// TODO: this assumes we're in a transformer
	root := filepath.Dir(os.Getenv("INPUT"))
	bundle = filepath.Join(root, "hack", "make", bundle)

	cmd := exec.CommandContext(ctx, "bash", "-c", "source "+bundle)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func TestMain(m *testing.M) {
	if err := os.MkdirAll("bundles", os.FileMode(0644)); err != nil {
		log.Fatal(err)
	}
	if err := bundle(".integration-daemon-start"); err != nil {
		log.Fatal(err)
	}
	if err := bundle(".integration-daemon-start"); err != nil {
		log.Fatal(err)
	}
}
