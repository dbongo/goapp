package store

import (
	"testing"

	"launchpad.net/gocheck"
)

func Test(t *testing.T) { gocheck.TestingT(t) }

type S struct{}

var _ = gocheck.Suite(&S{})

func (s *S) SetUpSuite(c *gocheck.C) {
	ticker.Stop()
}

func (s *S) TearDownSuite(c *gocheck.C) {
	db, err := Open("127.0.0.1:27017", "testdb")
	c.Assert(err, gocheck.IsNil)
	defer db.session.Close()
	db.session.DB("testdb").DropDatabase()
}

func (s *S) TearDownTest(c *gocheck.C) {
	connections = make(map[string]*session)
}

func (s *S) TestOpenConnectsToTheDatabase(c *gocheck.C) {
	db, err := Open("127.0.0.1:27017", "testdb")
	c.Assert(err, gocheck.IsNil)
	defer db.session.Close()
	c.Assert(db.session.Ping(), gocheck.IsNil)
}

func (s *S) TestOpenCopiesConnection(c *gocheck.C) {
	db, err := Open("127.0.0.1:27017", "testdb")
	c.Assert(err, gocheck.IsNil)
	defer db.session.Close()
	db2, err := Open("127.0.0.1:27017", "testdb")
	c.Assert(err, gocheck.IsNil)
	c.Assert(db.session, gocheck.Not(gocheck.Equals), db2.session)
}

func (s *S) TestOpenReconnects(c *gocheck.C) {
	db, err := Open("127.0.0.1:27017", "testdb")
	c.Assert(err, gocheck.IsNil)
	db.session.Close()
	db, err = Open("127.0.0.1:27017", "testdb")
	c.Assert(err, gocheck.IsNil)
	err = db.session.Ping()
	c.Assert(err, gocheck.IsNil)
}

func (s *S) TestOpenConnectionRefused(c *gocheck.C) {
	db, err := Open("127.0.0.1:27018", "testdb")
	c.Assert(db, gocheck.IsNil)
	c.Assert(err, gocheck.NotNil)
}

func (s *S) TestClose(c *gocheck.C) {
	defer func() {
		r := recover()
		c.Check(r, gocheck.NotNil)
	}()
	db, err := Open("127.0.0.1:27017", "testdb")
	c.Assert(err, gocheck.IsNil)
	db.Close()
	err = db.session.Ping()
	c.Check(err, gocheck.NotNil)
}

func (s *S) TestCollection(c *gocheck.C) {
	db, _ := Open("127.0.0.1:27017", "testdb")
	defer db.session.Close()
	collection := db.Collection("users")
	c.Assert(collection.FullName, gocheck.Equals, db.name+".users")
}
