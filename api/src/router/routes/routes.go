package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//Route represent all routes from api
type Route struct {
	URL                 string
	Method              string
	Function            func(http.ResponseWriter, *http.Request)
	RequireAuthenticate bool
}

//Configure all routes for router
func Configure(r mux.Router) *mux.Router {
	routes := userRoutes

	for _, route := range routes {
		r.HandleFunc(route.URL, route.Function).Methods(route.Method)
	}

	return &r
}
