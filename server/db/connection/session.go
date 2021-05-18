package connection

import (
	"errors"
	"log"
	"time"

	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/mysql"
)

var session db.Session

// OpenSession open database session
// There is only one db session at the same time
// Returns session, nil on success
func OpenSession(db_settings *mysql.ConnectionURL) (db.Session, error) {
	if session != nil {
		return session, errors.New("Session already exists")
	}
	if db_settings.Options == nil {
		db_settings.Options = make(map[string]string)
	}
	db_settings.Options["multiStatements"] = "true"
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
		return nil, errors.New("Session not created")
	}
	return session, nil
}

func GetSession() db.Session {
	return session
}
