package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/AA55hex/golang_bootcamp/server/config"
	"github.com/AA55hex/golang_bootcamp/server/db/connection"
	"github.com/AA55hex/golang_bootcamp/server/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// check database session
	if connection.GetSession() == nil {
		log.Fatal("Database session not created")
	}
	defer connection.GetSession().Close()

	// init router
	fmt.Println("Creating router")
	router := mux.NewRouter()
	router.HandleFunc("/books/{id:[0-9]+}", handlers.GetBookByIDHandler).Methods("GET")
	router.HandleFunc("/books/{id:[0-9]+}", handlers.UpdateBookHandler).Methods("PUT")
	router.HandleFunc("/books/{id:[0-9]+}", handlers.DeleteBookHandler).Methods("DELETE")
	router.HandleFunc("/books", handlers.GetBooksByFilterHandler).Methods("GET")
	router.HandleFunc("/books/new", handlers.CreateBookHandler).Methods("POST")
	// listen & serve
	fmt.Println("Creating server")
	server := http.Server{
		Handler:      router,
		Addr:         config.Server.Address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("Server created.")
	fmt.Println("Listening started on", server.Addr)
	log.Fatal(server.ListenAndServe())
}
