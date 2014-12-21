package handler

import (
	"log"
	"net/http"

	"github.com/dbongo/goapp/auth"
	"github.com/dbongo/goapp/logger"
	"github.com/dbongo/goapp/model"

	"github.com/zenazn/goji/web"
	"golang.org/x/crypto/bcrypt"
)

// Register ...
func Register(w http.ResponseWriter, r *http.Request) {
	user := new(model.User)
	err := parseBody(r.Body, user)
	if err != nil {
		logger.Error.Print(err)
	}
	err = user.Save()
	if err != nil {
		logger.Error.Print(err)
	}
	writeJSON(w, user.ToString(), http.StatusCreated)
	return
}

// Remove ...
func Remove(c *web.C, w http.ResponseWriter, r *http.Request) {
	user, err := GetCurrentUser(c)
	if err != nil {
		log.Fatal(err)
	}
	user.Delete()
	writeJSON(w, user.ToString(), http.StatusOK)
}

// Login ...
func Login(w http.ResponseWriter, r *http.Request) {
	user := new(model.User)
	err := parseBody(r.Body, user)
	if err != nil {
		logger.Error.Print(err)
	}
	token, err := login(user)
	if err != nil {
		logger.Error.Print(err)
	}
	res := &Response{"token": token.Data}
	ServeJSON(w, res, http.StatusOK)
	return
}

func login(u *model.User) (*auth.Token, error) {
	user, err := model.FindUserByEmail(u.Email)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password))
	if err != nil {
		return nil, err
	}
	token := auth.JWTToken(user)
	return token, nil
}
