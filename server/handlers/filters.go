package handlers

import (
	"errors"
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

	err := f.Price.Parse(filters)
	if err != nil {
		return err
	}

	if filters["genre"] != "" {
		genre64, err := strconv.ParseInt(filters["genre"], 10, 32)
		if err != nil {
			return errors.New("bad genre")
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
			return errors.New("bad minPrice")
		}
		buff := float32(minPrice)
		p.minPrice = &buff
	}

	if filters["maxPrice"] != "" {
		maxPrice, err := strconv.ParseFloat(filters["maxPrice"], 32)
		if err != nil {
			return errors.New("bad maxPrice")
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

	// check for 0 amoount
	query = query.Where("amount != 0")

	// name filtering
	if filter.Name != "" {
		query = query.And("name = ?", filter.Name)
	}

	// genre filtering
	if filter.Genre != nil {
		query = query.And("genre = ?", *filter.Genre)
	}

	// prcie filtering
	switch {
	case filter.Price.minPrice != nil && filter.Price.maxPrice != nil:
		query = query.And("price between ? and ?",
			*filter.Price.minPrice,
			*filter.Price.maxPrice)
	case filter.Price.minPrice != nil && filter.Price.maxPrice == nil:
		query = query.And("price >= ?", *filter.Price.minPrice)
	case filter.Price.minPrice == nil && filter.Price.maxPrice != nil:
		query = query.And("price <= ?", *filter.Price.maxPrice)
	}

	var result []entity.Book
	err := query.All(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
