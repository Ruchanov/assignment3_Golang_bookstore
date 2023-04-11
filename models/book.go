package models

type Book struct {
	Id          uint   `json:"id" gorm: "primary key"`
	Title       string `json: "title"`
	Description string `json: "description"`
	Cost        int    `json: "cost"`
}
