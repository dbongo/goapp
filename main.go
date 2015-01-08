package main

import (
	"flag"
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/zenazn/goji/web/middleware"

	mw "github.com/dbongo/hackapp/middleware"
	"github.com/dbongo/hackapp/router"
	"github.com/dbongo/hackapp/static"
)

var (
	//endpoint = flag.String("endp", "/var/run/docker.sock", "Docker endpoint")
	endpoint   = flag.String("endp", "/Users/dbongo/.boot2docker/boot2docker-vm.sock", "boot2docker endpoint")
	apiport    = flag.String("apip", ":3000", "Port to serve hackapp api")
	uiport     = flag.String("uip", ":8080", "Port to serve hackapp ui")
	assetspath = flag.String("assetsp", "./static/dist", "Path to the assets")
)

func main() {
	flag.Parse()

	// Runtime profiling data. To view all available profiles, open
	// http://localhost:6060/debug/pprof/ in the browser
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	api := router.New()
	api.Use(middleware.RequestID)
	api.Use(mw.Options)
	api.Use(mw.SetHeaders)
	api.Use(mw.LogHTTP)
	api.Use(mw.Recovery)
	go func() {
		log.Fatal(http.ListenAndServe(*apiport, api))
	}()

	ui := static.Handler(*assetspath, *endpoint)
	if err := http.ListenAndServe(*uiport, ui); err != nil {
		log.Fatal(err)
	}
}
