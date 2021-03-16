package go_simple_http

import (
	"encoding/json"
	"net/http"
)

// Response ...
type Response struct {
	response http.ResponseWriter
}

// NewResponse ...
func NewResponse(res http.ResponseWriter) Response {
	return Response{
		response: res,
	}
}

// Status
func (res *Response) Status(code int) *Response {
	return res
}

// Send ...
func (res *Response) Send(data string) {
	res.response.Write([]byte(data))
}

// Json ...
func (res *Response) Json(data interface{}) {
	buffer, _ := json.Marshal(data)
	res.response.Header().Add("Content-type", "application/json")
	res.response.Write(buffer)
}
