package client // import "github.com/docker/docker/client"

import (
	"fmt"
	"net/http"

	"github.com/docker/docker/api/types/versions"
	"github.com/docker/docker/errdefs"
	"github.com/pkg/errors"
)

// errConnectionFailed implements an error returned when connection failed.
type errConnectionFailed struct {
	host string
}

// Error returns a string representation of an errConnectionFailed
func (err errConnectionFailed) Error() string {
	if err.host == "" {
		return "Cannot connect to the balenaEngine daemon. Is the balenaEngine daemon running on this host?"
	}
	return fmt.Sprintf("Cannot connect to the balenaEngine daemon at %s. Is the balenaEngine daemon running?", err.host)
}

// IsErrConnectionFailed returns true if the error is caused by connection failed.
func IsErrConnectionFailed(err error) bool {
	return errors.As(err, &errConnectionFailed{})
}

// ErrorConnectionFailed returns an error with host in the error message when connection to docker daemon failed.
func ErrorConnectionFailed(host string) error {
	return errConnectionFailed{host: host}
}

// Deprecated: use the errdefs.NotFound() interface instead. Kept for backward compatibility
type notFound interface {
	error
	NotFound() bool
}

// IsErrNotFound returns true if the error is a NotFound error, which is returned
// by the API when some object is not found.
func IsErrNotFound(err error) bool {
	var e notFound
	if errors.As(err, &e) {
		return true
	}
	return errdefs.IsNotFound(err)
}

type objectNotFoundError struct {
	object string
	id     string
}

func (e objectNotFoundError) NotFound() {}

func (e objectNotFoundError) Error() string {
	return fmt.Sprintf("Error: No such %s: %s", e.object, e.id)
}

func wrapResponseError(err error, resp serverResponse, object, id string) error {
	switch {
	case err == nil:
		return nil
	case resp.statusCode == http.StatusNotFound:
		return objectNotFoundError{object: object, id: id}
	case resp.statusCode == http.StatusNotImplemented:
		return errdefs.NotImplemented(err)
	default:
		return err
	}
}

// unauthorizedError represents an authorization error in a remote registry.
type unauthorizedError struct {
	cause error
}

// Error returns a string representation of an unauthorizedError
func (u unauthorizedError) Error() string {
	return u.cause.Error()
}

// IsErrUnauthorized returns true if the error is caused
// when a remote registry authentication fails
func IsErrUnauthorized(err error) bool {
	if _, ok := err.(unauthorizedError); ok {
		return ok
	}
	return errdefs.IsUnauthorized(err)
}

type pluginPermissionDenied struct {
	name string
}

func (e pluginPermissionDenied) Error() string {
	return "Permission denied while installing plugin " + e.name
}

// IsErrPluginPermissionDenied returns true if the error is caused
// when a user denies a plugin's permissions
func IsErrPluginPermissionDenied(err error) bool {
	_, ok := err.(pluginPermissionDenied)
	return ok
}

type notImplementedError struct {
	message string
}

func (e notImplementedError) Error() string {
	return e.message
}

func (e notImplementedError) NotImplemented() bool {
	return true
}

// IsErrNotImplemented returns true if the error is a NotImplemented error.
// This is returned by the API when a requested feature has not been
// implemented.
func IsErrNotImplemented(err error) bool {
	if _, ok := err.(notImplementedError); ok {
		return ok
	}
	return errdefs.IsNotImplemented(err)
}

// NewVersionError returns an error if the APIVersion required
// if less than the current supported version
func (cli *Client) NewVersionError(APIrequired, feature string) error {
	if cli.version != "" && versions.LessThan(cli.version, APIrequired) {
		return fmt.Errorf("%q requires API version %s, but the Docker daemon API version is %s", feature, APIrequired, cli.version)
	}
	return nil
}
