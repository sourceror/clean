package books_test

import (
	"fmt"
	"testing"

	"github.com/sourceror/clean/books"
)

type newBookInput struct {
	params books.NewBookParameters
}

type newBookOutput struct {
	book *books.Book
	err  error
}

func TestNewBook(t *testing.T) {
	tests := map[string]struct {
		in   newBookInput
		want newBookOutput
	}{
		"assert error: empty Title": {
			in: newBookInput{
				params: books.NewBookParameters{
					Title:       "",
					Description: "The Night Watch has a Dragon on their hands...",
					AuthorID:    "55b16d39-17e9-457d-8c31-54f55c81563a",
				},
			},
			want: newBookOutput{
				book: nil,
				err:  fmt.Errorf("%w: empty Title", books.ErrBookValidation),
			},
		},
		"assert error: empty Description": {
			in: newBookInput{
				params: books.NewBookParameters{
					Title:       "Guards Guards!",
					Description: "",
					AuthorID:    "55b16d39-17e9-457d-8c31-54f55c81563a",
				},
			},
			want: newBookOutput{
				book: nil,
				err:  fmt.Errorf("%w: empty Description", books.ErrBookValidation),
			},
		},
		"assert error: empty AuthorID": {
			in: newBookInput{
				params: books.NewBookParameters{
					Title:       "Guards Guards!",
					Description: "The Night Watch has a Dragon on their hands...",
					AuthorID:    "",
				},
			},
			want: newBookOutput{
				book: nil,
				err:  fmt.Errorf("%w: empty AuthorID", books.ErrBookValidation),
			},
		},
		"book is created correctly": {
			in: newBookInput{
				params: books.NewBookParameters{
					Title:       "Guards Guards!",
					Description: "The Night Watch has a Dragon on their hands...",
					AuthorID:    "55b16d39-17e9-457d-8c31-54f55c81563a",
				},
			},
			want: newBookOutput{
				book: &books.Book{
					Title:       "Guards Guards!",
					Description: "The Night Watch has a Dragon on their hands...",
					AuthorID:    "55b16d39-17e9-457d-8c31-54f55c81563a",
				},
				err: nil,
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			book, err := books.NewBook(tt.in.params)
			if err != nil {
			}
			assertNewBookOutput(t, newBookOutput{
				book: book,
				err:  err,
			}, tt.want)
		})
	}
}

func assertNewBookOutput(t *testing.T, got, want newBookOutput) {
	t.Helper()
	if got.err != want.err {
		t.Fatalf("mismatched error contents, got %s, want %s", got.err, want.err)
	}

	if got.book.Title != want.book.Title {
		t.Fatalf("mismatched Title, got %s, want %s", got.book.Title, want.book.Title)
	}
	if got.book.Description != want.book.Description {
		t.Fatalf("mismatched Description, got %s, want %s", got.book.Description, want.book.Description)
	}
	if got.book.AuthorID != want.book.AuthorID {
		t.Fatalf("mismatched Author, got %s, want %s", got.book.AuthorID, want.book.AuthorID)
	}
}
