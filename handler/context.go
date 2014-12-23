package handler

import (
	"github.com/dbongo/goapp/model"
	"github.com/zenazn/goji/web"
)

// ToUser returns the User from the current request context.
// If the User does not exist a nil value is returned.
func ToUser(c web.C) *model.User {
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
