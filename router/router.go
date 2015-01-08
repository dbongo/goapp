package router

import (
	"github.com/dbongo/hackapp/handler"
	"github.com/dbongo/hackapp/token"

	"github.com/zenazn/goji/web"
)

// New returns our api router
func New() *web.Mux {
	mux := web.New()
	mux.Post("/api/auth/login", handler.Login)
	mux.Post("/api/auth/register", handler.Register)

	api := web.New()
	api.Use(token.Validation)
	api.Get("/api/hello/world", handler.HelloWorld)
	api.Get("/api/hello/:name", handler.HelloName)

	mux.Handle("/api/hello/*", api)
	return mux
}
