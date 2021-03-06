package driver

import (
	"github.com/gorilla/mux"
	"net/http"
)

// NewRouter constructs new mux.Router.
func NewRouter(h Handler) *mux.Router {
	r := mux.NewRouter()

	setupRoutes(r, h)

	return r
}

func setupRoutes(r *mux.Router, h Handler) {
	r.Methods(http.MethodGet).Path("/json-p-target").HandlerFunc(h.JsonPTarget())
	r.Methods(http.MethodGet).Path("/cors-target").HandlerFunc(h.CorsTarget())
}
