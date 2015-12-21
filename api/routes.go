package api

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"UserIndex",
		"GET",
		"/users",
		UserIndex,
	},
	Route{
		"UserCreate",
		"POST",
		"/users",
		UserCreate,
	},
	Route{
		"UserShow",
		"GET",
		"/users/{userId}",
		UserShow,
	},
}
