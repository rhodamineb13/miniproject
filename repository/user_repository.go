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
	VerifyUser(context.Context, *dto.UserLoginDTO) error
}

func (u *userRepo) Insert(ctx context.Context, reg *dto.RegisterDBDTO) error {
	exec := fmt.Sprintf(`INSERT INTO users (name, date_of_birth, email, password, created_at, updated_at)
	VALUES
	($1, $2, $3, $4, NOW(), NOW())`)

	_, err := u.db.ExecContext(ctx, exec, reg.Name, reg.DOB, reg.Email, reg.Password)
	return err
}

func (u *userRepo) VerifyUser(ctx context.Context, login *dto.UserLoginDTO) error {
	var user dto.GetUserDTO
	query := fmt.Sprintf(`SELECT * FROM users WHERE email = $1 AND password = $2`)

	err := u.db.GetContext(ctx, &user, query, login.Email, login.Password)

	return err
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}
