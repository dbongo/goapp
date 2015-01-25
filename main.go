package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/zenazn/goji/web/middleware"

	mw "github.com/dbongo/hackapp/middleware"
	"github.com/dbongo/hackapp/router"
)

var (
	port = flag.String("p", ":3000", "server port")
)

func main() {
	flag.Parse()

	// create api router
	mux := router.New()

	// configure api middleware
	mux.Use(middleware.RequestID)
	mux.Use(mw.Options)
	mux.Use(mw.SetHeaders)
	mux.Use(mw.HTTPLogger)
	mux.Use(mw.Recovery)

	// set api handler
	http.Handle("/api/", mux)

	// start server on port 3000
	if err := http.ListenAndServe(*port, nil); err != nil {
		log.Fatal(err)
	}
}
