package handler

import (
	"fmt"
	"net/http"

	"github.com/zenazn/goji/web"
)

// HelloWorld ...
func HelloWorld(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "Hello, world!")
}

// HelloName ...
func HelloName(c web.C, rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "Hello, %s!", c.URLParams["name"])
}
