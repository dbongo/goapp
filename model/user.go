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

var fails []Fail

type (
	// Fail ...
	Fail struct {
		Timestamp time.Time `bson:"authfail" json:"authfail,omitempty"`
		Message   string    `bson:"message" json:"message,omitempty"`
	}

	// User ...
	User struct {
		Name         string    `bson:"name" json:"name,omitempty"`
		Email        string    `bson:"email" json:"email,omitempty"`
		Username     string    `bson:"username" json:"username,omitempty"`
		Password     string    `bson:"password" json:"-"`
		Created      time.Time `bson:"created" json:"created,omitempty"`
		LastLogin    time.Time `bson:"lastlogin" json:"lastlogin,omitempty"`
		FailedLogins []Fail    `bson:"fails" json:"fails"`
	}
)

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
	u.Created = time.Now()
	if err := u.Save(); err != nil {
		return nil, err
	}
	return &u, nil
}

// AuthUser ...
func AuthUser(email, password string) (*User, error) {
	//fail := Fail{}
	u, err := FindUserByEmail(email)
	if err != nil {
		// log.Println()
		// log.Printf("string error: %v", err)
		// fail.Timestamp = time.Now()
		// fail.Message = err.Error()
		// log.Printf("struct fail: %v", fail)
		// fails = append(fails, &fail)
		// log.Printf("var fails: %v", fails)
		// log.Println()
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		// log.Println()
		// log.Printf("string error: %v", err)
		// fail.Timestamp = time.Now()
		// fail.Message = err.Error()
		// log.Printf("struct fail: %v", fail)
		// fails = append(fails, &fail)
		// log.Printf("var fails: %v", fails)
		// log.Println()
		return nil, err
	}
	u.LastLogin = time.Now()
	//u.LoginFails = &fails
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
	u.Created = time.Now()
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
	return ds.Users().Update(bson.M{"email": u.Email}, u)
}

func (u *User) hashPassword(password string) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	u.Password = string(hash[:])
}
