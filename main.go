package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/goji/context"

	"github.com/dbongo/hackapp/middleware"
	"github.com/dbongo/hackapp/router"
)

var (
	port = flag.String("p", ":3000", "server port")
)

func main() {
	flag.Parse()

	// create the router and add middleware
	mux := router.New()
	mux.Use(middleware.RequestID)
	mux.Use(middleware.Options)
	mux.Use(context.Middleware)
	mux.Use(middleware.SetHeaders)
	mux.Use(middleware.HTTPLogger)
	mux.Use(middleware.SetUser)
	mux.Use(middleware.Recovery)
	http.Handle("/api/", mux)

	if err := http.ListenAndServe(*port, nil); err != nil {
		log.Fatal(err)
	}
}
