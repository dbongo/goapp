package auth

import (
	"io/ioutil"
	"time"

	"github.com/dbongo/goapp/logger"
	"github.com/dbongo/goapp/model"
	jwt "github.com/dgrijalva/jwt-go"
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

// TokenInfo ...
type TokenInfo struct {
	Raw string `json:"token"`
}

// Token ...
func Token(user *model.User) *TokenInfo {
	t := jwt.New(jwt.GetSigningMethod("RS256"))
	t.Claims["ID"] = user.ID.Hex()
	t.Claims["email"] = user.Email
	t.Claims["exp"] = time.Now().Add(time.Minute * 60 * 730).Unix()
	signed, err := t.SignedString(signKey)
	if err != nil {
		logger.Error.Println(err)
	}
	return &TokenInfo{signed}
}

// UserFromJWT ...
func UserFromJWT(token string) *model.User {
	var t, err = jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return verifyKey, nil
	})
	if err != nil || !t.Valid {
		return nil
	}
	var email, ok = t.Claims["email"].(string)
	if !ok {
		return nil
	}
	var user, _ = model.FindUserByEmail(email)
	return user
}
