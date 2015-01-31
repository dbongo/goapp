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

// UserCollection ...
type UserCollection struct {
	*mgo.Collection
}

// NewUserCollection ...
func NewUserCollection(users *mgo.Collection) *UserCollection {
	email := mgo.Index{
		Key:        []string{"email"},
		Unique:     true,
		Background: true,
	}
	users.EnsureIndex(email)
	return &UserCollection{users}
}

// GetUser ...
func (u *UserCollection) GetUser(email string) (*model.User, error) {
	user := new(model.User)
	err := u.Find(bson.M{"email": email}).One(user)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}

// GetUserList ...
func (u *UserCollection) GetUserList() ([]*model.User, error) {
	var users []*model.User
	err := u.Find(nil).All(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// UpdateUser ...
func (u *UserCollection) UpdateUser(user *model.User) error {
	user.Updated = time.Now().Format(time.RFC3339)
	return u.Update(bson.M{"_id": user.ID}, user)
}

// AuthUser ...
func (u *UserCollection) AuthUser(email, password string) (*model.User, error) {
	user, err := u.GetUser(email)
	if err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}
	user.LastLogin = time.Now().Format(time.RFC3339)
	if err := u.UpdateUser(user); err != nil {
		return nil, err
	}
	return user, nil
}

// CreateUser ...
func (u *UserCollection) CreateUser(email, username, password string) (*model.User, error) {
	_, err := u.GetUser(email)
	if email == "" || username == "" || password == "" {
		return nil, errors.New("email, username, password are required fields")
	} else if err == nil {
		return nil, errors.New("please provide another email, " + email + " is taken")
	}
	user := new(model.User)
	user.ID = bson.NewObjectId()
	user.Admin = isAdmin(u)
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

func isAdmin(u *UserCollection) bool {
	var users []*model.User
	users, _ = u.GetUserList()
	if users == nil || len(users) == 0 {
		return true
	}
	return false
}
