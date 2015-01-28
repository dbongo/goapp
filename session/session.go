package session

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/dbongo/hackapp/logger"
	"github.com/dbongo/hackapp/model"
	"github.com/dbongo/hackapp/testkeys"
	"golang.org/x/net/context"

	"github.com/dgrijalva/jwt-go"
)

var (
	// ErrClaimsEmail ...
	ErrClaimsEmail = errors.New("email from token not ok")
)

var (
	signKey   []byte
	verifyKey []byte
)

// initialize keys for signing/verifying jwt tokens
func init() {
	var err error
	if signKey, err = ioutil.ReadFile(testkeys.Private); err != nil {
		log.Fatal(err)
	}
	if verifyKey, err = ioutil.ReadFile(testkeys.Public); err != nil {
		log.Fatal(err)
	}
}

// New generates signed JWT token in string format
func New(email string) (string, error) {
	token := jwt.New(jwt.GetSigningMethod("RS256"))
	token.Claims["email"] = email
	token.Claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	signedToken, err := token.SignedString(signKey)
	return signedToken, err
}

// GetUser ...
func GetUser(c context.Context, req *http.Request) *model.User {
	if req.Header.Get("Authorization") != "" {
		user, err := tokenFromRequest(c, req)
		if err != nil {
			logger.Error.Println(err)
		}
		return user
	}
	return nil
}

func tokenFromRequest(c context.Context, req *http.Request) (*model.User, error) {
	token, err := jwt.ParseFromRequest(req, func(token *jwt.Token) (interface{}, error) {
		return verifyKey, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	return userFromToken(c, token)
}

func userFromToken(c context.Context, token *jwt.Token) (*model.User, error) {
	email, ok := token.Claims["email"].(string)
	if !ok {
		return nil, ErrClaimsEmail
	}
	user, err := model.FindUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}
