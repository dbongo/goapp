package token

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/dbongo/hackapp/model"
	"github.com/dbongo/hackapp/testkeys"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/zenazn/goji/web"
)

var (
	signKey   []byte
	verifyKey []byte
	err       error
)

// initialize keys for signing/verifying jwt tokens
func init() {
	if signKey, err = ioutil.ReadFile(testkeys.Private); err != nil {
		log.Fatal(err)
	}
	if verifyKey, err = ioutil.ReadFile(testkeys.Public); err != nil {
		log.Fatal(err)
	}
}

// Validation ...
func Validation(c *web.C, h http.Handler) http.Handler {
	fn := func(rw http.ResponseWriter, req *http.Request) {
		token, _ := jwt.ParseFromRequest(req, func(token *jwt.Token) (interface{}, error) {
			return verifyKey, nil
		})
		if !token.Valid {
			msg := "invalid access_token"
			http.Error(rw, msg, http.StatusUnauthorized)
			return
		}
		h.ServeHTTP(rw, req)
	}
	return http.HandlerFunc(fn)
}

// New generates signed JWT token in string format
func New(email string) (*jwt.Token, error) {
	token := jwt.New(jwt.GetSigningMethod("RS256"))
	token.Claims["email"] = email
	token.Claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	token.Raw, err = token.SignedString(signKey)
	if err != nil {
		return nil, err
	}
	token.Valid = true
	return token, nil
}

// UserFromToken ...
func UserFromToken(token string) *model.User {
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
