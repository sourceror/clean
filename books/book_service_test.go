package books_test

import (
	"testing"

	"github.com/sourceror/clean/books"
)

type newBookServiceInput struct {
	bookRepository books.BookRepository
}

type newBookServiceOutput struct {
	err error
}

func TestNewBookService(t *testing.T) {
	tests := map[string]struct {
		in   newBookServiceInput
		want newBookServiceOutput
	}{
		"assert error: nil BookRepository": {
			in: newBookServiceInput{
				bookRepository: nil,
			},
			want: newBookServiceOutput{
				err: books.ErrNilBookRepository,
			},
		},
		"service is created correctly": {
			in: newBookServiceInput{
				bookRepository: bookRepositoryMock{},
			},
			want: newBookServiceOutput{
				err: nil,
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := books.NewBookService(tt.in.bookRepository)
			assertNewBookServiceOutput(t, newBookServiceOutput{
				err: err,
			}, tt.want)
		})
	}
}

func assertNewBookServiceOutput(t *testing.T, got, want newBookServiceOutput) {
	t.Helper()
	if got.err != want.err {
		t.Fatalf("mismatched error contents, got %s, want %s", got.err, want.err)
	}
}

type bookRepositoryMock struct {
	storeBookError error
}

func (b bookRepositoryMock) StoreBook(book *books.Book) error {
	return b.storeBookError
}
