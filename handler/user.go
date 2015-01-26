package handler

import (
	"encoding/json"
	"log"
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
	log.Printf("c.Env : %v", c.Env)
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
	c.Env["access_token"] = t.Raw
	log.Printf("access_token: %v", c.Env["access_token"])
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&Session{
		Email:    user.Email,
		Username: user.Username,
		Token:    t.Raw,
	})
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
	c.Env["access_token"] = t.Raw
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&Session{
		Email:    user.Email,
		Username: user.Username,
		Token:    t.Raw,
	})
}

// UpdateUser ...
func UpdateUser(c web.C, w http.ResponseWriter, req *http.Request) {
	log.Printf("access_token: %v", c.Env["access_token"])
	t := c.Env["access_token"].(string)
	user := session.UserFromToken(t)
	u := model.User{}
	defer req.Body.Close()
	if err := json.NewDecoder(req.Body).Decode(&u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if len(u.Email) != 0 {
		user.SetEmail(u.Email)
	}
	if len(u.Name) != 0 {
		user.Name = u.Name
	}
	if err := user.Update(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
