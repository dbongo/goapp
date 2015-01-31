package router

import (
	"github.com/dbongo/hackapp/handler"
	"github.com/dbongo/hackapp/middleware"

	"github.com/zenazn/goji/web"
)

// New ...
func New() *web.Mux {
	mux := web.New()
	mux.Post("/api/login", handler.LoginUser)
	mux.Post("/api/register", handler.RegisterUser)

	hello := web.New()
	hello.Get("/api/hello", handler.HelloWorld)
	hello.Get("/api/hello/:name", handler.HelloName)
	mux.Handle("/api/hello", hello)
	mux.Handle("/api/hello/*", hello)

	user := web.New()
	user.Use(middleware.RequireUser)
	user.Get("/api/user", handler.GetCurrentUser)
	user.Put("/api/user", handler.PutUser)
	mux.Handle("/api/user", user)

	users := web.New()
	users.Use(middleware.RequireAdminUser)
	users.Get("/api/users", handler.GetUserList)
	mux.Handle("/api/users", users)

	return mux
}
