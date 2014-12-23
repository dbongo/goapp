package middleware

import (
	"net/http"
	"runtime/debug"

	"github.com/dbongo/goapp/logger"
)

// Recovery is a middleware that recovers from panics, logs the panic (and a backtrace), and returns a HTTP 500 (Internal Server Error) status if possible.
func Recovery(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if e := recover(); e != nil {
				logger.Error.Printf("Recovering from error '%s'", e)
				logger.Trace.Printf(string(debug.Stack()))
				http.Error(w, http.StatusText(500), 500)
				return
			}
		}()
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
