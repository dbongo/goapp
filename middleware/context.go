package middleware

import (
	"github.com/dbongo/hackapp/model"
	"github.com/zenazn/goji/web"
)

// UserToC ...
func UserToC(c *web.C, user *model.User) {
	c.Env["user"] = user
}

// ToUser ...
func ToUser(c *web.C) *model.User {
	var v = c.Env["user"]
	if v == nil {
		return nil
	}
	u, ok := v.(*model.User)
	if !ok {
		return nil
	}
	return u
}
