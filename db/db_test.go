package db

import (
	"os"
	"reflect"
	"testing"

	"github.com/dbongo/goapp/db/store"

	"launchpad.net/gocheck"
)

type hasIndexChecker struct{}
type hasUniqueIndexChecker struct{}

var HasIndex gocheck.Checker = &hasIndexChecker{}
var HasUniqueIndex gocheck.Checker = &hasUniqueIndexChecker{}

func (c *hasIndexChecker) Info() *gocheck.CheckerInfo {
	return &gocheck.CheckerInfo{Name: "HasIndexChecker", Params: []string{"collection", "key"}}
}

func (c *hasUniqueIndexChecker) Info() *gocheck.CheckerInfo {
	return &gocheck.CheckerInfo{Name: "HasUniqueField", Params: []string{"collection", "key"}}
}

func (c *hasIndexChecker) Check(params []interface{}, names []string) (bool, string) {
	collection, ok := params[0].(*store.Collection)
	if !ok {
		return false, "first parameter should be a Collection"
	}
	key, ok := params[1].([]string)
	if !ok {
		return false, "second parameter should be the key, as used for mgo index declaration (slice of strings)"
	}
	indexes, err := collection.Indexes()
	if err != nil {
		return false, "failed to get collection indexes: " + err.Error()
	}
	for _, index := range indexes {
		if reflect.DeepEqual(index.Key, key) {
			return true, ""
		}
	}
	return false, ""
}

func (c *hasUniqueIndexChecker) Check(params []interface{}, names []string) (bool, string) {
	collection, ok := params[0].(*store.Collection)
	if !ok {
		return false, "first parameter should be a Collection"
	}
	key, ok := params[1].([]string)
	if !ok {
		return false, "second parameter should be the key, as used for mgo index declaration (slice of strings)"
	}
	indexes, err := collection.Indexes()
	if err != nil {
		return false, "failed to get collection indexes: " + err.Error()
	}
	for _, index := range indexes {
		if reflect.DeepEqual(index.Key, key) {
			return index.Unique, ""
		}
	}
	return false, ""
}

func Test(t *testing.T) { gocheck.TestingT(t) }

type S struct{}

var _ = gocheck.Suite(&S{})

func (s *S) SetUpSuite(c *gocheck.C) {
	os.Setenv("MONGODB_ADDRESS", "127.0.0.1:27017")
	os.Setenv("MONGODB_NAME", "testdb")
}

func (s *S) TearDownSuite(c *gocheck.C) {
	store, err := Connect()
	c.Assert(err, gocheck.IsNil)
	defer store.Close()
	store.Collection("apps").Database.DropDatabase()
}

func (s *S) TestUsers(c *gocheck.C) {
	store, err := Connect()
	c.Assert(err, gocheck.IsNil)
	users := store.Users()
	usersc := store.Collection("users")
	c.Assert(users, gocheck.DeepEquals, usersc)
	c.Assert(users, HasUniqueIndex, []string{"email"})
}
