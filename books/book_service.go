package books

import "errors"

var (
	ErrNilBookRepository = errors.New("nil BookRepository")
)

func NewBookService(bookRepository BookRepository) (*BookService, error) {
	if bookRepository == nil {
		return nil, ErrNilBookRepository
	}
	return &BookService{bookRepository: bookRepository}, nil
}

type BookService struct {
	bookRepository BookRepository
}

func (b BookService) NewBook(parameters NewBookParameters) error {
	book, err := NewBook(parameters)
	if err != nil {
		return err
	}
	return b.bookRepository.StoreBook(book)
}

type BookRepository interface {
	StoreBook(book *Book) error
}
