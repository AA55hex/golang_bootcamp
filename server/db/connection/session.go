package connection

import (
	"database/sql"
	"fmt"

	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/mysql"
)

var session db.Session

// open database session & check for migrations
// todo: add .env support
func init() {
	db_settings := mysql.ConnectionURL{
		Database: `book_store`,
		Host:     `mysql_docker`,
		User:     `root`,
		Password: `pseudo_pass`,
		Options: map[string]string{
			"multiStatements": "true",
		},
	}

	// open db session
	fmt.Println("Try open session: ", db_settings)
	var err error
	session, err = mysql.Open(db_settings)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println("Session created")

	// try migrate
	internalSQLDriver := session.Driver().(*sql.DB)
	err = try_migrate(internalSQLDriver)
	if err != nil {
		fmt.Print(err)
	}
}

func GetSession() db.Session {
	return session
}
