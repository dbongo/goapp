package db

import (
	"os"

	"github.com/dbongo/goapp/db/storage"
	"gopkg.in/mgo.v2"
)

const (
	defaultAddress = "127.0.0.1:27017"
	defaultName    = "appdb"
)

// Storage ...
type Storage struct {
	*storage.Storage
}

func connect() (*storage.Storage, error) {
	addr := os.Getenv("MONGODB_PORT_27017_TCP_ADDR")
	if addr == "" {
		addr = defaultAddress
	}
	name := os.Getenv("DB_NAME")
	if name == "" {
		name = defaultName
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
	emailIndex := mgo.Index{
		Key:        []string{"email"},
		Unique:     true,
		Background: true,
	}
	c := s.Collection("users")
	c.EnsureIndex(emailIndex)
	//storage.LogStats()
	return c
}
