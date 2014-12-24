package middleware

import (
	"net/http"

	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"
)

// RequestID ...
func RequestID(c *web.C, h http.Handler) http.Handler {
	fn := func(rw http.ResponseWriter, req *http.Request) {
		reqID := middleware.GetReqID(*c)
		rw.Header().Set("Request-Id", reqID)
		h.ServeHTTP(rw, req)
	}
	return http.HandlerFunc(fn)
}
