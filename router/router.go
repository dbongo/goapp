package router

import (
	"github.com/dbongo/hackapp/handler"
	ha "github.com/dbongo/hackapp/middleware"
	"github.com/dbongo/hackapp/token"

	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"
)

var (
	app *web.Mux
	api *web.Mux
)

func init() {
	// initialize app router / middleware
	app = web.New()
	app.Use(middleware.RequestID)
	app.Use(middleware.AutomaticOptions)
	app.Use(ha.RequestID)
	app.Use(ha.HTTPLogger)
	app.Use(ha.Recovery)
	app.Use(ha.SetHeaders)
	app.Post("/login", handler.Login)
	app.Post("/register", handler.Register)

	// initialize api router / token validation middleware
	api = web.New()
	api.Use(token.Validation)
	api.Get("/api/hello", handler.HelloWorld)
	api.Get("/api/hello/:name", handler.HelloName)
}

// Init returns the initialized router
func Init() *web.Mux {
	app.Handle("/api/*", api)
	return app
}
