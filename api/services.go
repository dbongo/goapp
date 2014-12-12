package api

import (
	"net/http"

	"github.com/zenazn/goji/web"
)

// ServicesHandler ...
type ServicesHandler struct {
	ApiHandler
}

// Index ...
func (handler *ServicesHandler) Index(c *web.C, w http.ResponseWriter, r *http.Request) *HTTPResponse {
	response := &HTTPResponse{StatusCode: http.StatusOK, Payload: "Hello World"}
	return response
}
