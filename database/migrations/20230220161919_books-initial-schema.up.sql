BEGIN;
    CREATE SCHEMA IF NOT EXISTS books;
    CREATE TABLE books.book (
        id UUID PRIMARY KEY,
        title TEXT NOT NULL CHECK (title <> ''),
        description TEXT NOT NULL CHECK (description <> ''),
        author_id TEXT NOT NULL CHECK (author_id <> '')
    );
COMMIT;