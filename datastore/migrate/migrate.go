package migrate

import (
	"github.com/BurntSushi/migration"
)

// Setup is the database migration function that
// will setup the initial SQL database structure.
func Setup(tx migration.LimitedTx) error {
	var stmts = []string{
		userTable,
	}
	for _, stmt := range stmts {
		_, err := tx.Exec(stmt)
		if err != nil {
			return err
		}
	}
	return nil
}

var userTable = `
CREATE TABLE IF NOT EXISTS users (
	 user_id           INTEGER PRIMARY KEY AUTO_INCREMENT
  ,user_admin        BOOLEAN
	,user_username     VARCHAR(255)
	,user_name         VARCHAR(255)
	,user_email        VARCHAR(255)
  ,user_password     VARCHAR(255)
	,user_token        VARCHAR(255)
	,user_active       BOOLEAN
	,user_created      INTEGER
	,user_updated      INTEGER
	,UNIQUE(user_token)
	,UNIQUE(user_username, user_email)
);
`
