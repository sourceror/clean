package database

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/sourceror/clean/books"
)

var (
	ErrNilDatabase  = errors.New("nil database object")
	ErrBadStatement = errors.New("bad statement")
)

const (
	storeBook = "STORE_BOOK"
)

const (
	bookTable = "books.book"
)

func NewAssetRepository(db *sql.DB) (*BookRepository, error) {
	if db == nil {
		return nil, ErrNilDatabase
	}
	repo := &BookRepository{
		db:            db,
		sqlStatements: make(map[string]*sql.Stmt),
	}
	for name, query := range repo.statements() {
		stmt, err := db.Prepare(query)
		if err != nil {
			return nil, fmt.Errorf("%w: %s", ErrBadStatement, err.Error())
		}
		repo.sqlStatements[name] = stmt
	}
	return repo, nil
}

type BookRepository struct {
	db            *sql.DB
	sqlStatements map[string]*sql.Stmt
}

func (b *BookRepository) statements() map[string]string {
	return map[string]string{
		storeBook: `INSERT INTO ` + bookTable + `(title, description, author_id) VALUES ($1, $2, $3)`,
	}
}

func (b *BookRepository) StoreBook(book *books.Book) error {
	stmt, ok := b.sqlStatements[storeBook]
	if !ok {
		return fmt.Errorf("statement does not exist")
	}
	_, err := stmt.Exec(book.Title, book.Description, book.AuthorID)
	return err
}
