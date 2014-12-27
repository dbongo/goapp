package db

import (
	"os"

	"github.com/dbongo/hackapp/db/storage"
	"gopkg.in/mgo.v2"
)

const (
	defaultAddress = "127.0.0.1:27017"
	dbName         = "hackdb"
)

var (
	strg Storage
	addr string
	err  error
)

// Storage ...
type Storage struct {
	*storage.Storage
}

func connect() (*storage.Storage, error) {
	if addr = os.Getenv("MONGODB_PORT_27017_TCP_ADDR"); addr == "" {
		addr = defaultAddress
	}
	return storage.MongoDB(addr, dbName)
}

// Connect ...
func Connect() (*Storage, error) {
	strg.Storage, err = connect()
	return &strg, err
}

// Users returns the users collection from mongo.
func (s *Storage) Users() *storage.Collection {
	emailIndex := mgo.Index{
		Key:        []string{"email"},
		Unique:     true,
		Background: true,
	}
	c := s.Collection("users")
	c.EnsureIndex(emailIndex)
	return c
}
