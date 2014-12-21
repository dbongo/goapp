package main

import (
	"github.com/subosito/gotenv"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web/middleware"

	"github.com/dbongo/goapp/handler"
	"github.com/dbongo/goapp/logger"
	mw "github.com/dbongo/goapp/middleware"
	"github.com/dbongo/goapp/router"
)

func init() {
	gotenv.Load(".env")
}

func main() {
	logger.Info.Println("initializing server")

	goji.Abandon(middleware.Logger)
	goji.NotFound(mw.NotFound)

	goji.Use(mw.RequestID)
	goji.Use(middleware.EnvInit)
	goji.Use(middleware.Recoverer)
	goji.Use(mw.RequestLogger)

	goji.Post("/users", handler.Register)
	goji.Post("/login", handler.Login)

	mux := router.New()

	goji.Handle("/api/*", mux)
	goji.Serve()
}
