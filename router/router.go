package router

import (
	"github.com/dbongo/hackapp/handler"
	"github.com/dbongo/hackapp/token"

	"github.com/zenazn/goji/web"
)

// New returns our api router
func New() *web.Mux {
	mux := web.New()
	mux.Post("/api/login", handler.Login)
	mux.Post("/api/register", handler.Register)

	hello := web.New()
	hello.Use(token.Validation)
	hello.Get("/api/hello", handler.HelloWorld)
	hello.Get("/api/hello/:name", handler.HelloName)
	mux.Handle("/api/hello", hello)
	mux.Handle("/api/hello/*", hello)
	return mux
}
