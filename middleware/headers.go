package middleware

import (
	"net/http"

	"github.com/zenazn/goji/web"
)

// SetHeaders applies default headers to each request
func SetHeaders(c *web.C, h http.Handler) http.Handler {
	fn := func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Add("Access-Control-Allow-Origin", "*")
		rw.Header().Add("Content-Type", "application/json")
		h.ServeHTTP(rw, req)
	}
	return http.HandlerFunc(fn)
}
