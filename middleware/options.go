package middleware

import (
	"net/http"

	"github.com/zenazn/goji/web"
)

// Options automatically return an appropriate "Allow" header when the
// request method is OPTIONS and the request would have otherwise been 404'd.
func Options(c *web.C, h http.Handler) http.Handler {
	fn := func(rw http.ResponseWriter, req *http.Request) {
		if req.Method == "OPTIONS" {
			rw.Header().Set("Access-Control-Allow-Origin", "*")
			rw.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
			rw.Header().Set("Access-Control-Allow-Headers", "Authorization")
			rw.Header().Set("Allow", "HEAD,GET,POST,PUT,DELETE,OPTIONS")
			rw.WriteHeader(http.StatusOK)
			return
		}
		h.ServeHTTP(rw, req)
	}
	return http.HandlerFunc(fn)
}
