package database

import (
	"github.com/dbongo/hackapp/datastore"

	"gopkg.in/mgo.v2"
)

const (
	users = "users"
)

// NewDatastore ...
func NewDatastore(db *mgo.Database) datastore.Datastore {
	return struct {
		*Userstore
	}{
		NewUserstore(db.C(users)),
	}
}
