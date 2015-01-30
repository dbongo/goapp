package middleware

import (
	"net/http"

	"github.com/zenazn/goji/web"
)

// SetHeaders ...
func SetHeaders(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		reqID := GetReqID(*c)
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Set("Request-Id", reqID)
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
