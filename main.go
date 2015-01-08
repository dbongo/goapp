package main

import (
	"flag"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	_ "net/http/pprof"
	"net/url"
	"os"
	"strings"

	"github.com/zenazn/goji/web/middleware"

	mw "github.com/dbongo/hackapp/middleware"
	"github.com/dbongo/hackapp/router"
)

var (
	//endpoint = flag.String("e", "/var/run/docker.sock", "Docker endpoint")
	endpoint = flag.String("e", "/Users/dbongo/.boot2docker/boot2docker-vm.sock", "boot2docker endpoint")
	addr     = flag.String("p", ":8080", "Address and port to serve hackapp")
	assets   = flag.String("a", "./static/dist", "Path to the assets")
)

// UnixHandler ...
type UnixHandler struct {
	path string
}

func (h *UnixHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := net.Dial("unix", h.path)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	c := httputil.NewClientConn(conn, nil)
	defer c.Close()
	res, err := c.Do(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	defer res.Body.Close()
	copyHeader(w.Header(), res.Header)
	if _, err := io.Copy(w, res.Body); err != nil {
		log.Println(err)
	}
}

func copyHeader(dest, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dest.Add(k, v)
		}
	}
}

func tcpHandler(e string) http.Handler {
	u, err := url.Parse(e)
	if err != nil {
		log.Fatal(err)
	}
	return httputil.NewSingleHostReverseProxy(u)
}

func unixHandler(e string) http.Handler {
	return &UnixHandler{e}
}

func muxHandler(dir string, e string) http.Handler {
	var (
		h http.Handler

		mux         = router.New()
		fileHandler = http.FileServer(http.Dir(dir))
	)
	if strings.Contains(e, "http") {
		h = tcpHandler(e)
	} else {
		if _, err := os.Stat(e); err != nil {
			if os.IsNotExist(err) {
				log.Fatalf("unix socket %s does not exist", e)
			}
			log.Fatal(err)
		}
		h = unixHandler(e)
	}
	mux.Use(middleware.RequestID)
	mux.Use(mw.Options)
	mux.Use(mw.SetHeaders)
	mux.Use(mw.LogHTTP)
	mux.Use(mw.Recovery)
	mux.Handle("/static/", http.StripPrefix("/static", h))
	mux.Handle("/*", fileHandler)
	return mux
}

func main() {
	flag.Parse()

	// Runtime profiling data. To view all available profiles, open
	// http://localhost:6060/debug/pprof/ in the browser
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	mux := muxHandler(*assets, *endpoint)
	if err := http.ListenAndServe(*addr, mux); err != nil {
		log.Fatal(err)
	}
}
