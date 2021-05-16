package session

import (
	"errors"
	"fmt"

	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/mysql"
)

var db_settings mysql.ConnectionURL

func Create() (db.Session, error) {
	// configure connection
	db_settings = mysql.ConnectionURL{
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
	session, err := mysql.Open(db_settings)
	if err != nil {
		return nil, errors.New("session.Create: " + err.Error())
	}
	fmt.Println("Session created")
	return session, nil
}
