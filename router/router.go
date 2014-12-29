package router

import (
	"github.com/dbongo/hackapp/handler"
	mw "github.com/dbongo/hackapp/middleware"
	"github.com/dbongo/hackapp/token"

	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"
)

var (
	app *web.Mux
	api *web.Mux
)

func init() {

	// initialize app router and core middleware
	app = web.New()
	app.Use(middleware.RequestID)
	app.Use(middleware.AutomaticOptions)
	app.Use(mw.RequestID)
	app.Use(mw.HTTPLogger)
	app.Use(mw.Recovery)
	app.Use(mw.SetHeaders)

	// attach app routes
	app.Post("/login", handler.Login)
	app.Post("/register", handler.Register)

	// initialize api router and token validation middleware
	api = web.New()
	api.Use(token.Validation)

	// attach api routes
	api.Get("/api/hello", handler.HelloWorld)
	api.Get("/api/hello/:name", handler.HelloName)
}

// New returns an instance of the app router
func New() *web.Mux {

	// attach api router to app
	app.Handle("/api/*", api)

	return app
}
