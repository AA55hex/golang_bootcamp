package handlers

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/AA55hex/golang_bootcamp/server/db/connection"
	"github.com/AA55hex/golang_bootcamp/server/db/entity"
)

// FilterMap is map type for filter values
// using for BookFilter parsing
type FilterMap map[string]string

type PriceFilter struct {
	minPrice *float32
	maxPrice *float32
}

// BookFilter is filter for GetBooks func
type BookFilter struct {
	Name  string
	Price PriceFilter
	Genre *int32
}

// Parse into structure filter parameters
func (f *BookFilter) Parse(filters FilterMap) error {

	f.Name = filters["name"]
	f.Price.Parse(filters)

	if filters["genre"] != "" {
		genre64, err := strconv.ParseInt(filters["genre"], 10, 32)
		if err != nil {
			return errors.New("Bad genre")
		}
		genre32 := int32(genre64)
		f.Genre = &genre32
	}
	return nil
}

// Parse into structure filter parameters
// Returns nil on success
func (p *PriceFilter) Parse(filters FilterMap) error {
	if filters["minPrice"] != "" {
		minPrice, err := strconv.ParseFloat(filters["minPrice"], 32)

		if err != nil {
			return errors.New("Bad minPrice")
		}
		buff := float32(minPrice)
		p.minPrice = &buff

	}

	if filters["minPrice"] != "" {
		maxPrice, err := strconv.ParseFloat(filters["minPrice"], 32)
		if err != nil {
			return errors.New("Bad maxPrice")
		}
		buff := float32(maxPrice)
		p.maxPrice = &buff
	}

	return nil
}

// GetBooks create sql-query for db and returns result on success
func GetBooks(filter *BookFilter) ([]entity.Book, error) {
	query := connection.GetSession().SQL().SelectFrom("book")
	// create variable to determine next query function
	next_func := query.Where

	// name filtering
	if filter.Name != "" {
		query = next_func("name = ?", filter.Name)
		next_func = query.And
	}

	// genre filtering
	if filter.Genre != nil {
		query = next_func("genre = ?", *filter.Genre)
		next_func = query.And
	}

	// prcie filtering
	switch {
	case filter.Price.minPrice != nil && filter.Price.maxPrice != nil:
		query = next_func("price between ? and ?",
			*filter.Price.minPrice,
			*filter.Price.maxPrice)
		next_func = query.And
	case filter.Price.minPrice != nil && filter.Price.maxPrice == nil:
		query = next_func("price > ?", *filter.Price.minPrice)
		next_func = query.And
	case filter.Price.minPrice == nil && filter.Price.maxPrice != nil:
		query = next_func("price < ?", *filter.Price.maxPrice)
		next_func = query.And
	}

	// check for 0 amoount
	query = next_func("amount != 0")

	var result []entity.Book
	err := query.All(&result)
	fmt.Println(query.String())
	if err != nil {
		return nil, err
	}
	return result, nil
}
