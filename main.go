package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"

	"github.com/dbongo/hackapp/logger"
	"github.com/dbongo/hackapp/router"
	"github.com/zenazn/goji/web"
)

var (
	port int
	mux  *web.Mux
)

func init() {
	flag.IntVar(&port, "p", 3000, "port")

	mux = router.Init()
}

func main() {
	flag.Parse()

	// Runtime profiling data. To view all available profiles, open http://localhost:6060/debug/pprof/ in the browser
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", port),
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	logger.Info.Printf("hackapp listening on %s", s.Addr)
	log.Fatal(s.ListenAndServe())
}
