package main

import "net/http"

// Route is a struct used to differentiate between different
// requests that can be made against the library server.
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

// Routes that map destination URLs to functions
// in which book operations are performed.
var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"LookupBook",
		"GET",
		"/lookup/{bookId}",
		LookupBook,
	},
	Route{
		"AddBook",
		"POST",
		"/add",
		AddBook,
	},
	Route{
		"RemoveBook",
		"POST",
		"/remove/{bookId}",
		RemoveBook,
	},
}
