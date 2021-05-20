package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/AA55hex/golang_bootcamp/server/db/connection"
	"github.com/AA55hex/golang_bootcamp/server/db/entity"
	"github.com/gorilla/mux"
)

func jsonResponse(w http.ResponseWriter, httpStatus int, jsonBody interface{}) error {
	w.WriteHeader(httpStatus)
	w.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(jsonBody)
	return err
}

func textResponse(w http.ResponseWriter, httpStatus int, body []byte) {
	w.WriteHeader(httpStatus)
	w.Write(body)
}

// GetBookByIDHandler is http handler for GET /books/{id:[0-9]+}
var GetBookByIDHandler = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	book_id, _ := strconv.ParseInt(vars["id"], 10, 32)
	book, _ := entity.GetBook(int32(book_id), connection.GetSession())
	if book == nil {
		textResponse(w, http.StatusNotFound, nil)
		return
	}

	jsonResponse(w, http.StatusFound, book)
}

// GetBooksByFilterHandler is http handler for GET /books with filtres
var GetBooksByFilterHandler = func(w http.ResponseWriter, r *http.Request) {
	// Getting and parsing filters
	filters := FilterMap{}
	var filter BookFilter
	filters["name"] = r.URL.Query().Get("name")
	filters["minPrice"] = r.URL.Query().Get("minPrice")
	filters["maxPrice"] = r.URL.Query().Get("maxPrice")
	filters["genre"] = r.URL.Query().Get("genre")

	err := filter.Parse(filters)
	if err != nil {
		textResponse(w, http.StatusBadRequest, []byte(err.Error()))
		return
	}

	books, err := GetBooks(&filter)
	if err != nil {
		textResponse(w, http.StatusInternalServerError, []byte("Server error"))
		return
	}

	jsonResponse(w, http.StatusFound, books)

}

// CreateBookHandler is http handler for POST /books/new
var CreateBookHandler = func(w http.ResponseWriter, r *http.Request) {
	book := &entity.Book{}
	err := json.NewDecoder(r.Body).Decode(book)
	if err != nil {
		textResponse(w, http.StatusBadRequest, []byte(err.Error()))
		return
	}

	err = book.Insert(connection.GetSession())
	if err != nil {
		textResponse(w, http.StatusBadRequest, []byte(err.Error()))
		return
	}

	id := []byte(strconv.FormatInt(int64(book.Id), 10))
	textResponse(w, http.StatusCreated, id)
}

// UpdateBookHandler is http handler for PUT /books/{id:[0-9]+}
var UpdateBookHandler = func(w http.ResponseWriter, r *http.Request) {
	book := &entity.Book{}
	err := json.NewDecoder(r.Body).Decode(book)
	if err != nil {
		textResponse(w, http.StatusBadRequest, []byte(err.Error()))
		return
	}

	vars := mux.Vars(r)
	book_id, _ := strconv.ParseInt(vars["id"], 10, 32)
	book.Id = int32(book_id)

	err = book.Update(connection.GetSession())
	if err != nil {
		textResponse(w, http.StatusBadRequest, []byte(err.Error()))
		return
	}

	id := []byte(strconv.FormatInt(int64(book.Id), 10))
	textResponse(w, http.StatusOK, id)
}

// DeleteBookHandler is http handler for DELETE /books/{id:[0-9]+}
var DeleteBookHandler = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	book_id, _ := strconv.ParseInt(vars["id"], 10, 32)
	err := entity.DeleteBook(int32(book_id), connection.GetSession())
	if err != nil {
		textResponse(w, http.StatusNotFound, []byte(err.Error()))
	}
	/// http: superfluous response.WriteHeader
	// w.WriteHeader(http.StatusNoContent)
}
