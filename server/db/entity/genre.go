package entity

type Genre struct {
	Id   int    `json,db:"id,omitempty"`
	Name string `json:"name"`
}
