package storage

import (
	"sync"
	"time"

	"gopkg.in/mgo.v2"
)

const period time.Duration = 2 * 60 * time.Second

var (
	connections = make(map[string]*session)
	mut         sync.RWMutex
	ticker      *time.Ticker
)

type session struct {
	master *mgo.Session
	used   time.Time
}

// Storage ...
type Storage struct {
	session *mgo.Session
	name    string
}

// Collection represents a mongo collection.
// It embeds mgo.Collection for operations, and holds a session to MongoDB
type Collection struct {
	*mgo.Collection
}

// Close closes the session with the database
func (c *Collection) Close() {
	c.Collection.Database.Session.Close()
}

func mongodb(addr, dbName string) (*Storage, error) {
	conn, err := mgo.Dial(addr)
	if err != nil {
		return nil, err
	}
	copy := conn.Clone()
	storage := &Storage{session: copy, name: dbName}
	mut.Lock()
	connections[addr] = &session{master: conn, used: time.Now()}
	mut.Unlock()
	return storage, nil
}

// MongoDB ...
func MongoDB(addr, dbName string) (storage *Storage, err error) {
	defer func() {
		if r := recover(); r != nil {
			storage, err = mongodb(addr, dbName)
		}
	}()
	mut.RLock()
	if session, ok := connections[addr]; ok {
		mut.RUnlock()
		if err = session.master.Ping(); err == nil {
			mut.Lock()
			session.used = time.Now()
			connections[addr] = session
			mut.Unlock()
			copy := session.master.Clone()
			return &Storage{copy, dbName}, nil
		}
		return mongodb(addr, dbName)
	}
	mut.RUnlock()
	return mongodb(addr, dbName)
}

// Close the storage and releases the connection
func (s *Storage) Close() {
	s.session.Close()
}

// Collection returns a collection by its name.
// If the collection does not exist, MongoDB will create it.
func (s *Storage) Collection(name string) *Collection {
	return &Collection{s.session.DB(s.name).C(name)}
}

// DB ...
func (s *Storage) DB() *mgo.Database {
	return s.session.DB(s.name)
}

func init() {
	ticker = time.NewTicker(time.Second)
	go retire(ticker)
}

func retire(t *time.Ticker) {
	for _ = range t.C {
		now := time.Now()
		var old []string
		mut.RLock()
		for k, v := range connections {
			if now.Sub(v.used) >= period {
				old = append(old, k)
			}
		}
		mut.RUnlock()
		mut.Lock()
		for _, c := range old {
			connections[c].master.Close()
			delete(connections, c)
		}
		mut.Unlock()
	}
}
