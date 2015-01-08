package handler

import (
	"encoding/json"
	"net/http"

	"github.com/dbongo/hackapp/model"
	"github.com/dbongo/hackapp/token"
)

type loginUser struct {
	Email    string
	Password string
}

type registerUser struct {
	Email    string
	Username string
	Password string
}

type response struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

// Login ...
func Login(w http.ResponseWriter, req *http.Request) {
	lu := loginUser{}
	defer req.Body.Close()
	if err := json.NewDecoder(req.Body).Decode(&lu); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := model.AuthUser(lu.Email, lu.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	t, err := token.New(user.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res := response{user.Email, user.Username, t.Raw}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&res)
}

// Register ...
func Register(w http.ResponseWriter, req *http.Request) {
	ru := registerUser{}
	defer req.Body.Close()
	if err := json.NewDecoder(req.Body).Decode(&ru); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := model.NewUser(ru.Email, ru.Username, ru.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	t, err := token.New(user.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res := response{user.Email, user.Username, t.Raw}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&res)
}
