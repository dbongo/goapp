package database

import (
	"database/sql"
	"log"
	"os"

	"github.com/BurntSushi/migration"
	"github.com/dbongo/hackapp/datastore"
	"github.com/dbongo/hackapp/datastore/migrate"
	_ "github.com/go-sql-driver/mysql"
	"github.com/russross/meddler"

	"gopkg.in/mgo.v2"
)

const (
	postgres = "postgres"
	sqlite   = "sqlite3"
	mysql    = "mysql"
	users    = "users"
)

// Connect ...
func Connect(driver, datasource string) (*sql.DB, error) {
	switch driver {
	case postgres:
		meddler.Default = meddler.PostgreSQL
	case sqlite:
		meddler.Default = meddler.SQLite
	case mysql:
		meddler.Default = meddler.MySQL
	}
	migration.DefaultGetVersion = migrate.GetVersion
	migration.DefaultSetVersion = migrate.SetVersion
	migrations := []migration.Migrator{
		migrate.Setup,
	}
	return migration.Open(driver, datasource, migrations)
}

// MustConnect ...
func MustConnect(driver, datasource string) *sql.DB {
	db, err := Connect(driver, datasource)
	if err != nil {
		panic(err)
	}
	return db
}

func mustConnectTest() *sql.DB {
	var (
		driver     = os.Getenv("TEST_DRIVER")
		datasource = os.Getenv("TEST_DATASOURCE")
	)
	if len(driver) == 0 {
		driver = mysql
		datasource = "root@/test_db"
	}
	db, err := Connect(driver, datasource)
	if err != nil {
		panic(err)
	}
	return db
}

// New ...
func New(addr, name string) *mgo.Database {
	session, err := mgo.Dial(addr)
	if err != nil {
		log.Fatal(err)
	}
	s := session.Clone()
	return s.DB(name)
}

// NewDatastore ...
func NewDatastore(db *sql.DB) datastore.Datastore {
	return struct {
		*Userstore
	}{
		NewUserstore(db),
	}
}

// func NewDatastore(db *mgo.Database) datastore.Datastore {
// 	return struct {
// 		*UserCollection
// 	}{
// 		NewUserCollection(db.C(users)),
// 	}
// }
