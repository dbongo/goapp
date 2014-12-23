package model

import (
	"errors"

	"code.google.com/p/go.crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/dbongo/goapp/db"
	"github.com/dbongo/goapp/logger"
)

// User ...
type User struct {
	ID       bson.ObjectId `bson:"_id" json:"_id"`
	Name     string        `bson:"name" json:"name,omitempty"`
	Email    string        `bson:"email" json:"email,omitempty"`
	Username string        `bson:"username" json:"username,omitempty"`
	Password string        `bson:"password" json:"password,omitempty"`
}

// Save ...
func (u *User) Save() error {
	conn, err := db.Connect()
	if err != nil {
		logger.Error.Println(err)
		return err
	}
	defer conn.Close()
	if u.Email == "" || u.Username == "" || u.Password == "" {
		return errors.New("email / username / password fields can not be blank")
	}
	u.ID = bson.NewObjectId()
	u.HashPassword()
	return conn.Users().Insert(u)
}

// Delete ...
func (u *User) Delete() error {
	conn, err := db.Connect()
	if err != nil {
		logger.Error.Println(err)
		return err
	}
	defer conn.Close()
	return conn.Users().Remove(bson.M{"email": u.Email})
}

// Update ...
func (u *User) Update() error {
	conn, err := db.Connect()
	if err != nil {
		logger.Error.Println(err)
		return err
	}
	defer conn.Close()
	return conn.Users().Update(bson.M{"email": u.Email}, u)
}

// HashPassword ...
func (u *User) HashPassword() {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.Error.Println(err)
	}
	u.Password = string(hash[:])
}

// Valid ...
func (u *User) Valid() bool {
	_, err := FindUserByEmail(u.Email)
	if err == nil {
		return true
	}
	return false
}

// FindUserByID ...
func FindUserByID(id bson.ObjectId) (*User, error) {
	conn, err := db.Connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	user := &User{}
	err = conn.Users().FindId(id).One(user)
	if err != nil {
		return nil, err
	} else if user.ID == "" {
		return nil, err
	}
	return user, nil
}

// FindUserByEmail ...
func FindUserByEmail(email string) (*User, error) {
	conn, err := db.Connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	user := &User{}
	err = conn.Users().Find(bson.M{"email": email}).One(user)
	if err == mgo.ErrNotFound {
		return nil, mgo.ErrNotFound
	}
	return user, nil
}
