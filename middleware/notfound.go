package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	msg := "The resource you are looking for was not found."
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	body, _ := json.Marshal(msg)
	fmt.Fprint(w, string(body))
	return
}
