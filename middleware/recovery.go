package middleware

import (
	"net/http"
	"runtime/debug"

	"github.com/dbongo/hackapp/logger"
)

// Recovery is a middleware handler that recovers from panics, logs the panic
// (and a backtrace), and returns a HTTP 500 (Internal Server Error) status if possible.
func Recovery(h http.Handler) http.Handler {
	fn := func(rw http.ResponseWriter, req *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				logger.Error.Printf("Recovering from error '%s'", err)
				logger.Trace.Printf(string(debug.Stack()))
				http.Error(rw, http.StatusText(500), 500)
				return
			}
		}()
		h.ServeHTTP(rw, req)
	}
	return http.HandlerFunc(fn)
}
