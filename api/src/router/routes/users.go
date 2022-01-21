package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URL:                 "/users",
		Method:              http.MethodPost,
		Function:            controllers.CreateUser,
		RequireAuthenticate: false,
	},
	{
		URL:                 "/users",
		Method:              http.MethodGet,
		Function:            controllers.GetAllUsers,
		RequireAuthenticate: false,
	},

	{
		URL:                 "/users/{id}",
		Method:              http.MethodGet,
		Function:            controllers.GetUser,
		RequireAuthenticate: false,
	},

	{
		URL:                 "/users/{id}",
		Method:              http.MethodPut,
		Function:            controllers.UpdateUser,
		RequireAuthenticate: false,
	},

	{
		URL:                 "/users/{id}",
		Method:              http.MethodDelete,
		Function:            controllers.DeleteUser,
		RequireAuthenticate: false,
	},
}
