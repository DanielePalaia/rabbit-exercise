package main

import (
	"net/http"
	"rabbit-exercise/controllers"

	"github.com/gorilla/mux"
)

// Route struct contains route configuration
type Route struct {
	name    string
	method  string
	pattern string
	handler http.Handler
}

// Routes ia a vector of Route
type Routes []Route

// NewRouter returns a router to manage different routes
func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.method).
			Path(route.pattern).
			Name(route.name).
			Handler(route.handler)
	}

	return router
}

var routes = Routes{
	// Publish to an exchange
	Route{
		"getMetrics",
		"GET",
		"/publish",
		http.HandlerFunc(controllers.HandlePublish),
	},
	// Consume from a queue
	Route{
		"getMetrics",
		"GET",
		"/consume",
		http.HandlerFunc(controllers.HandleConsume),
	},
}
