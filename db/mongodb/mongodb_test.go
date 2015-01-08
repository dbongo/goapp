package mongodb

// import (
// 	"testing"
//
// 	"launchpad.net/gocheck"
// )
//
// func Test(t *testing.T) { gocheck.TestingT(t) }
//
// type S struct{}
//
// var _ = gocheck.Suite(&S{})
//
// func (s *S) SetUpSuite(c *gocheck.C) {
// 	ticker.Stop()
// }
//
// func (s *S) TearDownSuite(c *gocheck.C) {
// 	storage, err := MongoDB("127.0.0.1:27017", "appdb_test")
// 	c.Assert(err, gocheck.IsNil)
// 	defer storage.session.Close()
// 	storage.session.DB("appdb_test").DropDatabase()
// }
//
// func (s *S) TearDownTest(c *gocheck.C) {
// 	connections = make(map[string]*session)
// }
//
// func (s *S) TestMongoDBConnectsToTheDatabase(c *gocheck.C) {
// 	storage, err := MongoDB("127.0.0.1:27017", "appdb_test")
// 	c.Assert(err, gocheck.IsNil)
// 	defer storage.session.Close()
// 	c.Assert(storage.session.Ping(), gocheck.IsNil)
// }
//
// func (s *S) TestMongoDBCopiesConnection(c *gocheck.C) {
// 	storage, err := MongoDB("127.0.0.1:27017", "appdb_test")
// 	c.Assert(err, gocheck.IsNil)
// 	defer storage.session.Close()
// 	storage2, err := MongoDB("127.0.0.1:27017", "appdb_test")
// 	c.Assert(err, gocheck.IsNil)
// 	c.Assert(storage.session, gocheck.Not(gocheck.Equals), storage2.session)
// }
//
// func (s *S) TestMongoDBReconnects(c *gocheck.C) {
// 	storage, err := MongoDB("127.0.0.1:27017", "appdb_test")
// 	c.Assert(err, gocheck.IsNil)
// 	storage.session.Close()
// 	storage, err = MongoDB("127.0.0.1:27017", "appdb_test")
// 	c.Assert(err, gocheck.IsNil)
// 	err = storage.session.Ping()
// 	c.Assert(err, gocheck.IsNil)
// }
//
// func (s *S) TestMongoDBConnectionRefused(c *gocheck.C) {
// 	storage, err := MongoDB("127.0.0.1:27018", "appdb_test")
// 	c.Assert(storage, gocheck.IsNil)
// 	c.Assert(err, gocheck.NotNil)
// }
//
// func (s *S) TestClose(c *gocheck.C) {
// 	defer func() {
// 		r := recover()
// 		c.Check(r, gocheck.NotNil)
// 	}()
// 	storage, err := MongoDB("127.0.0.1:27017", "appdb_test")
// 	c.Assert(err, gocheck.IsNil)
// 	storage.Close()
// 	err = storage.session.Ping()
// 	c.Check(err, gocheck.NotNil)
// }
//
// func (s *S) TestCollection(c *gocheck.C) {
// 	storage, _ := MongoDB("127.0.0.1:27017", "appdb_test")
// 	defer storage.session.Close()
// 	collection := storage.Collection("users")
// 	c.Assert(collection.FullName, gocheck.Equals, storage.name+".users")
// }
