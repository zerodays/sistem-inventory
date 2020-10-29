package router

import "net/http"

type Route struct {
	// Name of the route, used for creating URL from name.
	Name string

	// Path of the route.
	Path string

	// Can this route be accessed only by authorized users
	AuthorizedOnly bool

	// Will the route use a custom Content-Type header. Default is application/json.
	CustomContentType bool

	// Handlers for different http methods.
	GET, POST, PUT, DELETE http.Handler
}
