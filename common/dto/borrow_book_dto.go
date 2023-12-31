package dto

import "time"

type BorrowBookRequestDTO struct {
	BookID uint `json:"book_id"`
	UserID uint `json:"user_id"`
}

type BorrowListDTO struct {
	ID         uint      `db:"id" json:"id"`
	BookID     uint      `db:"book_id" json:"book_id"`
	Title      string    `db:"title" json:"title"`
	UserID     uint      `db:"user_id" json:"user_id"`
	UserName   string    `db:"name" json:"name"`
	BorrowedAt time.Time `db:"borrowed_at" json:"borrowed_at"`
	ReturnedAt time.Time `db:"returned_at" json:"returned_at"`
}

type ReturnBookRequest struct {
	BookID uint `json:"book_id"`
	UserID uint
}
