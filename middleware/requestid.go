package middleware

import (
	"net/http"

	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"
)

// RequestID ...
func RequestID(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		reqID := middleware.GetReqID(*c)
		w.Header().Set("Request-Id", reqID)
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
