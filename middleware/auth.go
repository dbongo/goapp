package middleware

import (
	"io/ioutil"
	"net/http"

	"github.com/dbongo/goapp/logger"
	jwt "github.com/dgrijalva/jwt-go"

	"github.com/zenazn/goji/web"
)

const (
	privateKey = "keys/app.rsa"     // openssl genrsa -out app.rsa 1024
	publicKey  = "keys/app.rsa.pub" // openssl rsa -in app.rsa -pubout > app.rsa.pub
)

var signKey, verifyKey []byte

func init() {
	var err error
	if signKey, err = ioutil.ReadFile(privateKey); err != nil {
		logger.Error.Println(err)
	}
	if verifyKey, err = ioutil.ReadFile(publicKey); err != nil {
		logger.Error.Println(err)
	}
}

// Auth verifies the jwt token in the request headers
func Auth(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		token, err := jwt.ParseFromRequest(r, func(token *jwt.Token) (interface{}, error) {
			return verifyKey, nil
		})
		if token.Valid {
			w.WriteHeader(http.StatusOK)
			logger.Info.Printf("%v", token)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			logger.Error.Printf("%v", err)
			return
		}
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
