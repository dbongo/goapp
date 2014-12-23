package main

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/subosito/gotenv"
	"github.com/zenazn/goji"
	gmw "github.com/zenazn/goji/web/middleware"
	"golang.org/x/crypto/bcrypt"

	"github.com/dbongo/goapp/auth"
	"github.com/dbongo/goapp/logger"
	"github.com/dbongo/goapp/middleware"
	"github.com/dbongo/goapp/model"
	"github.com/dbongo/goapp/router"
)

// intialize app env vars
func init() {
	gotenv.Load(".env")

	// remove default goji middleware
	goji.Abandon(gmw.Logger)
	goji.Abandon(gmw.Recoverer)

	// add core middleware handlers
	goji.Use(middleware.RequestID)
	goji.Use(middleware.RequestLogger)
	goji.Use(middleware.Recovery)
}

func main() {
	logger.Info.Println("initializing server")

	// user login/registration routes
	goji.Post("/users", register)
	goji.Post("/login", login)

	// create api router and require jwt token authentication
	// for every route prefixed with /api/*
	mux := router.New()
	mux.Use(middleware.Auth)

	goji.Handle("/api/*", mux)
	goji.Serve()
}

func register(w http.ResponseWriter, r *http.Request) {
	user := &model.User{}
	if err := parseBody(r.Body, user); err != nil {
		logger.Error.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := user.Save(); err != nil {
		logger.Error.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	data := struct {
		Token string `json:"token"`
	}{auth.Token(user).Raw}
	json.NewEncoder(w).Encode(&data)
}

func login(w http.ResponseWriter, r *http.Request) {
	user := &model.User{}
	if err := parseBody(r.Body, user); err != nil {
		logger.Error.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := verify(user); err != nil {
		logger.Error.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	data := struct {
		Token string `json:"token"`
	}{auth.Token(user).Raw}
	json.NewEncoder(w).Encode(&data)
}

func verify(u *model.User) error {
	user, err := model.FindUserByEmail(u.Email)
	if err != nil {
		return err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password)); err != nil {
		return err
	}
	return nil
}

func parseBody(body io.ReadCloser, r interface{}) error {
	defer body.Close()
	b, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(b, &r); err != nil {
		return errors.New("The request was bad-formed.")
	}
	return nil
}
