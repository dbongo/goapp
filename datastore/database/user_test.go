package database

import (
	"fmt"
	"testing"

	"github.com/dbongo/hackapp/model"
	"github.com/franela/goblin"
)

const (
	testaddr = "127.0.0.1:27017"
	testname = "hackapp_testdb"
)

func TestUserstore(t *testing.T) {

	db := mustConnectTest()
	us := NewUserstore(db)
	defer db.Close()

	g := goblin.Goblin(t)

	g.Describe("Userstore", func() {

		// before each test be sure to purge the package table data from the database.
		g.BeforeEach(func() {
			db.Exec("DELETE FROM users")
		})

		g.It("Should Post a User", func() {
			user := model.User{
				Username: "joe",
				Name:     "Joe Sixpack",
				Email:    "foo@bar.com",
				Password: "password",
				Token:    "e42080dddf012c718e476da161d21ad5",
			}
			err := us.PostUser(&user)
			g.Assert(err == nil).IsTrue()
			g.Assert(user.ID != 0).IsTrue()
			fmt.Printf("%v\n", user)
		})

		g.It("Should Put a User", func() {
			user := model.User{
				Username: "joe",
				Name:     "Joe Sixpack",
				Email:    "foo@bar.com",
				Password: "password",
				Token:    "e42080dddf012c718e476da161d21ad5",
			}
			err := us.PostUser(&user)
			g.Assert(err == nil).IsTrue()
			g.Assert(user.ID != 0).IsTrue()

			u, err := us.GetUser(user.ID)
			g.Assert(err == nil).IsTrue()
			g.Assert(u.Email == "foo@bar.com")

			u.Email = "bar@foo.com"

			err = us.PutUser(u)
			g.Assert(user.Email == "bar@foo.com")
			fmt.Printf("%v\n", user)
		})

		g.It("Should Get a User", func() {
			user := model.User{
				Username: "joe",
				Name:     "Joe Sixpack",
				Email:    "foo@bar.com",
				Password: "password",
				Token:    "e42080dddf012c718e476da161d21ad5",
				Active:   true,
				Admin:    true,
				Created:  1398065343,
				Updated:  1398065344,
			}
			us.PostUser(&user)
			getuser, err := us.GetUser(user.ID)
			g.Assert(err == nil).IsTrue()
			g.Assert(user.ID).Equal(getuser.ID)
			g.Assert(user.Username).Equal(getuser.Username)
			g.Assert(user.Name).Equal(getuser.Name)
			g.Assert(user.Email).Equal(getuser.Email)
			g.Assert(user.Token).Equal(getuser.Token)
			g.Assert(user.Active).Equal(getuser.Active)
			g.Assert(user.Admin).Equal(getuser.Admin)
			g.Assert(user.Created).Equal(getuser.Created)
			g.Assert(user.Updated).Equal(getuser.Updated)
		})

		g.It("Should Get a User List", func() {
			user1 := model.User{
				Username: "jane",
				Name:     "Jane Doe",
				Email:    "foo@bar.com",
				Password: "password",
				Token:    "ab20g0ddaf012c744e136da16aa21ad9",
			}
			user2 := model.User{
				Username: "joe",
				Name:     "Joe Sixpack",
				Email:    "foo@bar.com",
				Password: "password",
				Token:    "e42080dddf012c718e476da161d21ad5",
			}
			us.PostUser(&user1)
			us.PostUser(&user2)
			users, err := us.GetUserList()
			g.Assert(err == nil).IsTrue()
			g.Assert(len(users)).Equal(2)
			g.Assert(users[0].Username).Equal(user1.Username)
			g.Assert(users[0].Name).Equal(user1.Name)
			g.Assert(users[0].Email).Equal(user1.Email)
			g.Assert(users[0].Token).Equal(user1.Token)
		})

		//g.It("Should Get a User", func() {
		//user := model.User{
		// Login:    "joe",
		// Remote:   "github.com",
		// Access:   "f0b461ca586c27872b43a0685cbc2847",
		// Secret:   "976f22a5eef7caacb7e678d6c52f49b1",
		// Name:     "Joe Sixpack",
		// Email:    "foo@bar.com",
		// Gravatar: "b9015b0857e16ac4d94a0ffd9a0b79c8",
		// Token:    "e42080dddf012c718e476da161d21ad5",
		// Active:   true,
		// Admin:    true,
		// Created:  1398065343,
		// Updated:  1398065344,
		// Synced:   1398065345,
		//}
		//getuser, err := us.GetUser(user.ID)
		//g.Assert(err == nil).IsTrue()
		//g.Assert(user.ID).Equal(getuser.ID)
		// g.Assert(user.Login).Equal(getuser.Login)
		// g.Assert(user.Remote).Equal(getuser.Remote)
		// g.Assert(user.Access).Equal(getuser.Access)
		// g.Assert(user.Secret).Equal(getuser.Secret)
		// g.Assert(user.Name).Equal(getuser.Name)
		// g.Assert(user.Email).Equal(getuser.Email)
		// g.Assert(user.Gravatar).Equal(getuser.Gravatar)
		// g.Assert(user.Token).Equal(getuser.Token)
		// g.Assert(user.Active).Equal(getuser.Active)
		// g.Assert(user.Admin).Equal(getuser.Admin)
		// g.Assert(user.Created).Equal(getuser.Created)
		// g.Assert(user.Updated).Equal(getuser.Updated)
		// g.Assert(user.Synced).Equal(getuser.Synced)
		//})
	})
}

// func TestUserCollection(t *testing.T) {
// 	db := New(testaddr, testname)
// 	uc := NewUserCollection(db.C(users))
// 	defer db.Session.Close()
//
// 	test := goblin.Goblin(t)
//
// 	test.Describe("UserCollection", func() {
//
// 		test.Before(func() {
// 			uc.DropCollection()
// 		})
//
// 		test.It("Should create the admin user", func() {
// 			data := struct {
// 				Email    string
// 				Username string
// 				Password string
// 			}{
// 				Email:    "admin@email.com",
// 				Username: "admin",
// 				Password: "abc123",
// 			}
// 			user, err := uc.CreateUser(data.Email, data.Username, data.Password)
// 			test.Assert(err == nil).IsTrue()
// 			test.Assert(user.ID != "").IsTrue()
// 			test.Assert(user.Admin == true).IsTrue()
// 		})
//
// 		test.It("Should create a user", func() {
// 			data := struct {
// 				Email    string
// 				Username string
// 				Password string
// 			}{
// 				Email:    "user@email.com",
// 				Username: "user",
// 				Password: "user123",
// 			}
// 			user, err := uc.CreateUser(data.Email, data.Username, data.Password)
// 			test.Assert(err == nil).IsTrue()
// 			test.Assert(user.ID != "").IsTrue()
// 			test.Assert(user.Admin == false).IsTrue()
// 		})
// 	})
// }
