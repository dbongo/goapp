package router

import (
	"github.com/dbongo/hackapp/handler"
	"github.com/zenazn/goji/web"
)

// New returns an instance of the api router
func New() *web.Mux {
	mux := web.New()
	mux.Get("/api/hello", handler.HelloWorld)
	mux.Get("/api/hello/:name", handler.HelloName)
	return mux
}
