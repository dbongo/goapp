package handler

import (
	"fmt"
	"net/http"

	"github.com/zenazn/goji/web"
)

// HelloWorld ...
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}

// HelloName ...
func HelloName(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", c.URLParams["name"])
}
