package repository

import (
	"context"
	"fmt"
	"miniproject/common/dto"

	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}

type UserRepo interface {
	Insert(context.Context, *dto.RegisterDBDTO) error
}

func (u *userRepo) Insert(ctx context.Context, reg *dto.RegisterDBDTO) error {
	exec := fmt.Sprintf(`INSERT INTO users (name, date_of_birth, email, password)
	VALUES
	($1, $2, $3, $4)`)

	_, err := u.db.ExecContext(ctx, exec, reg.Name, reg.DOB, reg.Email, reg.Password)
	return err
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}
