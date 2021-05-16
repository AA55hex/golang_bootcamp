package entity

type Book struct {
	Id     int32   `json,db:"id,omitempty"`
	Name   string  `json:"name"`
	Price  float32 `json:"price"`
	Genre  int     `json:"genre"`
	Amount int     `json:"amount"`
}
