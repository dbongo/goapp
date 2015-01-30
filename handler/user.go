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

// LoginUser ...
func LoginUser(c web.C, w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	data := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ctx := context.FromC(c)
	user, err := datastore.AuthUser(ctx, data.Email, data.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	token, err := session.New(user.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response := struct {
		*model.User
		Token string `json:"token"`
	}{user, token}
	json.NewEncoder(w).Encode(&response)
}

// RegisterUser ...
func RegisterUser(c web.C, w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	data := struct {
		Email    string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ctx := context.FromC(c)
	user, err := datastore.CreateUser(ctx, data.Email, data.Username, data.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	token, err := session.New(user.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response := struct {
		*model.User
		Token string `json:"token"`
	}{user, token}
	json.NewEncoder(w).Encode(&response)
}

// GetCurrentUser ...
func GetCurrentUser(c web.C, w http.ResponseWriter, r *http.Request) {
	user := ToUser(c)
	if user == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	json.NewEncoder(w).Encode(user)
}

// PutUser ...
func PutUser(c web.C, w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	user := ToUser(c)
	if user == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	update := model.User{}
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if len(update.Email) != 0 {
		user.Email = update.Email
	}
	if len(update.Name) != 0 {
		user.Name = update.Name
	}
	ctx := context.FromC(c)
	if err := datastore.UpdateUser(ctx, user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}
