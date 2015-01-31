package datastore

import (
	"code.google.com/p/go.net/context"

	"github.com/dbongo/hackapp/model"
)

// UserCollection ...
type UserCollection interface {
	GetUser(email string) (*model.User, error)
	GetUserList() ([]*model.User, error)
	UpdateUser(user *model.User) error
	AuthUser(email, password string) (*model.User, error)
	CreateUser(email, username, password string) (*model.User, error)
}

// GetUser ...
func GetUser(c context.Context, email string) (*model.User, error) {
	return FromContext(c).GetUser(email)
}

// GetUserList ...
func GetUserList(c context.Context) ([]*model.User, error) {
	return FromContext(c).GetUserList()
}

// UpdateUser ...
func UpdateUser(c context.Context, user *model.User) error {
	return FromContext(c).UpdateUser(user)
}

// AuthUser ...
func AuthUser(c context.Context, email, password string) (*model.User, error) {
	return FromContext(c).AuthUser(email, password)
}

// CreateUser ...
func CreateUser(c context.Context, email, username, password string) (*model.User, error) {
	return FromContext(c).CreateUser(email, username, password)
}
