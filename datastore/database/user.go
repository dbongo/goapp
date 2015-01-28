package database

// import (
// 	"time"
//
// 	"github.com/dbongo/hackapp/model"
// 	"gopkg.in/mgo.v2"
// 	"gopkg.in/mgo.v2/bson"
// )
//
// const (
// 	users = "users"
// )
//
// // Userstore ...
// type Userstore struct {
// 	*mgo.Database
// 	C *mgo.Collection
// }
//
// // NewUserstore ...
// func NewUserstore(db *mgo.Database) *Userstore {
// 	return &Userstore{db, db.C(users)}
// }
//
// // GetUser ...
// func (db *Userstore) GetUser(email string) (*model.User, error) {
// 	user := &model.User{}
// 	db.Session = db.Open()
// 	defer db.Close()
// 	if err := db.C.Find(bson.M{"email": email}).One(user); err == mgo.ErrNotFound {
// 		return nil, mgo.ErrNotFound
// 	} else if err != nil {
// 		return nil, err
// 	}
// 	return user, nil
// }
//
// // PutUser ...
// func (db *Userstore) PutUser(user *model.User) error {
// 	db.Session = db.Open()
// 	defer db.Close()
// 	user.Updated = time.Now().Format(time.RFC3339)
// 	return db.C.Update(bson.M{"_id": user.ID}, user)
// }
