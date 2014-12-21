package middleware

import (
	"github.com/dbongo/goapp/logger"
	"github.com/dbongo/goapp/model"
	"github.com/zenazn/goji/web"
)

// UserToC sets the User in the current web context.
func UserToC(c *web.C, user *model.User) {
	c.Env["user"] = user
}

// ToUser returns the User from the current request context.
// If the User does not exist a nil value is returned.
func ToUser(c *web.C) *model.User {
	var v = c.Env["user"]
	logger.Info.Printf("current user from c.Env[user] is %v", v)
	if v == nil {
		return nil
	}
	u, ok := v.(*model.User)
	if !ok {
		return nil
	}
	return u
}
