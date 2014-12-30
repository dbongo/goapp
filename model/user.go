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
	Name      string    `bson:"name" json:"name,omitempty"`
	Email     string    `bson:"email" json:"email,omitempty"`
	Username  string    `bson:"username" json:"username,omitempty"`
	Password  string    `bson:"password" json:"-"`
	Created   time.Time `bson:"created" json:"created"`
	LastLogin time.Time `bson:"lastlogin" json:"lastlogin"`
}

// NewUser ...
func NewUser(email, username, password string) (*User, error) {
	if email == "" || username == "" || password == "" {
		return nil, errors.New("email, username, password are required fields")
	} else if UserExists(email) {
		return nil, errors.New("please provide another email, " + email + " is taken")
	}
	u := User{}
	u.Email = email
	u.Username = username
	u.hashPassword(password)
	if err := u.Save(); err != nil {
		return nil, err
	}
	return &u, nil
}

// AuthUser authenticates the user's login credentials
func AuthUser(email, password string) (*User, error) {
	u, err := FindUserByEmail(email)
	if err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return nil, err
	}
	u.LastLogin = time.Now()
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
	conn, err := db.Connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	user := &User{}
	if err := conn.Users().Find(bson.M{"email": email}).One(user); err == mgo.ErrNotFound {
		return nil, mgo.ErrNotFound
	} else if err != nil {
		return nil, err
	}
	return user, nil
}

// Save ...
func (u *User) Save() error {
	conn, err := db.Connect()
	if err != nil {
		return err
	}
	defer conn.Close()
	u.Created = time.Now()
	return conn.Users().Insert(u)
}

// Delete ...
func (u *User) Delete() error {
	conn, err := db.Connect()
	if err != nil {
		return err
	}
	defer conn.Close()
	return conn.Users().Remove(bson.M{"email": u.Email})
}

// Update ...
func (u *User) Update() error {
	conn, err := db.Connect()
	if err != nil {
		return err
	}
	defer conn.Close()
	return conn.Users().Update(bson.M{"email": u.Email}, u)
}

func (u *User) hashPassword(password string) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	u.Password = string(hash[:])
}
