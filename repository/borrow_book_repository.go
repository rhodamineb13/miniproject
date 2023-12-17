package repository

import (
	"context"
	"database/sql"
	"fmt"
	"miniproject/common/dto"
	"miniproject/common/helper"

	"github.com/jmoiron/sqlx"
)

type borrowRepo struct {
	db *sqlx.DB
}

type BorrowRepo interface {
	Borrow(context.Context, *dto.BorrowBookRequestDTO) error
}

func (b *borrowRepo) Borrow(ctx context.Context, req *dto.BorrowBookRequestDTO) error {
	var bookID uint

	queryFind := fmt.Sprintf(`SELECT 1 FROM borrow_records WHERE book_id = $1 AND user_id = $2`)

	if err := b.db.GetContext(ctx, &bookID, queryFind, req.BookID, req.UserID); err == nil {
		return helper.ErrBookAlreadyBorrowed
	}

	tx, err := b.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
	})

	if err != nil {
		return err
	}

	execQueryBorrow := fmt.Sprintf(`INSERT INTO borrow_records(user_id, book_id, borrowed_at)
	VALUES
	($1, $2, NOW())`)

	_, err = tx.ExecContext(ctx, execQueryBorrow, req.UserID, req.BookID)

	if err != nil {
		return err
	}

	execQueryQty := fmt.Sprintf(`UPDATE books
	SET quantity = quantity - 1
	WHERE id = $1`)

	_, err = tx.ExecContext(ctx, execQueryQty, req.BookID)

	if err != nil {
		return err
	}

	return nil
}

func NewBorrowRepo(db *sqlx.DB) BorrowRepo {
	return &borrowRepo{
		db: db,
	}
}