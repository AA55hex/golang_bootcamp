package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/AA55hex/golang_bootcamp/server/db/entity"
	"github.com/gorilla/mux"
)

// Http handler for GET /books/{id:[0-9]+}
var GetBookByIdHandler = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	book_id, _ := strconv.ParseInt(vars["id"], 10, 32)
	book, _ := entity.GetBook(int32(book_id))
	if book == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 page not found"))
		return
	}

	w.WriteHeader(http.StatusFound)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

var GetBooksByFilterHandler = func(w http.ResponseWriter, r *http.Request) {
	// Getting and parsing filters
	filters := FilterMap{}
	var filter BookFilter
	filters["name"] = r.URL.Query().Get("name")
	filters["minPrice"] = r.URL.Query().Get("minPrice")
	filters["maxPrice"] = r.URL.Query().Get("maxPrice")
	filters["genre"] = r.URL.Query().Get("genre")
	filter.Parse(filters)
	books, err := GetBooks(&filter)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusFound)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)

}

// Http handler for POST /books/new
var CreateBookHandler = func(w http.ResponseWriter, r *http.Request) {
	book := &entity.Book{}
	err := json.NewDecoder(r.Body).Decode(book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = book.Insert()
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
	id := []byte(strconv.FormatInt(int64(book.Id), 10))
	w.Write([]byte(id))
}
