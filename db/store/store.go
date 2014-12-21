package store

import (
	"sync"
	"time"

	"gopkg.in/mgo.v2"
)

const period time.Duration = 7 * 24 * time.Hour

var (
	connections = make(map[string]*session)
	mut         sync.RWMutex
	ticker      *time.Ticker
)

type session struct {
	master *mgo.Session
	used   time.Time
}

// Store holds the connection with the database
type Store struct {
	session *mgo.Session
	name    string
}

// Collection represents a mongodb collection.
// It embeds mgo.Collection for operations, and holds a session to MongoDB
type Collection struct {
	*mgo.Collection
}

// Close closes the session with the database
func (c *Collection) Close() {
	c.Collection.Database.Session.Close()
}

func open(addr, dbname string) (*Store, error) {
	sess, err := mgo.Dial(addr)
	if err != nil {
		return nil, err
	}
	copy := sess.Clone()
	store := &Store{session: copy, name: dbname}
	mut.Lock()
	connections[addr] = &session{master: sess, used: time.Now()}
	mut.Unlock()
	return store, nil
}

// Open dials to the MongoDB database, and return the connection (represented by the type Store)
func Open(addr, dbname string) (store *Store, err error) {
	defer func() {
		if r := recover(); r != nil {
			store, err = open(addr, dbname)
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
			return &Store{copy, dbname}, nil
		}
		return open(addr, dbname)
	}
	mut.RUnlock()
	return open(addr, dbname)
}

// Close closes the store and releases the connection
func (s *Store) Close() {
	s.session.Close()
}

// Collection returns a collection by its name.
// If the collection does not exist, MongoDB will create it.
func (s *Store) Collection(name string) *Collection {
	return &Collection{s.session.DB(s.name).C(name)}
}

// DB ...
func (s *Store) DB() *mgo.Database {
	return s.session.DB(s.name)
}

func init() {
	ticker = time.NewTicker(time.Hour)
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
