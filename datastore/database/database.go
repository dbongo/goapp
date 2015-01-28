package database

// import (
// 	"os"
//
// 	"github.com/dbongo/hackapp/datastore"
//
// 	"gopkg.in/mgo.v2"
// )
//
// const (
// 	defaultAddress  = "127.0.0.1:27017"
// 	defaultName     = "hackdb"
// 	usersCollection = "users"
// )
//
// // Database ...
// type Database struct {
// 	Session *mgo.Session
// 	Use     *mgo.Database
// }
//
// // Open ...
// func (db *Database) Open() *mgo.Session {
// 	return db.Session.Clone()
// }
//
// // Close ...
// func (db *Database) Close() {
// 	db.Session.Close()
// }
//
// // NewDatastore ...
// func NewDatastore(db *mgo.Database) datastore.Datastore {
// 	return struct {
// 		*Userstore
// 	}{
// 		NewUserstore(db),
// 	}
// }
//
// // New ...
// func New() *mgo.Database {
// 	db := &Database{}
// 	db.Session = session()
// 	db.Use = db.Session.DB(defaultName)
// 	return db.Use
// }
//
// func session() *mgo.Session {
// 	var address string
// 	if address = os.Getenv("MONGODB_PORT_27017_TCP_ADDR"); address == "" {
// 		address = defaultAddress
// 	}
// 	session, err := connect(address)
// 	if err != nil {
// 		return nil
// 	}
// 	return session
// }
//
// func connect(address string) (*mgo.Session, error) {
// 	session, err := mgo.Dial(address)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return session, nil
// }
