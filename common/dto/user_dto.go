package dto

import "time"

type RegisterUserDTO struct {
	Name     string `json:"name" binding:"required"`
	DOB      string `json:"date_of_birth" format:"2006-01-02"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type UserLoginDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type GetUserDTO struct {
	ID        uint       `db:"id"`
	Name      string     `db:"name"`
	DOB       time.Time  `db:"date_of_birth"`
	Email     string     `db:"email"`
	Password  string     `db:"password"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}

type RegisterDBDTO struct {
	Name     string    `db:"name"`
	DOB      time.Time `db:"date_of_birth"`
	Email    string    `db:"email"`
	Password string    `db:"password"`
}

type AccessTokenDTO struct {
	ID          uint   `json:"-"`
	AccessToken string `json:"access_token"`
}

type ChangePasswordDTO struct {
	ID          uint
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}
