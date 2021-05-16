package main

import (
	"github.com/AA55hex/golang_bootcamp/server/db/connection"
)

func main() {
	// create database session
	if connection.GetSession() == nil {
		return
	}
	defer connection.GetSession().Close()
	// create router

	// listen & serve
}
