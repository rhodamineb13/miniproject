package service

import (
	"context"
	"log"
	"miniproject/common/dto"
	"miniproject/common/helper"
	"miniproject/repository"
	"time"
)

type userService struct {
	userRepo repository.UserRepo
}

type UserService interface {
	Register(context.Context, *dto.RegisterUserDTO) error
}

func (u *userService) Register(ctx context.Context, reg *dto.RegisterUserDTO) error {
	DOB, err := time.Parse(helper.TimeParseLayout, reg.DOB)
	if err != nil {
		return helper.ErrParseTimeFormat
	}

	regDB := &dto.RegisterDBDTO{
		Name:     reg.Name,
		DOB:      DOB,
		Email:    reg.Email,
		Password: reg.Password,
	}
	err = u.userRepo.Insert(ctx, regDB)

	if err != nil {
		log.Println(err)
		return helper.ErrRegisterFailed
	}
	return nil
}

func NewUserService(userRepo repository.UserRepo) UserService {
	return &userService{
		userRepo: userRepo,
	}
}
