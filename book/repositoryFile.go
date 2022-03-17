package book

import (
	"fmt"
)

type fileRepository struct {
}

func NewFileRepository() *fileRepository {
	return &fileRepository{}
}

func (r *fileRepository) FileAll() ([]Book, error) {
	var book []Book
	fmt.Println("file all")
	// }
	return book, nil
}

func (r *fileRepository) FindByID(ID int) (Book, error) {
	var book Book
	fmt.Println("file find by id")
	return book, nil
}

func (r *fileRepository) Create(book Book) (Book, error) {
	// var item book.Book
	fmt.Println("file create")
	return book, nil
}
