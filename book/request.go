package book

import "encoding/json"

type BooksRequest struct {
	Title       string      `json:"title" validate:"required"`
	Price       json.Number `json:"price" validate:"required,number"`
	Discount    json.Number `json:"discount" validate:"required,number"`
	Rating      json.Number `json:"rating" validate:"required,number"`
	Description string      `json:"description" validate:"required"`
}
