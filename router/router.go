package router

import (
	"github.com/dbongo/goapp/handler"
	"github.com/zenazn/goji/web"
)

// New ...
func New() *web.Mux {
	mux := web.New()
	mux.Get("/api/hello", handler.HelloWorld)
	mux.Get("/api/hello/:name", handler.HelloName)
	// user := web.New()
	// user.Get("/api/user", handler.CurrentUser)
	// mux.Handle("/api/user/*", user)
	return mux
}
