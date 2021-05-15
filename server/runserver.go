package main

import (
	"fmt"
	"log"

	"github.com/upper/db/v4/adapter/mysql"
)

var db_settings mysql.ConnectionURL

func main() {
	// create database connection
	db_settings = mysql.ConnectionURL{
		Database: `book_store`,
		Host:     `mysql_docker`,
		User:     `root`,
		Password: `pseudo_pass`,
	}
	// open db session
	fmt.Println("Open session: ", db_settings)
	session, err := mysql.Open(db_settings)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer session.Close()
	fmt.Println("Session created")
	// execute migration scripts
	// ...
	// get router
	// ...
	// listen & serve
}
