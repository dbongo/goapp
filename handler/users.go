package handler

import (
	"net/http"

	"github.com/dbongo/hackapp/datastore"
	"github.com/goji/context"
	"github.com/zenazn/goji/web"
)

// GetUserList ...
func GetUserList(c web.C, w http.ResponseWriter, r *http.Request) {
	var ctx = context.FromC(c)
	users, err := datastore.GetUserList(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	jsonResponseWriter(w, users)
}
