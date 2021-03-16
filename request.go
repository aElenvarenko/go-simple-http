package go_simple_http

import "net/http"

// Request ...
type Request struct {
	Params  map[string]interface{}
	request *http.Request
}

// NewRequest ...
func NewRequest(r *http.Request) Request {
	return Request{
		request: r,
	}
}
