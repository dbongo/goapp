package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/zenazn/goji/web/middleware"

	mw "github.com/dbongo/hackapp/middleware"
	"github.com/dbongo/hackapp/router"
)

var port = flag.String("p", ":3000", "server port")

func main() {
	flag.Parse()

	var mux = router.New()

	mux.Use(middleware.RequestID)
	mux.Use(mw.Options)
	mux.Use(mw.SetHeaders)
	mux.Use(mw.HTTPLogger)
	mux.Use(mw.Recovery)
	http.Handle("/api/", mux)

	if err := http.ListenAndServe(*port, nil); err != nil {
		log.Fatal(err)
	}
}
