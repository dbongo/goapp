package handler

import (
	"encoding/json"
	"net/http"

	"github.com/dbongo/hackapp/datastore"
	"github.com/dbongo/hackapp/model"
	"github.com/dbongo/hackapp/session"

	"github.com/goji/context"
	"github.com/zenazn/goji/web"
)

var (
	err error
)

// LoginUser ...
func LoginUser(c web.C, w http.ResponseWriter, r *http.Request) {
	data := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}
	res := struct {
		User  *model.User `json:"user"`
		Token string      `json:"token"`
	}{}
	if jsonRequest(r, &data) {
		ctx := context.FromC(c)
		if res.User, err = datastore.AuthUser(ctx, data.Email, data.Password); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if res.Token, err = session.New(res.User.Email); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
	jsonResponseWriter(w, res)
}

// RegisterUser ...
func RegisterUser(c web.C, w http.ResponseWriter, r *http.Request) {
	data := struct {
		Email    string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
	}{}
	res := struct {
		User  *model.User `json:"user"`
		Token string      `json:"token"`
	}{}
	if jsonRequest(r, &data) {
		ctx := context.FromC(c)
		if res.User, err = datastore.CreateUser(ctx, data.Email, data.Username, data.Password); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if res.Token, err = session.New(res.User.Email); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
	jsonResponseWriter(w, res)
}

// GetCurrentUser accepts a request to retrieve the currently authenticated
// user from the datastore and return in JSON format.
//
// GET /api/user
//
func GetCurrentUser(c web.C, w http.ResponseWriter, r *http.Request) {
	user := ToUser(c)
	if user == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	data := struct {
		*model.User
	}{user}
	json.NewEncoder(w).Encode(&data)
}

// GetCurrentUser ...
// func GetCurrentUser(c web.C, w http.ResponseWriter, r *http.Request) {
// 	res := struct {
// 		User *model.User `json:"user"`
// 	}{}
// 	if res.User = ToUser(c); res.User == nil {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		return
// 	}
// 	jsonResponseWriter(w, res)
// }

// PutUser ...
func PutUser(c web.C, w http.ResponseWriter, r *http.Request) {
	data := struct {
		Name     string `json:"name"`
		Username string `json:"username"`
	}{}
	res := struct {
		User *model.User `json:"user"`
	}{}
	if res.User = ToUser(c); res.User == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if jsonRequest(r, &data) {
		if len(data.Name) != 0 {
			res.User.Name = data.Name
		}
		if len(data.Username) != 0 {
			res.User.Username = data.Username
		}
		ctx := context.FromC(c)
		if err := datastore.UpdateUser(ctx, res.User); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	jsonResponseWriter(w, res)
}
