package network // import "github.com/docker/docker/api/server/router/network"

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/docker/docker/api/server/httputils"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/api/types/versions"
	"github.com/docker/docker/errdefs"
	"github.com/docker/libnetwork"
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
		return errdefs.InvalidParameter(err)
	}

	var list []types.NetworkResource

	localNetworks, err := n.backend.GetNetworks(filter, types.NetworkListConfig{Detailed: versions.LessThan(httputils.VersionFromContext(ctx), "1.28")})
	if err != nil {
		return err
	}

	// Removed unnecessary filtering
	list = localNetworks

	if list == nil {
		list = []types.NetworkResource{}
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

type ambigousResultsError string

func (e ambigousResultsError) Error() string {
	return "network " + string(e) + " is ambiguous"
}

func (ambigousResultsError) InvalidParameter() {}

func nameConflict(name string) error {
	return errdefs.Conflict(libnetwork.NetworkNameError(name))
}

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
	scope := r.URL.Query().Get("scope")

	// In case multiple networks have duplicate names, return error.
	// TODO (yongtang): should we wrap with version here for backward compatibility?

	// First find based on full ID, return immediately once one is found.
	// If a network appears both in swarm and local, assume it is in local first

	// For full name and partial ID, save the result first, and process later
	// in case multiple records was found based on the same term
	listByFullName := map[string]types.NetworkResource{}
	listByPartialID := map[string]types.NetworkResource{}

	// TODO(@cpuguy83): All this logic for figuring out which network to return does not belong here
	// Instead there should be a backend function to just get one network.
	filter := filters.NewArgs(filters.Arg("idOrName", term))
	if scope != "" {
		filter.Add("scope", scope)
	}
	nw, _ := n.backend.GetNetworks(filter, types.NetworkListConfig{Detailed: true, Verbose: verbose})
	for _, network := range nw {
		if network.ID == term {
			return httputils.WriteJSON(w, http.StatusOK, network)
		}
		if network.Name == term {
			// No need to check the ID collision here as we are still in
			// local scope and the network ID is unique in this scope.
			listByFullName[network.ID] = network
		}
		if strings.HasPrefix(network.ID, term) {
			// No need to check the ID collision here as we are still in
			// local scope and the network ID is unique in this scope.
			listByPartialID[network.ID] = network
		}
	}

	// Find based on full name, returns true only if no duplicates
	if len(listByFullName) == 1 {
		for _, v := range listByFullName {
			return httputils.WriteJSON(w, http.StatusOK, v)
		}
	}
	if len(listByFullName) > 1 {
		return errors.Wrapf(ambigousResultsError(term), "%d matches found based on name", len(listByFullName))
	}

	// Find based on partial ID, returns true only if no duplicates
	if len(listByPartialID) == 1 {
		for _, v := range listByPartialID {
			return httputils.WriteJSON(w, http.StatusOK, v)
		}
	}
	if len(listByPartialID) > 1 {
		return errors.Wrapf(ambigousResultsError(term), "%d matches found based on ID prefix", len(listByPartialID))
	}

	return libnetwork.ErrNoSuchNetwork(term)
}

func (n *networkRouter) postNetworkCreate(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	var create types.NetworkCreateRequest

	if err := httputils.ParseForm(r); err != nil {
		return err
	}

	if err := httputils.CheckForJSON(r); err != nil {
		return err
	}

	if err := json.NewDecoder(r.Body).Decode(&create); err != nil {
		if err == io.EOF {
			return errdefs.InvalidParameter(errors.New("got EOF while reading request body"))
		}
		return errdefs.InvalidParameter(err)
	}

	nw, err := n.backend.CreateNetwork(create)
	if err != nil {
		if _, ok := err.(libnetwork.NetworkNameError); ok {
			// check if user defined CheckDuplicate, if set true, return err
			if create.CheckDuplicate {
				return nameConflict(create.Name)
			}
		}
		return err
	}

	return httputils.WriteJSON(w, http.StatusCreated, nw)
}

func (n *networkRouter) postNetworkConnect(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	var connect types.NetworkConnect
	if err := httputils.ParseForm(r); err != nil {
		return err
	}

	if err := httputils.CheckForJSON(r); err != nil {
		return err
	}

	if err := json.NewDecoder(r.Body).Decode(&connect); err != nil {
		if err == io.EOF {
			return errdefs.InvalidParameter(errors.New("got EOF while reading request body"))
		}
		return errdefs.InvalidParameter(err)
	}

	// Unlike other operations, we does not check ambiguity of the name/ID here.
	// The reason is that, In case of attachable network in swarm scope, the actual local network
	// may not be available at the time. At the same time, inside daemon `ConnectContainerToNetwork`
	// does the ambiguity check anyway. Therefore, passing the name to daemon would be enough.
	return n.backend.ConnectContainerToNetwork(connect.Container, vars["id"], connect.EndpointConfig)
}

func (n *networkRouter) postNetworkDisconnect(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	var disconnect types.NetworkDisconnect
	if err := httputils.ParseForm(r); err != nil {
		return err
	}

	if err := httputils.CheckForJSON(r); err != nil {
		return err
	}

	if err := json.NewDecoder(r.Body).Decode(&disconnect); err != nil {
		if err == io.EOF {
			return errdefs.InvalidParameter(errors.New("got EOF while reading request body"))
		}
		return errdefs.InvalidParameter(err)
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

// findUniqueNetwork will search network across different scopes (both local and swarm).
// NOTE: This findUniqueNetwork is different from FindNetwork in the daemon.
// In case multiple networks have duplicate names, return error.
// First find based on full ID, return immediately once one is found.
// If a network appears both in swarm and local, assume it is in local first
// For full name and partial ID, save the result first, and process later
// in case multiple records was found based on the same term
// TODO (yongtang): should we wrap with version here for backward compatibility?
//
func (n *networkRouter) findUniqueNetwork(term string) (types.NetworkResource, error) {
	listByFullName := map[string]types.NetworkResource{}
	listByPartialID := map[string]types.NetworkResource{}

	filter := filters.NewArgs(filters.Arg("idOrName", term))
	nw, _ := n.backend.GetNetworks(filter, types.NetworkListConfig{Detailed: true})
	for _, network := range nw {
		if network.ID == term {
			return network, nil
		}
		if network.Name == term && !network.Ingress {
			// No need to check the ID collision here as we are still in
			// local scope and the network ID is unique in this scope.
			listByFullName[network.ID] = network
		}
		if strings.HasPrefix(network.ID, term) {
			// No need to check the ID collision here as we are still in
			// local scope and the network ID is unique in this scope.
			listByPartialID[network.ID] = network
		}
	}

	// Find based on full name, returns true only if no duplicates
	if len(listByFullName) == 1 {
		for _, v := range listByFullName {
			return v, nil
		}
	}
	if len(listByFullName) > 1 {
		return types.NetworkResource{}, errdefs.InvalidParameter(errors.Errorf("network %s is ambiguous (%d matches found based on name)", term, len(listByFullName)))
	}

	// Find based on partial ID, returns true only if no duplicates
	if len(listByPartialID) == 1 {
		for _, v := range listByPartialID {
			return v, nil
		}
	}
	if len(listByPartialID) > 1 {
		return types.NetworkResource{}, errdefs.InvalidParameter(errors.Errorf("network %s is ambiguous (%d matches found based on ID prefix)", term, len(listByPartialID)))
	}

	return types.NetworkResource{}, errdefs.NotFound(libnetwork.ErrNoSuchNetwork(term))
}
