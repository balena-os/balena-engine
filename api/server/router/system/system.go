package system // import "github.com/docker/docker/api/server/router/system"

import (
	"github.com/docker/docker/api/server/router"
	"github.com/docker/docker/api/types"
	buildkit "github.com/docker/docker/builder/builder-next"
	"github.com/docker/docker/builder/fscache"
)

// systemRouter provides information about the Docker system overall.
// It gathers information about host, daemon and container events.
type systemRouter struct {
	backend        Backend
	routes         []router.Route
	fscache        *fscache.FSCache // legacy
	builder        *buildkit.Builder
	builderVersion types.BuilderVersion
}

// NewRouter initializes a new system router
func NewRouter(b Backend, fscache *fscache.FSCache, builder *buildkit.Builder, bv types.BuilderVersion) router.Router {
	r := &systemRouter{
		backend:        b,
		fscache:        fscache,
		builder:        builder,
		builderVersion: bv,
	}

	r.routes = []router.Route{
		router.NewOptionsRoute("/{anyroute:.*}", optionsHandler),
		router.NewGetRoute("/_ping", r.pingHandler),
		router.NewGetRoute("/events", r.getEvents, router.WithCancel),
		router.NewGetRoute("/info", r.getInfo),
		router.NewGetRoute("/version", r.getVersion),
		router.NewGetRoute("/system/df", r.getDiskUsage, router.WithCancel),
		router.NewPostRoute("/auth", r.postAuth),
	}

	return r
}

// Routes returns all the API routes dedicated to the docker system
func (s *systemRouter) Routes() []router.Route {
	return s.routes
}
