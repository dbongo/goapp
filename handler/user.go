package handler

import (
	"encoding/json"
	"net/http"

	"github.com/dbongo/hackapp/model"
	"github.com/dbongo/hackapp/session"
	"github.com/zenazn/goji/web"
)

// LoginUser ...
func LoginUser(c web.C, w http.ResponseWriter, req *http.Request) {
	login := struct {
		Email, Password string
	}{}
	defer req.Body.Close()
	if err := json.NewDecoder(req.Body).Decode(&login); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := model.AuthUser(login.Email, login.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	token, err := session.New(user.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	resp := struct {
		*model.User
		Token string `json:"token"`
	}{user, token.Raw}
	json.NewEncoder(w).Encode(&resp)
}

// RegisterUser ...
func RegisterUser(c web.C, w http.ResponseWriter, req *http.Request) {
	register := struct {
		Email, Username, Password string
	}{}
	defer req.Body.Close()
	if err := json.NewDecoder(req.Body).Decode(&register); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := model.NewUser(register.Email, register.Username, register.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	token, err := session.New(user.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	resp := struct {
		*model.User
		Token string `json:"token"`
	}{user, token.Raw}
	json.NewEncoder(w).Encode(&resp)
}

// GetCurrentUser ...
func GetCurrentUser(c web.C, w http.ResponseWriter, r *http.Request) {
	var user = ToUser(c)
	if user == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	resp := struct {
		*model.User
	}{user}
	json.NewEncoder(w).Encode(&resp)
}

// PutUser ...
func PutUser(c web.C, w http.ResponseWriter, r *http.Request) {
	//var ctx = context.FromC(c)
	var user = ToUser(c)
	if user == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	defer r.Body.Close()
	in := model.User{}
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if len(in.Email) != 0 {
		user.Email = in.Email
	}
	if len(in.Name) != 0 {
		user.Name = in.Name
	}
	if err := user.Update(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}
