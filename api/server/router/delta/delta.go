package delta

import "github.com/docker/docker/api/server/router"

// deltaRouter is a router to talk with the deltas controller
type deltaRouter struct {
	backend Backend
	routes  []router.Route
}

// NewRouter initializes a new delta router
func NewRouter(b Backend) router.Router {
	r := &deltaRouter{
		backend: b,
	}
	r.initRoutes()
	return r
}

// Routes returns the available routes to the deltas controller
func (r *deltaRouter) Routes() []router.Route {
	return r.routes
}

func (r *deltaRouter) initRoutes() {
	r.routes = []router.Route{
		// POST
		router.NewPostRoute("/deltas/create", r.postDeltasCreate),
	}
}
