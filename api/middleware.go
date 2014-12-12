package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dbongo/goapp/err"
	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"
)

// RequestIDHandler ...
func RequestIDHandler(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		reqID := middleware.GetReqID(*c)
		w.Header().Set("Request-Id", reqID)
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

// NotFoundHandler ...
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	notFound := &err.HTTPErr{
		StatusCode: http.StatusNotFound,
		Message:    "The resource you are looking for was not found.",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(notFound.StatusCode)
	body, _ := json.Marshal(notFound)
	fmt.Fprint(w, string(body))
	return
}
