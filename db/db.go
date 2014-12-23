package db

import (
	"os"

	"github.com/dbongo/goapp/db/storage"
	"gopkg.in/mgo.v2"
)

const (
	defaultDBAddress = "127.0.0.1:27017"
	defaultDBName    = "appdb"
)

// Storage ...
type Storage struct {
	*storage.Storage
}

func connect() (*storage.Storage, error) {
	addr := os.Getenv("MONGODB_ADDRESS")
	if addr == "" {
		addr = defaultDBAddress
	}
	name := os.Getenv("MONGODB_NAME")
	if name == "" {
		name = defaultDBName
	}
	return storage.MongoDB(addr, name)
}

// Connect ...
func Connect() (*Storage, error) {
	var (
		strg Storage
		err  error
	)
	strg.Storage, err = connect()
	return &strg, err
}

// Users returns the users collection from mongo.
func (s *Storage) Users() *storage.Collection {
	email := mgo.Index{Key: []string{"email"}, Unique: true, Background: false}
	c := s.Collection("users")
	c.EnsureIndex(email)
	return c
}
