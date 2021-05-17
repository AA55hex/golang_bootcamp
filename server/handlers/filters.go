package handlers

import (
	"strconv"

	"github.com/AA55hex/golang_bootcamp/server/db/connection"
	"github.com/AA55hex/golang_bootcamp/server/db/entity"
)

type PriceFilter struct {
	minPrice *float32
	maxPrice *float32
}
type BookFilter struct {
	Name  string
	Price PriceFilter
	Genre *int32
}

func (f *BookFilter) Parse(filters map[string]string) {
	//get name
	f.Name = filters["name"]
	f.Price.Parse(filters)

	genre, err := strconv.ParseInt(filters["minPrice"], 10, 32)
	if err == nil {
		buff := int32(genre)
		f.Genre = &buff
	}
}

func (p *PriceFilter) Parse(filters map[string]string) {
	minPrice, err := strconv.ParseFloat(filters["minPrice"], 32)
	if err == nil {
		buff := float32(minPrice)
		p.minPrice = &buff
	}

	maxPrice, err := strconv.ParseFloat(filters["maxPrice"], 32)
	if err == nil {
		buff := float32(maxPrice)
		p.maxPrice = &buff
	}
}

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
		query = next_func("genre = ?", filter.Genre)
		next_func = query.And
	}

	// prcie filtering
	switch {
	case filter.Price.minPrice != nil && filter.Price.maxPrice != nil:
		query = next_func("genre between ? and ?",
			filter.Price.minPrice,
			filter.Price.maxPrice)
		next_func = query.And
	case filter.Price.minPrice != nil && filter.Price.maxPrice == nil:
		query = next_func("genre > ", filter.Price.minPrice)
		next_func = query.And
	case filter.Price.minPrice == nil && filter.Price.maxPrice != nil:
		query = next_func("genre < ", filter.Price.maxPrice)
		next_func = query.And
	}

	// check for 0 amoount
	query = next_func("amount != ", "0")

	var result []entity.Book
	err := query.All(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
