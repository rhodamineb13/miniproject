package repository

import (
	"context"
	"fmt"
	"miniproject/common/dto"
	"miniproject/common/helper"

	"github.com/jmoiron/sqlx"
)

type bookRepo struct {
	db *sqlx.DB
}

type BookRepo interface {
	Insert(context.Context, *dto.AddBookDTO) error
	Select(context.Context) ([]dto.ShowBooksDTO, error)
	Get(context.Context, uint) (dto.ShowBooksDTO, error)
}

func (b *bookRepo) Insert(ctx context.Context, book *dto.AddBookDTO) error {
	var bookID uint

	query := fmt.Sprintf(`SELECT id FROM books WHERE isbn = $1 AND title = $2 AND author = $3`)

	if err := b.db.GetContext(ctx, &bookID, query, book.ISBN, book.Title, book.Author); err == nil {
		return helper.ErrDuplicatedBook
	}

	exec := fmt.Sprintf(`INSERT INTO books(isbn, title, author, description, quantity, created_at, updated_at)
	VALUES
	($1, $2, $3, $4, $5, NOW(), NOW())`)

	_, err := b.db.ExecContext(ctx, exec, book.ISBN, book.Title, book.Author, book.Description, book.Quantity)

	return err
}

func (b *bookRepo) Select(ctx context.Context) ([]dto.ShowBooksDTO, error) {
	var books []dto.ShowBooksDTO
	query := fmt.Sprintf(`SELECT id, isbn, title, author, description, quantity FROM books ORDER BY id`)

	err := b.db.SelectContext(ctx, &books, query)

	return books, err
}

func (b *bookRepo) Get(ctx context.Context, id uint) (dto.ShowBooksDTO, error) {
	var book dto.ShowBooksDTO
	query := fmt.Sprintf(`SELECT id, isbn, title, author, description, quantity FROM books WHERE id = $1`)

	err := b.db.GetContext(ctx, &book, query, id)

	return book, err
}

func NewBookRepo(db *sqlx.DB) BookRepo {
	return &bookRepo{
		db: db,
	}
}
