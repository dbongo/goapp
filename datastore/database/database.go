package database

import (
	"log"

	"github.com/dbongo/hackapp/datastore"

	"gopkg.in/mgo.v2"
)

const (
	users = "users"
)

// New ...
func New(addr, name string) *mgo.Database {
	session, err := mgo.Dial(addr)
	if err != nil {
		log.Fatal(err)
	}
	s := session.Clone()
	return s.DB(name)
}

// NewDatastore ...
func NewDatastore(db *mgo.Database) datastore.Datastore {
	return struct {
		*UserCollection
	}{
		NewUserCollection(db.C(users)),
	}
}
