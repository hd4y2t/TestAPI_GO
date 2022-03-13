package book

import "time"

type Book struct {
	ID         int
	Title      string
	Desciption string
	Price      int
	Rating     int
	Discoud    int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
