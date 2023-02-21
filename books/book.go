package books

import (
	"errors"
	"fmt"
)

var (
	ErrBookValidation = errors.New("error validating Book")
)

type NewBookParameters struct {
	Title       string
	Description string
	AuthorID    string
}

func NewBook(params NewBookParameters) (*Book, error) {
	book := Book(params)
	err := book.validate()
	if err != nil {
		return nil, err
	}
	return &book, nil
}

// Book is the main domain object
type Book struct {
	Title       string
	Description string
	AuthorID    string
}

func (b *Book) validate() error {
	if b.Title == "" {
		return fmt.Errorf("%w: empty Title", ErrBookValidation)
	}
	if b.Description == "" {
		return fmt.Errorf("%w: empty Description", ErrBookValidation)
	}
	if b.AuthorID == "" {
		return fmt.Errorf("%w: empty AuthorID", ErrBookValidation)
	}
	return nil
}
