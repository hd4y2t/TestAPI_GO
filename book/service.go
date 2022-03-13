package book

type Service interface {
	FindAll() ([]Book, error)
	FindByID(id int) (Book, error)
	Create(bookRequest BooksRequest) (Book, error)
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
	discount, err := bookRequest.Discoud.Int64()

	book := Book{
		Title:      bookRequest.Title,
		Price:      int(price),
		Desciption: bookRequest.Desciption,
		Discoud:    int(discount),
	}
	newBook, err := s.repository.Create(book)
	return newBook, err
}
