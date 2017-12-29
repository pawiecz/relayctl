package main

import "github.com/julienschmidt/httprouter"

// Route represents single API route and stores its handler function.
type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc httprouter.Handle
}

// Routes groups API routes for passing them between functions.
type Routes []Route

// AllRoutes defines all supported API routes.
func AllRoutes() Routes {
	routes := Routes{
		Route{"Index", "GET", "/", Index},
		Route{"RelayIndex", "GET", "/relays", RelayIndex},
		Route{"RelayShow", "GET", "/relays/:id", RelayShow},
	}
	return routes
}

// NewRouter reads from the routes slice to set the values for httprouter.Handle.
func NewRouter(routes Routes) *httprouter.Router {
	router := httprouter.New()
	for _, route := range routes {
		router.Handle(route.Method, route.Path, route.HandlerFunc)
	}
	return router
}
