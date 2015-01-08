package static

import (
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
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

// Handler ...
func Handler(dir string, e string) http.Handler {
	var (
		h http.Handler

		mux         = http.NewServeMux()
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
	mux.Handle("/static/", http.StripPrefix("/static", h))
	mux.Handle("/", fileHandler)
	return mux
}
