package service

import (
	"context"
	"errors"
	"miniproject/common/dto"
	"miniproject/common/helper"
	"miniproject/repository"
)

type borrowService struct {
	borrowRepo repository.BorrowRepo
}

type BorrowService interface {
	Borrow(context.Context, *dto.BorrowBookRequestDTO) error
	ReturnBook(context.Context, *dto.ReturnBookRequest) error
}

func (b *borrowService) Borrow(ctx context.Context, req *dto.BorrowBookRequestDTO) error {
	err := b.borrowRepo.Borrow(ctx, req)
	if err != nil {
		switch {
		case errors.Is(err, helper.ErrBookAlreadyBorrowed) || errors.Is(err, helper.ErrBookEmpty):
			return err
		default:
			return helper.ErrBorrowBook
		}
	}
	return nil
}

func (b *borrowService) ReturnBook(ctx context.Context, ret *dto.ReturnBookRequest) error {
	if err := b.borrowRepo.ReturnBook(ctx, ret); err != nil {
		return helper.ErrReturnBook
	}

	return nil
}

func NewBorrowService(borrowRepo repository.BorrowRepo) BorrowService {
	return &borrowService{
		borrowRepo: borrowRepo,
	}
}
