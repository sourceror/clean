package rpc

import (
	"errors"

	"github.com/sourceror/clean/books"
)

var (
	ErrNilNewBook = errors.New("nil NewBook use case")
)

type NewBook interface {
	NewBook(parameters books.NewBookParameters) error
}

func createNewBookHandler(useCase NewBook) (*newBookHandler, error) {
	if useCase == nil {
		return nil, ErrNilNewBook
	}
	return &newBookHandler{useCase: useCase}, nil
}

type newBookHandler struct {
	useCase NewBook
}
