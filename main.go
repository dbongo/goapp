package main

import (
	"log"
	_ "net/http/pprof"
)

import (
	"net/http"

	"github.com/dbongo/hackapp/logger"
	"github.com/dbongo/hackapp/router"
)

func init() {
	logger.Info.Println("intializing server")

	// create app router
	mux := router.New()

	// attach app router to http server
	http.Handle("/", mux)
}

func main() {
	logger.Info.Println("server running on port 3000")

	// expose dev routes
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	// start the app
	http.ListenAndServe(":3000", nil)
}
