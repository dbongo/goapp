package handler

import (
	"encoding/json"
	"net/http"

	"github.com/dbongo/hackapp/model"
	"github.com/dbongo/hackapp/session"
	"github.com/zenazn/goji/web"
)

// Login ...
type Login struct {
	Email    string
	Password string
}

// Register ...
type Register struct {
	Email    string
	Username string
	Password string
}

// Session ...
type Session struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

// LoginUser ...
func LoginUser(c web.C, w http.ResponseWriter, req *http.Request) {
	l := Login{}
	defer req.Body.Close()
	if err := json.NewDecoder(req.Body).Decode(&l); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := model.AuthUser(l.Email, l.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	t, err := session.New(user.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res := &Session{user.Email, user.Username, t.Raw}
	json.NewEncoder(w).Encode(res)
}

// RegisterUser ...
func RegisterUser(c web.C, w http.ResponseWriter, req *http.Request) {
	r := Register{}
	defer req.Body.Close()
	if err := json.NewDecoder(req.Body).Decode(&r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := model.NewUser(r.Email, r.Username, r.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	t, err := session.New(user.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res := &Session{user.Email, user.Username, t.Raw}
	json.NewEncoder(w).Encode(res)
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
