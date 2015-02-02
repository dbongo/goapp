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

	test := goblin.Goblin(t)

	test.Describe("UserCollection", func() {

		test.Before(func() {
			uc.DropCollection()
		})

		test.It("Should create the admin user", func() {
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
			test.Assert(err == nil).IsTrue()
			test.Assert(user.ID != "").IsTrue()
			test.Assert(user.Admin == true).IsTrue()
		})

		test.It("Should create a user", func() {
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
			test.Assert(err == nil).IsTrue()
			test.Assert(user.ID != "").IsTrue()
			test.Assert(user.Admin == false).IsTrue()
		})
	})
}
