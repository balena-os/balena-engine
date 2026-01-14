package network // import "github.com/docker/docker/api/server/router/network"

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/docker/docker/api/server/httputils"
	"github.com/docker/docker/api/types/backend"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/api/types/versions"
	"github.com/docker/docker/errdefs"
	"github.com/docker/docker/libnetwork"
	"github.com/pkg/errors"
)

func (n *networkRouter) getNetworksList(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	if err := httputils.ParseForm(r); err != nil {
		return err
	}

	filter, err := filters.FromJSON(r.Form.Get("filters"))
	if err != nil {
		return err
	}

	if err := network.ValidateFilters(filter); err != nil {
		return err
	}

	var list []network.Summary

	localNetworks, err := n.backend.GetNetworks(filter, backend.NetworkListConfig{Detailed: versions.LessThan(httputils.VersionFromContext(ctx), "1.28")})
	if err != nil {
		return err
	}

	list = localNetworks
	if list == nil {
		list = []network.Summary{}
	}

	return httputils.WriteJSON(w, http.StatusOK, list)
}

type invalidRequestError struct {
	cause error
}

func (e invalidRequestError) Error() string {
	return e.cause.Error()
}

func (e invalidRequestError) InvalidParameter() {}

type ambiguousResultsError string

func (e ambiguousResultsError) Error() string {
	return "network " + string(e) + " is ambiguous"
}

func (ambiguousResultsError) InvalidParameter() {}

func (n *networkRouter) getNetwork(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	if err := httputils.ParseForm(r); err != nil {
		return err
	}

	term := vars["id"]
	var (
		verbose bool
		err     error
	)
	if v := r.URL.Query().Get("verbose"); v != "" {
		if verbose, err = strconv.ParseBool(v); err != nil {
			return errors.Wrapf(invalidRequestError{err}, "invalid value for verbose: %s", v)
		}
	}
	networkScope := r.URL.Query().Get("scope")

	// In case multiple networks have duplicate names, return error.
	// TODO (yongtang): should we wrap with version here for backward compatibility?

	// First find based on full ID, return immediately once one is found.
	// For full name and partial ID, save the result first, and process later
	// in case multiple records was found based on the same term
	listByFullName := map[string]network.Inspect{}
	listByPartialID := map[string]network.Inspect{}

	// TODO(@cpuguy83): All this logic for figuring out which network to return does not belong here
	// Instead there should be a backend function to just get one network.
	filter := filters.NewArgs(filters.Arg("idOrName", term))
	if networkScope != "" {
		filter.Add("scope", networkScope)
	}
	networks, _ := n.backend.GetNetworks(filter, backend.NetworkListConfig{Detailed: true, Verbose: verbose})
	for _, nw := range networks {
		if nw.ID == term {
			return httputils.WriteJSON(w, http.StatusOK, nw)
		}
		if nw.Name == term {
			// No need to check the ID collision here as we are still in
			// local scope and the network ID is unique in this scope.
			listByFullName[nw.ID] = nw
		}
		if strings.HasPrefix(nw.ID, term) {
			// No need to check the ID collision here as we are still in
			// local scope and the network ID is unique in this scope.
			listByPartialID[nw.ID] = nw
		}
	}

	// Find based on full name, returns true only if no duplicates
	if len(listByFullName) == 1 {
		for _, v := range listByFullName {
			return httputils.WriteJSON(w, http.StatusOK, v)
		}
	}
	if len(listByFullName) > 1 {
		return errors.Wrapf(ambiguousResultsError(term), "%d matches found based on name", len(listByFullName))
	}

	// Find based on partial ID, returns true only if no duplicates
	if len(listByPartialID) == 1 {
		for _, v := range listByPartialID {
			return httputils.WriteJSON(w, http.StatusOK, v)
		}
	}
	if len(listByPartialID) > 1 {
		return errors.Wrapf(ambiguousResultsError(term), "%d matches found based on ID prefix", len(listByPartialID))
	}

	return libnetwork.ErrNoSuchNetwork(term)
}

func (n *networkRouter) postNetworkCreate(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	if err := httputils.ParseForm(r); err != nil {
		return err
	}

	var create network.CreateRequest
	if err := httputils.ReadJSON(r, &create); err != nil {
		return err
	}

	nw, err := n.backend.CreateNetwork(create)
	if err != nil {
		return err
	}

	return httputils.WriteJSON(w, http.StatusCreated, nw)
}

func (n *networkRouter) postNetworkConnect(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	if err := httputils.ParseForm(r); err != nil {
		return err
	}

	var connect network.ConnectOptions
	if err := httputils.ReadJSON(r, &connect); err != nil {
		return err
	}

	// Unlike other operations, we does not check ambiguity of the name/ID here.
	// Inside daemon `ConnectContainerToNetwork` does the ambiguity check anyway.
	// Therefore, passing the name to daemon would be enough.
	return n.backend.ConnectContainerToNetwork(ctx, connect.Container, vars["id"], connect.EndpointConfig)
}

func (n *networkRouter) postNetworkDisconnect(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	if err := httputils.ParseForm(r); err != nil {
		return err
	}

	var disconnect network.DisconnectOptions
	if err := httputils.ReadJSON(r, &disconnect); err != nil {
		return err
	}

	return n.backend.DisconnectContainerFromNetwork(disconnect.Container, vars["id"], disconnect.Force)
}

func (n *networkRouter) deleteNetwork(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	if err := httputils.ParseForm(r); err != nil {
		return err
	}

	nw, err := n.findUniqueNetwork(vars["id"])
	if err != nil {
		return err
	}
	if err := n.backend.DeleteNetwork(nw.ID); err != nil {
		return err
	}
	w.WriteHeader(http.StatusNoContent)
	return nil
}

func (n *networkRouter) postNetworksPrune(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	if err := httputils.ParseForm(r); err != nil {
		return err
	}

	pruneFilters, err := filters.FromJSON(r.Form.Get("filters"))
	if err != nil {
		return err
	}

	pruneReport, err := n.backend.NetworksPrune(ctx, pruneFilters)
	if err != nil {
		return err
	}
	return httputils.WriteJSON(w, http.StatusOK, pruneReport)
}

// findUniqueNetwork will search network.
// NOTE: This findUniqueNetwork is different from FindNetwork in the daemon.
// In case multiple networks have duplicate names, return error.
// First find based on full ID, return immediately once one is found.
// For full name and partial ID, save the result first, and process later
// in case multiple records was found based on the same term
// TODO (yongtang): should we wrap with version here for backward compatibility?
func (n *networkRouter) findUniqueNetwork(term string) (network.Inspect, error) {
	listByFullName := map[string]network.Inspect{}
	listByPartialID := map[string]network.Inspect{}

	filter := filters.NewArgs(filters.Arg("idOrName", term))
	networks, _ := n.backend.GetNetworks(filter, backend.NetworkListConfig{Detailed: true})
	for _, nw := range networks {
		if nw.ID == term {
			return nw, nil
		}
		if nw.Name == term && !nw.Ingress {
			// No need to check the ID collision here as we are still in
			// local scope and the network ID is unique in this scope.
			listByFullName[nw.ID] = nw
		}
		if strings.HasPrefix(nw.ID, term) {
			// No need to check the ID collision here as we are still in
			// local scope and the network ID is unique in this scope.
			listByPartialID[nw.ID] = nw
		}
	}

	// Find based on full name, returns true only if no duplicates
	if len(listByFullName) == 1 {
		for _, v := range listByFullName {
			return v, nil
		}
	}
	if len(listByFullName) > 1 {
		return network.Inspect{}, errdefs.InvalidParameter(errors.Errorf("network %s is ambiguous (%d matches found based on name)", term, len(listByFullName)))
	}

	// Find based on partial ID, returns true only if no duplicates
	if len(listByPartialID) == 1 {
		for _, v := range listByPartialID {
			return v, nil
		}
	}
	if len(listByPartialID) > 1 {
		return network.Inspect{}, errdefs.InvalidParameter(errors.Errorf("network %s is ambiguous (%d matches found based on ID prefix)", term, len(listByPartialID)))
	}

	return network.Inspect{}, errdefs.NotFound(libnetwork.ErrNoSuchNetwork(term))
}
