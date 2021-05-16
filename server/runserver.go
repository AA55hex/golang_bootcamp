package main

import (
	"database/sql"
	"log"

	"github.com/AA55hex/golang_bootcamp/server/db/migrations"
	"github.com/AA55hex/golang_bootcamp/server/db/session"
)

func main() {
	// create database session
	con, err := session.Create()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer con.Close()

	// execute migration scripts
	internalSQLDriver := con.Driver().(*sql.DB)
	err = migrations.Migrate(internalSQLDriver)
	if err != nil {
		log.Fatal(err)
		return
	}
	// get router
	// ...
	// listen & serve
}
