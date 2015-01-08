package handler

import (
	"encoding/json"
	"net/http"

	"github.com/dbongo/hackapp/model"
	"github.com/dbongo/hackapp/token"
)

type credentials struct {
	Email    string
	Password string
}

type registration struct {
	Email    string
	Username string
	Password string
}

type profile struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

// Login ...
func Login(w http.ResponseWriter, req *http.Request) {
	login := credentials{}
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
	t, err := token.New(user.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&profile{
		Email:    user.Email,
		Username: user.Username,
		Token:    t.Raw,
	})
}

// Register ...
func Register(w http.ResponseWriter, req *http.Request) {
	register := registration{}
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
	t, err := token.New(user.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&profile{
		Email:    user.Email,
		Username: user.Username,
		Token:    t.Raw,
	})
}
