package datastore

import (
	"code.google.com/p/go.net/context"

	"github.com/dbongo/hackapp/model"
)

// Userstore ...
type Userstore interface {
	// GetUser ...
	GetUser(email string) (*model.User, error)

	// PutUser ...
	PutUser(user *model.User) error

	// PostUserLogin ...
	PostUserLogin(email, password string) (*model.User, error)

	// PostUserRegistration ...
	PostUserRegistration(email, username, password string) (*model.User, error)
}

// GetUser ...
func GetUser(c context.Context, email string) (*model.User, error) {
	return FromContext(c).GetUser(email)
}

// PutUser ...
func PutUser(c context.Context, user *model.User) error {
	return FromContext(c).PutUser(user)
}

// PostUserLogin ...
func PostUserLogin(c context.Context, email, password string) (*model.User, error) {
	return FromContext(c).PostUserLogin(email, password)
}

// PostUserRegistration ...
func PostUserRegistration(c context.Context, email, username, password string) (*model.User, error) {
	return FromContext(c).PostUserRegistration(email, username, password)
}
