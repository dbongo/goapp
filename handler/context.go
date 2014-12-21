package handler

import (
	"errors"

	"github.com/dbongo/goapp/model"
	"github.com/zenazn/goji/web"
)

const currentUser string = "user"

// ErrUserNotSigned ...
var ErrUserNotSigned = errors.New("User is not signed in.")

// SetCurrentUser ...
func SetCurrentUser(c *web.C, user interface{}) {
	user = user.(*model.User)
	c.Env[currentUser] = user
}

// GetCurrentUser ...
func GetCurrentUser(c *web.C) (*model.User, error) {
	user, ok := c.Env[currentUser].(*model.User)
	if !ok || !user.Valid() {
		return nil, ErrUserNotSigned
	}
	return user, nil
}
