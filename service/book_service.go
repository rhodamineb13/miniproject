package service

import (
	"context"
	"errors"
	"log"
	"miniproject/common/dto"
	"miniproject/common/helper"
	"miniproject/repository"
)

type bookService struct {
	bookRepo repository.BookRepo
}

type BookService interface {
	AddBook(context.Context, *dto.AddBookDTO) error
	GetAllBooks(context.Context) ([]dto.ShowBooksDTO, error)
	FindBookByID(context.Context, uint) (dto.ShowBooksDTO, error)
}

func (b *bookService) AddBook(ctx context.Context, book *dto.AddBookDTO) error {
	if err := b.bookRepo.Insert(ctx, book); err != nil {
		if errors.Is(err, helper.ErrDuplicatedBook) {
			return err
		}
		return helper.ErrAddBook
	}
	return nil
}

func (b *bookService) GetAllBooks(ctx context.Context) ([]dto.ShowBooksDTO, error) {
	books, err := b.bookRepo.Select(ctx)
	if err != nil {
		return nil, helper.ErrBooksEmpty
	}

	return books, nil
}

func (b *bookService) FindBookByID(ctx context.Context, bookID uint) (dto.ShowBooksDTO, error) {
	book, err := b.bookRepo.Get(ctx, bookID)
	if err != nil {
		log.Println(err)
		return book, helper.ErrBookNotFound
	}

	return book, nil
}

func NewBookService(bookRepo repository.BookRepo) BookService {
	return &bookService{
		bookRepo: bookRepo,
	}
}
