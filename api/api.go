// Package api implements an API server.
package api

import "net/http"

// Endpoint represents an API action which is served in
// a path and constructs an HTTP response by http.Handler
type Endpoint struct{
	Path string
	Handler http.Handler
}

// Start creates the API actions given as Endpoint.
// Then, serves the API on the specified address.
func Start(address string, endpoints ...Endpoint) error {
	for _, endpoint := range endpoints {
		http.Handle(endpoint.Path, endpoint.Handler)
	}

	err := http.ListenAndServe(address, nil)
	return err
}