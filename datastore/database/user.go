package database

import (
	"errors"
	"log"
	"time"

	"github.com/dbongo/hackapp/model"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Userstore ...
type Userstore struct {
	*mgo.Collection
}

// NewUserstore ...
func NewUserstore(users *mgo.Collection) *Userstore {
	email := mgo.Index{
		Key:        []string{"email"},
		Unique:     true,
		Background: true,
	}
	users.EnsureIndex(email)
	return &Userstore{users}
}

// GetUser ...
func (u *Userstore) GetUser(email string) (*model.User, error) {
	user := model.User{}
	err := u.Find(bson.M{"email": email}).One(&user)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

// PutUser ...
func (u *Userstore) PutUser(user *model.User) error {
	user.Updated = time.Now().Format(time.RFC3339)
	return u.Update(bson.M{"_id": user.ID}, user)
}

// PostUserLogin ...
func (u *Userstore) PostUserLogin(email, password string) (*model.User, error) {
	user, err := u.GetUser(email)
	if err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}
	user.LastLogin = time.Now().Format(time.RFC3339)
	if err := u.PutUser(user); err != nil {
		return nil, err
	}
	return user, nil
}

// PostUserRegistration ...
func (u *Userstore) PostUserRegistration(email, username, password string) (*model.User, error) {
	_, err := u.GetUser(email)
	if email == "" || username == "" || password == "" {
		return nil, errors.New("email, username, password are required fields")
	} else if err == nil {
		return nil, errors.New("please provide another email, " + email + " is taken")
	}
	user := &model.User{}
	user.ID = bson.NewObjectId()
	user.Email = email
	user.Username = username
	user.Password = hashPassword(password)
	user.Created = time.Now().Format(time.RFC3339)
	if err := u.Insert(user); err != nil {
		return nil, err
	}
	return user, nil
}

func hashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(hash[:])
}
