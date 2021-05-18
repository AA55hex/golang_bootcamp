package connection

import (
	"database/sql"
	"log"
	"time"

	"github.com/AA55hex/golang_bootcamp/server/config"
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/mysql"
)

var session db.Session

// open database session & check for migrations
// todo: add .env support
func init() {
	db_settings := mysql.ConnectionURL{
		Database: config.MySQL.Database,
		Host:     config.MySQL.Host,
		User:     config.MySQL.User,
		Password: config.MySQL.Password,
		Options: map[string]string{
			"multiStatements": "true",
		},
	}
	var err error
	// open db session
	log.Println("Connection: ", db_settings)
	for i := 0; i < 30; i++ {
		log.Print("Try open session ", i, ": ")

		session, err = mysql.Open(db_settings)
		if err != nil {
			log.Println("FAIL.")
			log.Println("Error", err)
			time.Sleep(3 * time.Second)
		} else {
			log.Println("SUCCESS!")
			break
		}
	}
	if session == nil {
		return
	}
	// try migrate
	internalSQLDriver := session.Driver().(*sql.DB)
	err = try_migrate(internalSQLDriver)
	if err != nil {
		session.Close()
		session = nil
		log.Print(err)
	}
}

func GetSession() db.Session {
	return session
}
