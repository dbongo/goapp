package database

import (
	"log"
	"os"

	"github.com/dbongo/hackapp/datastore"

	"gopkg.in/mgo.v2"
)

const (
	defaultAddress = "127.0.0.1:27017"
	defaultName    = "hackdb"
	users          = "users"
)

var (
	addr string
	name string
)

func init() {
	if addr = os.Getenv("MONGODB_PORT_27017_TCP_ADDR"); addr == "" {
		addr = defaultAddress
	}
	if name = os.Getenv("MONGODB_NAME"); name == "" {
		name = defaultName
	}
}

// New ...
func New() *mgo.Database {
	session, err := mgo.Dial(addr)
	if err != nil {
		log.Fatal(err)
	}
	db := session.DB(name)
	return db
}

// NewDatastore ...
func NewDatastore(db *mgo.Database) datastore.Datastore {
	return struct {
		*Userstore
	}{
		NewUserstore(db.C(users)),
	}
}
