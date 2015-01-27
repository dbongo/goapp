package model

import (
	"errors"
	"log"
	"time"

	"code.google.com/p/go.crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/dbongo/hackapp/db"
)

// User ...
type User struct {
	ID        bson.ObjectId `bson:"_id" json:"-"`
	Name      string        `bson:"name" json:"name"`
	Email     string        `bson:"email" json:"email"`
	Username  string        `bson:"username" json:"username"`
	Password  string        `bson:"password" json:"-"`
	Created   string        `bson:"created" json:"created"`
	LastLogin string        `bson:"lastlogin" json:"lastlogin"`
	Updated   string        `bson:"updated" json:"updated"`
}

// NewUser ...
func NewUser(email, username, password string) (*User, error) {
	if email == "" || username == "" || password == "" {
		return nil, errors.New("email, username, password are required fields")
	} else if UserExists(email) {
		return nil, errors.New("please provide another email, " + email + " is taken")
	}
	u := User{}
	u.ID = bson.NewObjectId()
	u.Email = email
	u.Username = username
	u.hashPassword(password)
	u.Created = time.Now().Format(time.RFC3339)
	if err := u.Save(); err != nil {
		return nil, err
	}
	return &u, nil
}

// AuthUser ...
func AuthUser(email, password string) (*User, error) {
	u, err := FindUserByEmail(email)
	if err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return nil, err
	}
	u.LastLogin = time.Now().Format(time.RFC3339)
	if err := u.Update(); err != nil {
		return nil, err
	}
	return u, nil
}

// UserExists ...
func UserExists(email string) bool {
	_, err := FindUserByEmail(email)
	if err == nil {
		return true
	}
	return false
}

// FindUserByEmail ...
func FindUserByEmail(email string) (*User, error) {
	ds, err := db.Conn()
	if err != nil {
		return nil, err
	}
	defer ds.Close()
	user := &User{}
	if err := ds.Users().Find(bson.M{"email": email}).One(user); err == mgo.ErrNotFound {
		return nil, mgo.ErrNotFound
	} else if err != nil {
		return nil, err
	}
	return user, nil
}

// Save ...
func (u *User) Save() error {
	ds, err := db.Conn()
	if err != nil {
		return err
	}
	defer ds.Close()
	return ds.Users().Insert(u)
}

// Delete ...
func (u *User) Delete() error {
	ds, err := db.Conn()
	if err != nil {
		return err
	}
	defer ds.Close()
	return ds.Users().Remove(bson.M{"email": u.Email})
}

// Update ...
func (u *User) Update() error {
	ds, err := db.Conn()
	if err != nil {
		return err
	}
	defer ds.Close()
	u.Updated = time.Now().Format(time.RFC3339)
	return ds.Users().Update(bson.M{"_id": u.ID}, u)
}

func (u *User) hashPassword(password string) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	u.Password = string(hash[:])
}
