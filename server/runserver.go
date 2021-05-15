package main

import (
	"log"

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
	// ...
	// get router
	// ...
	// listen & serve
}
