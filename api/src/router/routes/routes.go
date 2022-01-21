package routes

import "net/http"

//Route represent all routes from api
type Route struct {
	URL                 string
	Method              string
	Function            func(http.ResponseWriter, *http.Request)
	RequireAuthenticate bool
}
