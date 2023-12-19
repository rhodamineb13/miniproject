package repository

import (
	"context"
	"fmt"
	"miniproject/common/dto"
	"miniproject/common/helper"

	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}

type UserRepo interface {
	Insert(context.Context, *dto.RegisterDBDTO) error
	VerifyUser(context.Context, string) (dto.GetUserDTO, error)
	FindIDByEmail(context.Context, string) (uint, error)
	UpdatePassword(context.Context, *dto.ChangePasswordDTO) error
	FindEmailByID(context.Context, uint) (string, error)
}

func (u *userRepo) Insert(ctx context.Context, reg *dto.RegisterDBDTO) error {
	var userID uint

	query := fmt.Sprintf(`SELECT id FROM users WHERE email = $1`)

	if err := u.db.GetContext(ctx, &userID, query, reg.Email); err == nil {
		return helper.ErrUserExists
	}

	exec := fmt.Sprintf(`INSERT INTO users (name, date_of_birth, email, password, created_at, updated_at)
	VALUES
	($1, $2, $3, $4, NOW(), NOW())`)

	_, err := u.db.ExecContext(ctx, exec, reg.Name, reg.DOB, reg.Email, reg.Password)
	return err
}

func (u *userRepo) VerifyUser(ctx context.Context, email string) (dto.GetUserDTO, error) {
	var user dto.GetUserDTO
	query := fmt.Sprintf(`SELECT * FROM users WHERE email = $1`)

	err := u.db.GetContext(ctx, &user, query, email)

	return user, err
}

func (u *userRepo) FindIDByEmail(ctx context.Context, email string) (uint, error) {
	var id uint
	query := fmt.Sprintf(`SELECT id FROM users WHERE email = $1`)

	err := u.db.GetContext(ctx, &id, query, email)

	return id, err
}

func (u *userRepo) FindEmailByID(ctx context.Context, id uint) (string, error) {
	var email string
	query := fmt.Sprintf(`SELECT email FROM users WHERE id = $1`)

	err := u.db.GetContext(ctx, &email, query, id)

	return email, err
}

func (u *userRepo) UpdatePassword(ctx context.Context, change *dto.ChangePasswordDTO) error {
	queryUpdate := fmt.Sprintf(`UPDATE users SET password = $1 WHERE id = $3`)
	_, err := u.db.ExecContext(ctx, queryUpdate, change.NewPassword, change.ID)
	return err
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}
