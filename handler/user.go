package handler

import (
	"encoding/json"
	"net/http"

	"github.com/dbongo/hackapp/logger"
	"github.com/dbongo/hackapp/model"
	"github.com/dbongo/hackapp/token"
	"golang.org/x/crypto/bcrypt"
)

// Login ...
func Login(rw http.ResponseWriter, req *http.Request) {
	user := new(model.User)
	defer req.Body.Close()

	if err := json.NewDecoder(req.Body).Decode(user); err != nil {
		logger.Error.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := authenticate(user); err != nil {
		logger.Error.Println(err)
		rw.WriteHeader(http.StatusUnauthorized)
		return
	}

	response := struct {
		Token string `json:"token"`
	}{token.New(user.Email)}

	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(&response)
}

// Register ...
func Register(rw http.ResponseWriter, req *http.Request) {
	user := new(model.User)
	defer req.Body.Close()

	if err := json.NewDecoder(req.Body).Decode(user); err != nil {
		logger.Error.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := user.Save(); err != nil {
		logger.Error.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	response := struct {
		Token string `json:"token"`
	}{token.New(user.Email)}

	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(&response)
}

func authenticate(u *model.User) error {
	user, err := model.FindUserByEmail(u.Email)
	if err != nil {
		return err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password)); err != nil {
		return err
	}
	return nil
}

// // CurrentUser ...
// func CurrentUser(c web.C, w http.ResponseWriter, r *http.Request) {
// 	var user = ToUser(c)
// 	if user == nil {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		return
// 	}
// 	// return private data for the currently authenticated user,
// 	// specifically, their auth token.
// 	data := struct {
// 		*model.User
// 		Token string `json:"token"`
// 	}{user, user.Token}
// 	json.NewEncoder(w).Encode(&data)
// }
