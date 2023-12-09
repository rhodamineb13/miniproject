package dto

import "time"

type RegisterUserDTO struct {
	Name     string `json:"name" binding:"required"`
	DOB      string `json:"date_of_birth" format:"2006-01-02"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type RegisterDBDTO struct {
	Name     string    `db:"name"`
	DOB      time.Time `db:"date_of_birth"`
	Email    string    `db:"email"`
	Password string    `db:"password"`
}
