package auth

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"code.google.com/p/go.net/context"
	"github.com/dbongo/goapp/logger"
	"github.com/dbongo/goapp/model"
	jwt "github.com/dgrijalva/jwt-go"
)

const (
	privateKey = "keys/app.rsa"
	publicKey  = "keys/app.rsa.pub"
)

var (
	signKey   []byte
	verifyKey []byte
	err       error
)

// Token ...
type Token struct {
	Data string `json:"auth_token"`
}

func init() {
	signKey, err = ioutil.ReadFile(privateKey)
	if err != nil {
		logger.Error.Print(err)
	}

	verifyKey, err = ioutil.ReadFile(publicKey)
	if err != nil {
		logger.Error.Print(err)
	}
}

// GetUser gets the currently authenticated user for the http.Request.
// The user details will be stored as either a simple API token or JWT bearer token.
func GetUser(c context.Context, r *http.Request) *model.User {
	switch {
	case r.Header.Get("Authorization") != "":
		return getUserBearer(c, r)
	default:
		return nil
	}
}

// JWTToken ...
func JWTToken(user *model.User) *Token {
	t := jwt.New(jwt.GetSigningMethod("RS256"))
	t.Claims["ID"] = user.ID.Hex()
	t.Claims["email"] = user.Email
	t.Claims["exp"] = time.Now().Add(time.Minute * 60 * 730).Unix()
	signed, err := t.SignedString(signKey)
	if err != nil {
		logger.Error.Print(err)
	}
	return &Token{signed}
}

// getUserBearer gets the currently authenticated user for the given bearer token (JWT)
func getUserBearer(c context.Context, r *http.Request) *model.User {
	var tokenstr = r.Header.Get("Authorization")
	fmt.Sscanf(tokenstr, "Bearer %s", &tokenstr)
	return UserFromJWT(tokenstr)
}

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
