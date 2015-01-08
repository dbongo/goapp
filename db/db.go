package db

import (
	"os"

	"gopkg.in/mgo.v2"
)

const (
	defaultAddress = "127.0.0.1:27017"
	defaultName    = "hackdb"
)

var addr string

// DataStore holds the connection with the database.
type DataStore struct {
	*mgo.Session
	name string
}

// Conn ...
func Conn() (*DataStore, error) {
	sess, err := connection()
	ds := DataStore{sess, defaultName}
	return &ds, err
}

func connection() (*mgo.Session, error) {
	if addr = os.Getenv("MONGODB_PORT_27017_TCP_ADDR"); addr == "" {
		addr = defaultAddress
	}
	sess, err := mgo.Dial(addr)
	if err != nil {
		return nil, err
	}
	clone := sess.Clone()
	return clone, nil
}

// DB returns the named database specified in the DDataStore.
func (ds *DataStore) DB() *mgo.Database {
	return ds.Session.DB(ds.name)
}

// Close closes the session, releasing the connection.
func (ds *DataStore) Close() {
	ds.Session.Close()
}

// Users returns the users collection from the database.
func (ds *DataStore) Users() *mgo.Collection {
	email := mgo.Index{
		Key:        []string{"email"},
		Unique:     true,
		Background: true,
	}
	users := ds.Session.DB(ds.name).C("users")
	users.EnsureIndex(email)
	return users
}
