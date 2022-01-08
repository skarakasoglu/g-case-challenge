package inmem

import "net/http"

type Controller struct{}

func (c Controller) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
	case http.MethodGet:
	default:

	}
}
