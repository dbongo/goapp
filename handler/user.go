package handler

import (
	"encoding/json"
	"net/http"

	"github.com/dbongo/hackapp/model"
	"github.com/dbongo/hackapp/token"
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
func LoginUser(w http.ResponseWriter, req *http.Request) {
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
	t, err := token.New(user.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&Session{
		Email:    user.Email,
		Username: user.Username,
		Token:    t.Raw,
	})
}

// RegisterUser ...
func RegisterUser(w http.ResponseWriter, req *http.Request) {
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
	t, err := token.New(user.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&Session{
		Email:    user.Email,
		Username: user.Username,
		Token:    t.Raw,
	})
}
