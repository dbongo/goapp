package model

import (
	"encoding/json"
	"log"

	"code.google.com/p/go.crypto/bcrypt"
	"github.com/dbongo/goapp/db"
	"github.com/dbongo/goapp/errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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
		log.Fatal(err)
	}
	defer conn.Close()
	if u.Name == "" || u.Email == "" || u.Username == "" || u.Password == "" {
		message := "Name/Email/Username/Password cannot be empty."
		return &errors.ValidationError{Message: message}
	}
	u.ID = bson.NewObjectId()
	u.HashPassword()
	err = conn.Users().Insert(u)
	if mgo.IsDup(err) {
		message := "Someone already has that email. Could you try another?"
		return &errors.ValidationError{Message: message}
	}
	return err
}

// Delete ...
func (u *User) Delete() error {
	conn, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	err = conn.Users().Remove(bson.M{"email": u.Email})
	if err == mgo.ErrNotFound {
		message := "User not found."
		return &errors.ValidationError{Message: message}
	}
	return err
}

// Update ...
func (u *User) Update() error {
	conn, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	return conn.Users().Update(bson.M{"email": u.Email}, u)
}

// HashPassword ...
func (u *User) HashPassword() {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
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

// ToString returns a user minus their password.
func (u *User) ToString() string {
	u.Password = ""
	user, _ := json.Marshal(u)
	return string(user)
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
	var user User
	err = conn.Users().Find(bson.M{"email": email}).One(&user)
	if err == mgo.ErrNotFound {
		return nil, &errors.ValidationError{Message: "User not found"}
	}
	return &user, nil
}
