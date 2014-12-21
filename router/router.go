package router

import (
	"github.com/dbongo/goapp/handler"
	"github.com/dbongo/goapp/middleware"
	"github.com/zenazn/goji/web"
)

// New returns api router
func New() *web.Mux {
	mux := web.New()

	mux.Use(middleware.Auth)
	mux.Get("/api/hello", handler.HelloWorld)
	mux.Get("/api/hello/:name", handler.HelloName)

	return mux
}
