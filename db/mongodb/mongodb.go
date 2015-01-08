package mongodb

// import (
// 	"sync"
// 	"time"
//
// 	"gopkg.in/mgo.v2"
// )
//
// const period time.Duration = 3 * 24 * time.Hour
//
// var (
// 	connections = make(map[string]*session) // pool of connections
// 	mut         sync.RWMutex                // for pool thread safety
// 	ticker      *time.Ticker                // for garbage collection
// )
//
// type session struct {
// 	master *mgo.Session
// 	used   time.Time
// }
//
// // Session holds the connection with the database.
// type Session struct {
// 	session *mgo.Session
// 	name    string
// }
//
// // Collection represents a database collection. It embeds mgo.Collection for
// // operations, and holds a session to mongodb.
// type Collection struct {
// 	*mgo.Collection
// }
//
// // Close closes the session with the database.
// func (c *Collection) Close() {
// 	c.Collection.Database.Session.Close()
// }
//
// func database(addr, name string) (*Session, error) {
// 	conn, err := mgo.Dial(addr)
// 	if err != nil {
// 		return nil, err
// 	}
// 	clone := conn.Clone()
// 	s := &Session{session: clone, name: name}
// 	mut.Lock()
// 	connections[addr] = &session{master: conn, used: time.Now()}
// 	mut.Unlock()
// 	return s, nil
// }
//
// // New dials a mongodb database using the addr parameter(a mongodb
// // connection URI) and name parameter(the name of the database), returning a
// // connection represented by the type Session. This function returns either
// // a pointer to a Session, or a non-nil error.
// func New(addr, name string) (s *Session, err error) {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			s, err = database(addr, name)
// 		}
// 	}()
// 	mut.RLock()
// 	if session, ok := connections[addr]; ok {
// 		mut.RUnlock()
// 		if err = session.master.Ping(); err == nil {
// 			mut.Lock()
// 			session.used = time.Now()
// 			connections[addr] = session
// 			mut.Unlock()
// 			clone := session.master.Clone()
// 			return &Session{clone, name}, nil
// 		}
// 		return database(addr, name)
// 	}
// 	mut.RUnlock()
// 	return database(addr, name)
// }
//
// // Close closes the session, releasing the connection.
// func (s *Session) Close() {
// 	s.session.Close()
// }
//
// // Collection returns a database collection by name. If the collection does not
// // exist, it will be created.
// func (s *Session) Collection(name string) *Collection {
// 	return &Collection{s.session.DB(s.name).C(name)}
// }
//
// // DB returns the named database specified in the Session.
// func (s *Session) DB() *mgo.Database {
// 	return s.session.DB(s.name)
// }
//
// func init() {
// 	ticker = time.NewTicker(time.Second)
// 	go retire(ticker)
// }
//
// func retire(t *time.Ticker) {
// 	for _ = range t.C {
// 		now := time.Now()
// 		var old []string
// 		mut.RLock()
// 		for k, v := range connections {
// 			if now.Sub(v.used) >= period {
// 				old = append(old, k)
// 			}
// 		}
// 		mut.RUnlock()
// 		mut.Lock()
// 		for _, c := range old {
// 			connections[c].master.Close()
// 			delete(connections, c)
// 		}
// 		mut.Unlock()
// 	}
// }
