package datastore

import (
	"code.google.com/p/go.net/context"

	"github.com/dbongo/hackapp/model"
)

type Userstore interface {
	// GetUser retrieves a specific user from the datastore for the given ID.
	GetUser(id int64) (*model.User, error)

	// GetUserLogin retrieves a user from the datastore for the specified remote and login name.
	//GetUserLogin(remote, login string) (*model.User, error)

	// GetUserToken retrieves a user from the datastore with the specified token.
	//GetUserToken(token string) (*model.User, error)

	// GetUserList retrieves a list of all users from the datastore that are registered in the system.
	GetUserList() ([]*model.User, error)

	// PostUser saves a User in the datastore.
	PostUser(user *model.User) error

	// PutUser saves a user in the datastore.
	PutUser(user *model.User) error

	// DelUser removes the user from the datastore.
	//DelUser(user *model.User) error
}

// GetUser retrieves a specific user from the datastore for the given ID.
func GetUser(c context.Context, id int64) (*model.User, error) {
	return FromContext(c).GetUser(id)
}

// GetUserList retrieves a list of all users from the datastore that are registered in the system.
func GetUserList(c context.Context) ([]*model.User, error) {
	return FromContext(c).GetUserList()
}

// PostUser saves a User in the datastore.
func PostUser(c context.Context, user *model.User) error {
	return FromContext(c).PostUser(user)
}

// PutUser saves a user in the datastore.
func PutUser(c context.Context, user *model.User) error {
	return FromContext(c).PutUser(user)
}

// // UserCollection ...
// type UserCollection interface {
// 	GetUser(email string) (*model.User, error)
// 	GetUserList() ([]*model.User, error)
// 	UpdateUser(user *model.User) error
// 	AuthUser(email, password string) (*model.User, error)
// 	CreateUser(email, username, password string) (*model.User, error)
// }
//
// // GetUser ...
// func GetUser(c context.Context, email string) (*model.User, error) {
// 	return FromContext(c).GetUser(email)
// }
//
// // GetUserList ...
// func GetUserList(c context.Context) ([]*model.User, error) {
// 	return FromContext(c).GetUserList()
// }
//
// // UpdateUser ...
// func UpdateUser(c context.Context, user *model.User) error {
// 	return FromContext(c).UpdateUser(user)
// }
//
// // AuthUser ...
// func AuthUser(c context.Context, email, password string) (*model.User, error) {
// 	return FromContext(c).AuthUser(email, password)
// }
//
// // CreateUser ...
// func CreateUser(c context.Context, email, username, password string) (*model.User, error) {
// 	return FromContext(c).CreateUser(email, username, password)
// }
