package book

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Book, error)
	Create(book Book) (Book, error)
	FindByID(id int) (Book, error)
	Update(book Book) (Book, error)
	Delete(book Book) (Book, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Book, error) {
	var book []Book
	err := r.db.Find(&book).Error
	// if err != nil {

	// }
	return book, err
}

func (r *repository) FindByID(ID int) (Book, error) {
	var book Book
	err := r.db.Find(&book, ID).Error
	// if err != nil {

	// }
	return book, err
}

func (r *repository) Create(book Book) (Book, error) {
	// var item book.Book
	err := r.db.Create(&book).Error

	// item.Title = "hanya coba"
	// db.Save(&item)
	return book, err
}

func (r *repository) Update(book Book) (Book, error) {
	err := r.db.Save(&book).Error
	return book, err
}

func (r *repository) Delete(book Book) (Book, error) {
	err := r.db.Delete(&book).Error
	return book, err
}
