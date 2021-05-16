package main

import (
	"github.com/AA55hex/golang_bootcamp/server/db/connection"
)

func main() {
	// create database session
	if connection.GetSession() == nil {
		return
	}
	// create router

	// listen & serve
}
