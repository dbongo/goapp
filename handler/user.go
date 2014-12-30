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

type responseUser struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

// Login ...
func Login(rw http.ResponseWriter, req *http.Request) {
	lu := loginUser{}
	defer req.Body.Close()
	if err := json.NewDecoder(req.Body).Decode(&lu); err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := model.AuthUser(lu.Email, lu.Password)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	t, err := token.New(user.Email)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	response := responseUser{user.Email, user.Username, t.Raw}
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(&response)
}

// Register ...
func Register(rw http.ResponseWriter, req *http.Request) {
	ru := registerUser{}
	defer req.Body.Close()
	if err := json.NewDecoder(req.Body).Decode(&ru); err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := model.NewUser(ru.Email, ru.Username, ru.Password)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	t, err := token.New(user.Email)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	response := responseUser{user.Email, user.Username, t.Raw}
	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(&response)
}
