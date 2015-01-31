package database

import (
	"testing"

	"github.com/franela/goblin"
)

const (
	testaddr = "127.0.0.1:27017"
	testname = "hackapp_testdb"
)

func TestUserCollection(t *testing.T) {
	db := New(testaddr, testname)
	uc := NewUserCollection(db.C(users))
	defer db.Session.Close()
	g := goblin.Goblin(t)
	g.Describe("UserCollection", func() {
		g.Before(func() {
			uc.DropCollection()
		})
		g.It("Should create the admin user", func() {
			data := struct {
				Email    string
				Username string
				Password string
			}{
				Email:    "admin@email.com",
				Username: "admin",
				Password: "abc123",
			}
			user, err := uc.CreateUser(data.Email, data.Username, data.Password)
			g.Assert(err == nil).IsTrue()
			g.Assert(user.ID != "").IsTrue()
			g.Assert(user.Admin == true).IsTrue()
		})
		g.It("Should create a user", func() {
			data := struct {
				Email    string
				Username string
				Password string
			}{
				Email:    "user@email.com",
				Username: "user",
				Password: "user123",
			}
			user, err := uc.CreateUser(data.Email, data.Username, data.Password)
			g.Assert(err == nil).IsTrue()
			g.Assert(user.ID != "").IsTrue()
			g.Assert(user.Admin == false).IsTrue()
		})
	})
}
