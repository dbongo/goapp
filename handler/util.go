package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/dbongo/hackapp/logger"
)

var (
	// ErrEncode ...
	ErrEncode = errors.New("Error encoding json")
	// ErrDecode ...
	ErrDecode = errors.New("Error decoding json")
)

func jsonResponseWriter(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(v); err != nil {
		logger.Error.Printf("Error encoding json: %v", err)
		http.Error(w, ErrEncode.Error(), http.StatusBadRequest)
		return
	}
}

func jsonRequest(r *http.Request, v interface{}) bool {
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		logger.Error.Printf("Error decoding json: %v", err)
		return false
	}
	return true
}
