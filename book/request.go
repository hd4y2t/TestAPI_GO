package book

import "encoding/json"

type BooksRequest struct {
	Title      string      `json:"title" binding:"required"`
	Price      json.Number `json:"price" binding:"required,number"`
	Discoud    json.Number `json:"discoud" binding:"required,number"`
	Rating     json.Number `json:"rating" binding:"required,number"`
	Desciption string      `json:"descriptoin" binding:"required"`
}
