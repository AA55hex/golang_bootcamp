package entity

import (
	"errors"

	"github.com/AA55hex/golang_bootcamp/server/db/connection"
	"github.com/upper/db/v4"
)

type Book struct {
	Id     int32    `json:"id" db:"id,omitempty"`
	Name   *string  `json:"name" db:"name"`
	Price  *float32 `json:"price" db:"price"`
	Genre  *int     `json:"genre" db:"genre"`
	Amount *int     `json:"amount" db:"amount"`
}

// Full validate book for insertion to db
// with db requests
// Returns nil on success
func (b *Book) Validate() error {
	// is not unique
	books := connection.GetSession().Collection("book")
	genres := connection.GetSession().Collection("genre")

	// check for name duplications
	name_duplications, _ := books.Find(db.Cond{"name": b.Name}).Count()
	if name_duplications != 0 {
		return errors.New("Name is not unique")
	}

	// check for genre existence
	genre_existence, _ := genres.Find(db.Cond{"id": b.Genre}).Count()
	if genre_existence != 0 {
		return errors.New("Bad genre id")
	}

	// simple validations
	err := b.SimpleValidate()
	return err
}

// Validation without db requests
// Returns nil on success
func (b *Book) SimpleValidate() error {
	switch {
	case b.Price != nil && *b.Price < 0:
		return errors.New("Bad price")
	case b.Amount != nil && *b.Amount < 0:
		return errors.New("Bad amount")
	default:
		return nil
	}
}

// Perform simple validation and trying to insert book object in database
// If the operation succeeds, updates current
// object with data from the newly inserted row
// Returns nil on success
func (b *Book) Insert() error {

	if err := b.SimpleValidate(); err != nil {
		return errors.New("Insert simple validation failed: " + err.Error())
	}

	books := connection.GetSession().Collection("book")
	err := books.InsertReturning(b)
	if err != nil {
		return errors.New("Insertion failed: " + err.Error())
	}

	return nil
}

// Perform simple validation and trying
// to update book object in database
// Returns nil on success
func (b *Book) Update() error {
	books := connection.GetSession().Collection("book")

	res := books.Find(db.Cond{"id": b.Id})
	defer res.Close()

	// check for existence
	if count, _ := res.Count(); count != 1 {
		return errors.New("Update validate failed: Book not found")
	}

	// simple validations
	err := b.SimpleValidate()
	if err != nil {
		return errors.New("Update simple validate failed: " + err.Error())
	}

	// try to update
	err = res.Update(b)
	if err != nil {
		return errors.New("Update failed: " + err.Error())
	}
	return nil
}

// Validate and trying to delete book object from database
// Returns nil on success
func (b *Book) Delete() error {
	books := connection.GetSession().Collection("book")

	res := books.Find(db.Cond{"id": b.Id})
	defer res.Close()

	// check for existence
	if count, _ := res.Count(); count != 1 {
		return errors.New("Delete validation failed: Book not found")
	}

	// try to delete
	if err := res.Delete(); err != nil {
		return errors.New("Deleting failed: Book not found")
	}

	return nil
}

// Finds book by id
// Returns book, nil on success
func GetBook(book_id int32) (*Book, error) {
	books := connection.GetSession().Collection("book")
	result := &Book{}

	err := books.Find(db.Cond{"id": book_id}).One(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
