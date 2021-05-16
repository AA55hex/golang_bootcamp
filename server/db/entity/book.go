package entity

import (
	"errors"

	"github.com/AA55hex/golang_bootcamp/server/db/connection"
	"github.com/upper/db/v4"
)

type Book struct {
	Id     int32   `json,db:"id,omitempty"`
	Name   string  `json:"name"`
	Price  float32 `json:"price"`
	Genre  int     `json:"genre"`
	Amount int     `json:"amount"`
}

// Validate book for insertion to database
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
	switch {
	case b.Price < 0:
		return errors.New("Bad price")
	case b.Amount < 0:
		return errors.New("Bad amount")
	default:
		return nil
	}
}

// Validate and trying to insert book object in database
// If the operation succeeds, updates current
// object with data from the newly inserted row
// Returns nil on success
func (b *Book) Insert() error {

	if err := b.Validate(); err != nil {
		return errors.New("Insert validation failed: " + err.Error())
	}

	books := connection.GetSession().Collection("book")
	err := books.InsertReturning(b)
	if err != nil {
		return errors.New("Insertion failed: " + err.Error())
	}

	return nil
}

// Validate and trying to update book object in database
// Returns nil on success
func (b *Book) Update() error {
	books := connection.GetSession().Collection("book")
	genres := connection.GetSession().Collection("genre")

	// check for existence
	res := books.Find(db.Cond{"id": b.Id})
	if count, _ := res.Count(); count != 1 {
		return errors.New("Update validate failed: Book not found")
	}

	// get limited name for comparsion
	var limited_name string
	if len(b.Name) > 100 {
		limited_name = b.Name[0:100]
	} else {
		limited_name = b.Name[0:len(b.Name)]
	}

	// check for name duplications
	name_duplications, _ := books.
		Find(db.Cond{"name": limited_name, "id !=": b.Id}).
		Count()
	if name_duplications != 0 {
		return errors.New("Update validate failed: Name is not unique")
	}

	// check for genre existence
	genre_existence, _ := genres.Find(db.Cond{"id": b.Genre}).Count()
	if genre_existence != 0 {
		return errors.New("Update validate failed: Bad genre id")
	}

	// simple validations
	switch {
	case b.Price < 0:
		return errors.New("Update validate failed: Bad price")
	case b.Amount < 0:
		return errors.New("Update validate failed: Bad amount")
	}

	// try to update
	err := res.Update(b)
	if err != nil {
		return errors.New("Update failed: " + err.Error())
	}
	return nil
}

// Validate and trying to delete book object from database
// Returns nil on success
func (b *Book) Delete() error {
	books := connection.GetSession().Collection("book")

	// check for existence
	res := books.Find(db.Cond{"id": b.Id})
	if count, _ := res.Count(); count != 1 {
		return errors.New("Delete validation failed: Book not found")
	}

	// try to delete
	if err := res.Delete(); err != nil {
		return errors.New("Deleting failed: Book not found")
	}

	return nil
}
