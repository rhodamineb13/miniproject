package service

import (
	"context"
	"log"
	"miniproject/common/crypto"
	"miniproject/common/dto"
	"miniproject/common/helper"
	"miniproject/repository"
	"os"
	"time"
)

type userService struct {
	userRepo repository.UserRepo
}

type UserService interface {
	Register(context.Context, *dto.RegisterUserDTO) error
	Login(context.Context, *dto.UserLoginDTO) (*dto.AccessTokenDTO, error)
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
		return helper.ErrRegisterFailed
	}
	return nil
}

func (u *userService) Login(ctx context.Context, login *dto.UserLoginDTO) (*dto.AccessTokenDTO, error) {
	//Admin Access Token

	if login.Email == os.Getenv("ADMIN_EMAIL") && login.Password == os.Getenv("ADMIN_PASSWORD") {
		roles := []crypto.Role{crypto.ADMIN, crypto.USER}
		accessToken, err := crypto.GenerateNewToken(login.Email, roles)
		if err != nil {
			return nil, err
		}
		return &dto.AccessTokenDTO{
			AccessToken: accessToken,
		}, nil
	}

	id, err := u.userRepo.VerifyUser(ctx, login)
	if err != nil {
		log.Println(err)
		return nil, helper.ErrLogin
	}

	//User Access Token

	accessToken, err := crypto.GenerateNewToken(login.Email, []crypto.Role{crypto.USER})
	if err != nil {
		return nil, err
	}

	return &dto.AccessTokenDTO{
		ID:          id,
		AccessToken: accessToken,
	}, nil
}

func NewUserService(userRepo repository.UserRepo) UserService {
	return &userService{
		userRepo: userRepo,
	}
}
