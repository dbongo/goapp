package token

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dbongo/hackapp/logger"
	"github.com/dbongo/hackapp/model"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/zenazn/goji/web"
)

// openssl genrsa -out certs/jwt.rsa 1024
// openssl rsa -in certs/jwt.rsa -pubout > certs/jwt.rsa.pub
const (
	privateKey = "certs/jwt.rsa"
	publicKey  = "certs/jwt.rsa.pub"
)

var (
	signKey   []byte
	verifyKey []byte
	err       error
)

// initialize keys for signing/verifying jwt tokens
func init() {
	if signKey, err = ioutil.ReadFile(privateKey); err != nil {
		logger.Error.Println(err)
	}
	if verifyKey, err = ioutil.ReadFile(publicKey); err != nil {
		logger.Error.Println(err)
	}
}

// New generates signed JWT token in string format
func New(email string) string {
	t := jwt.New(jwt.GetSigningMethod("RS256"))
	t.Claims["email"] = email
	t.Claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	token, err := t.SignedString(signKey)
	if err != nil {
		logger.Error.Println(err)
		return ""
	}
	return token
}

// Validation ...
func Validation(c *web.C, h http.Handler) http.Handler {
	fn := func(rw http.ResponseWriter, req *http.Request) {
		token, err := jwt.ParseFromRequest(req, func(token *jwt.Token) (interface{}, error) {
			return verifyKey, nil
		})
		if !token.Valid || err != nil {
			rw.WriteHeader(http.StatusUnauthorized)
			logger.Error.Printf("%v", err)
			return
		}
		h.ServeHTTP(rw, req)
	}
	return http.HandlerFunc(fn)
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
