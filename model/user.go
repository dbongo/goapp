package model

import "gopkg.in/mgo.v2/bson"

// User ...
type User struct {
	ID        bson.ObjectId `bson:"_id" json:"-"`
	Admin     bool          `bson:"admin" json:"admin"`
	Name      string        `bson:"name" json:"name"`
	Email     string        `bson:"email" json:"email"`
	Username  string        `bson:"username" json:"username"`
	Password  string        `bson:"password" json:"-"`
	Created   string        `bson:"created" json:"created"`
	LastLogin string        `bson:"lastlogin" json:"lastlogin"`
	Updated   string        `bson:"updated" json:"updated"`
}
