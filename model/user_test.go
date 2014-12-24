package model

import (
	"os"
	"testing"

	"github.com/dbongo/goapp/db"

	"gopkg.in/mgo.v2"
	"launchpad.net/gocheck"
)

func Test(t *testing.T) { gocheck.TestingT(t) }

type S struct{}

var _ = gocheck.Suite(&S{})

func (s *S) SetUpSuite(c *gocheck.C) {
	os.Setenv("MONGO_ADDRESS", "127.0.0.1:27017")
	os.Setenv("MONGO_DATABASE", "appdb_test")
}

func (s *S) TearDownSuite(c *gocheck.C) {
	strg, err := db.Connect()
	c.Assert(err, gocheck.IsNil)
	defer strg.Close()
}

func (s *S) TestCreateUser(c *gocheck.C) {
	user := User{Name: "Alice", Email: "alice@example.org", Username: "alice", Password: "123456"}
	defer user.Delete()
	err := user.Save()
	c.Assert(err, gocheck.IsNil)
}

func (s *S) TestCreateUserWithSameEmail(c *gocheck.C) {
	user := &User{Name: "Alice", Email: "alice@example.org", Username: "alice", Password: "123456"}
	user.Save()
	defer user.Delete()

	user2 := &User{Name: "Bob", Email: "alice@example.org", Username: "bob", Password: "123456"}
	err := user2.Save()
	c.Assert(mgo.IsDup(err), gocheck.Equals, true)
}

func (s *S) TestCreateUserWithoutRequiredFields(c *gocheck.C) {
	user := User{}
	err := user.Save()
	c.Assert(err.Error(), gocheck.Equals, "email / username / password fields can not be blank")
}

func (s *S) TestValid(c *gocheck.C) {
	user := User{Name: "Alice", Email: "alice@example.org", Username: "alice", Password: "123456"}
	defer user.Delete()
	user.Save()
	valid := user.Valid()
	c.Assert(valid, gocheck.Equals, true)
}

func (s *S) TestValidWhenUserDoesNotExistInTheDB(c *gocheck.C) {
	user := User{Name: "Alice", Email: "alice@example.org", Username: "alice", Password: "123456"}
	valid := user.Valid()
	c.Assert(valid, gocheck.Equals, false)
}

func (s *S) TestFindUserByEmail(c *gocheck.C) {
	user := User{Name: "Alice", Email: "alice@example.org", Username: "alice", Password: "123456"}
	defer user.Delete()
	user.Save()
	found, err := FindUserByEmail(user.Email)
	c.Assert(err, gocheck.IsNil)
	c.Assert(found, gocheck.NotNil)
}

func (s *S) TestFindUserWithInvalidEmail(c *gocheck.C) {
	user := User{Name: "Alice", Email: "alice@example.org", Username: "alice", Password: "123456"}
	defer user.Delete()
	user.Save()
	_, err := FindUserByEmail("bob@example.org")
	c.Assert(err, gocheck.Equals, mgo.ErrNotFound)
}
