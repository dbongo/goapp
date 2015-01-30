package handler

import (
	"github.com/dbongo/hackapp/model"
	"github.com/zenazn/goji/web"
)

// ToUser ...
func ToUser(c web.C) *model.User {
	value := c.Env["user"]
	if value == nil {
		return nil
	}
	user, ok := value.(*model.User)
	if !ok {
		return nil
	}
	return user
}
