package db

import (
	"os"

	"github.com/dbongo/goapp/db/store"
	"gopkg.in/mgo.v2"
)

const (
	defaultDBURL  = "127.0.0.1:27017"
	defaultDBName = "appdb"
)

var err error

// Store ...
type Store struct {
	*store.Store
}

func connect() (*store.Store, error) {
	url := os.Getenv("MONGODB_ADDRESS")
	if url == "" {
		url = defaultDBURL
	}
	name := os.Getenv("MONGODB_NAME")
	if name == "" {
		name = defaultDBName
	}
	return store.Open(url, name)
}

// Connect ...
func Connect() (*Store, error) {
	var store Store
	store.Store, err = connect()
	return &store, err
}

// Users returns the users collection from MongoDB.
func (s *Store) Users() *store.Collection {
	email := mgo.Index{Key: []string{"email"}, Unique: true, Background: false}
	c := s.Collection("users")
	c.EnsureIndex(email)
	return c
}
