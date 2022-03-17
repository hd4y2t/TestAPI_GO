package book

type Service interface {
	FindAll() ([]Book, error)
	FindByID(id int) (Book, error)
	Create(bookRequest BooksRequest) (Book, error)
	Update(ID int, BookRequest BooksRequest) (Book, error)
	Delete(ID int) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Book, error) {
	books, err := s.repository.FindAll()
	return books, err
}

func (s *service) FindByID(ID int) (Book, error) {
	books, err := s.repository.FindByID(ID)
	return books, err
}

func (s *service) Create(bookRequest BooksRequest) (Book, error) {
	price, err := bookRequest.Price.Int64()
	discount, err := bookRequest.Discount.Int64()
	rating, err := bookRequest.Rating.Int64()

	book := Book{
		Title:       bookRequest.Title,
		Price:       int(price),
		Discount:    int(discount),
		Rating:      int(rating),
		Description: bookRequest.Description,
	}
	newBook, err := s.repository.Create(book)
	return newBook, err
}

func (s *service) Update(ID int, bookRequest BooksRequest) (Book, error) {
	book, err := s.repository.FindByID(ID)

	price, err := bookRequest.Price.Int64()
	discount, err := bookRequest.Discount.Int64()
	rating, err := bookRequest.Rating.Int64()

	book.Title = bookRequest.Title
	book.Price = int(price)
	book.Discount = int(discount)
	book.Rating = int(rating)
	book.Description = bookRequest.Description

	newBook, err := s.repository.Update(book)
	return newBook, err
}

func (s *service) Delete(ID int) (Book, error) {
	book, err := s.repository.FindByID(ID)

	deleteBook, err := s.repository.Delete(book)
	// books, err := s.repository.Delete(ID)
	return deleteBook, err
}
