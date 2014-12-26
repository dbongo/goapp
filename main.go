package main

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/subosito/gotenv"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	gmw "github.com/zenazn/goji/web/middleware"
	"golang.org/x/crypto/bcrypt"

	"github.com/dbongo/goapp/logger"
	"github.com/dbongo/goapp/middleware"
	"github.com/dbongo/goapp/model"
	"github.com/dbongo/goapp/router"
)

const (
	privateKey = "keys/app.rsa"     // openssl genrsa -out app.rsa 1024
	publicKey  = "keys/app.rsa.pub" // openssl rsa -in app.rsa -pubout > app.rsa.pub
)

var (
	mux       *web.Mux
	signKey   []byte
	verifyKey []byte
	err       error
)

func init() {
	logger.Info.Println("intializing server")

	// load app env variables
	gotenv.Load(".env")

	// load keys for generating/verifying jwt tokens
	if signKey, err = ioutil.ReadFile(privateKey); err != nil {
		logger.Error.Println(err)
	}
	if verifyKey, err = ioutil.ReadFile(publicKey); err != nil {
		logger.Error.Println(err)
	}

	// remove default goji middleware
	goji.Abandon(gmw.Logger)
	goji.Abandon(gmw.Recoverer)

	// add core middleware
	goji.Use(middleware.RequestID)
	goji.Use(middleware.RequestLogger)
	goji.Use(middleware.Recovery)

	// create api router and add middleware
	mux = router.New()
	mux.Use(authMiddleware)
	mux.Use(middleware.Options)
}

func main() {
	goji.Post("/register", register)
	goji.Post("/login", login)
	goji.Handle("/api/*", mux)
	logger.Info.Println("server listening on :8000")
	goji.Serve()
}

func authMiddleware(c *web.C, h http.Handler) http.Handler {
	fn := func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		token, err := jwt.ParseFromRequest(req, func(token *jwt.Token) (interface{}, error) {
			return verifyKey, nil
		})
		if !token.Valid || err != nil {
			rw.WriteHeader(http.StatusUnauthorized)
			logger.Error.Printf("%v", err)
			return
		}
		rw.WriteHeader(http.StatusOK)
		logger.Info.Printf("%v", token)
		h.ServeHTTP(rw, req)
	}
	return http.HandlerFunc(fn)
}

func register(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	user := &model.User{}
	if err := parseBody(req.Body, user); err != nil {
		logger.Error.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := user.Save(); err != nil {
		logger.Error.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	res := struct {
		Token string `json:"token"`
	}{token(user)}
	json.NewEncoder(rw).Encode(&res)
}

func login(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	user := &model.User{}
	if err := parseBody(req.Body, user); err != nil {
		logger.Error.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := verifyPassword(user); err != nil {
		logger.Error.Println(err)
		rw.WriteHeader(http.StatusUnauthorized)
		return
	}
	res := struct {
		Token string `json:"token"`
	}{token(user)}
	json.NewEncoder(rw).Encode(&res)
}

func verifyPassword(u *model.User) error {
	user, err := model.FindUserByEmail(u.Email)
	if err != nil {
		return err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password)); err != nil {
		return err
	}
	return nil
}

func token(user *model.User) string {
	t := jwt.New(jwt.GetSigningMethod("RS256"))
	t.Claims["ID"] = user.ID.Hex()
	t.Claims["email"] = user.Email
	t.Claims["exp"] = time.Now().Add(time.Minute * 60 * 730).Unix()
	signedToken, err := t.SignedString(signKey)
	if err != nil {
		logger.Error.Println(err)
	}
	return signedToken
}

func userFromToken(token string) *model.User {
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return verifyKey, nil
	})
	if err != nil || !t.Valid {
		return nil
	}
	email, ok := t.Claims["email"].(string)
	if !ok {
		return nil
	}
	user, _ := model.FindUserByEmail(email)
	return user
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
