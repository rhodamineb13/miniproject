package dto

import "time"

type RegisterUserDTO struct {
	Name     string    `json:"name" binding:"required"`
	DOB      time.Time `json:"date_of_birth" format:"2006-01-02"`
	Email    time.Time `json:"email" binding:"required,email"`
	Password string    `json:"password" binding:"required,min=8"`
}
