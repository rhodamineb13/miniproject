package service

import (
	"context"
	"errors"
	"miniproject/common/config"
	"miniproject/common/crypto"
	"miniproject/common/dto"
	"miniproject/common/helper"
	"miniproject/repository"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

type userService struct {
	redis    *redis.Client
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

	passHash, err := crypto.GeneratePassword(reg.Password)
	if err != nil {
		return err
	}

	regDB := &dto.RegisterDBDTO{
		Name:     reg.Name,
		DOB:      DOB,
		Email:    reg.Email,
		Password: passHash,
	}
	err = u.userRepo.Insert(ctx, regDB)

	if err != nil {
		if errors.Is(err, helper.ErrUserExists) {
			return err
		}
		return helper.ErrRegisterFailed
	}
	return nil
}

func (u *userService) Login(ctx context.Context, login *dto.UserLoginDTO) (*dto.AccessTokenDTO, error) {
	getInitCounter, err := u.redis.Get(ctx, login.Email).Int()
	if err == nil {
		if getInitCounter == 3 {
			return nil, helper.ErrTemporarilyBanned
		}
	}

	//Admin Access Token

	if login.Email == os.Getenv("ADMIN_EMAIL") {
		hashPwd, err := crypto.GeneratePassword(login.Password)
		if err != nil {
			return nil, helper.ErrGeneratePassword
		}

		if err := crypto.ComparePassword(hashPwd, config.Config.AdminPassword); err != nil {
			return nil, helper.ErrLogin
		}

		roles := []crypto.Role{crypto.ADMIN, crypto.USER}
		accessToken, err := crypto.GenerateNewToken(0, roles)
		if err != nil {
			return nil, helper.ErrGenerateToken
		}
		return &dto.AccessTokenDTO{
			AccessToken: accessToken,
		}, nil
	}

	//User Access Token

	var BanCounter int

	user, errEmail := u.userRepo.VerifyUser(ctx, login)

	errPwd := crypto.ComparePassword(user.Password, login.Password)

	if errEmail != nil || errPwd != nil {
		getCounter, err := u.redis.Get(ctx, login.Email).Int()
		if err != nil {
			BanCounter++
			u.redis.Set(ctx, login.Email, BanCounter, time.Minute*5)
		}
		if getCounter == 3 {
			u.redis.Set(ctx, login.Email, getCounter, time.Minute*15)
			return nil, helper.ErrTemporarilyBanned
		}
		getCounter++
		u.redis.Set(ctx, login.Email, getCounter, time.Minute*5)
		return nil, helper.ErrLogin
	}

	id, err := u.userRepo.FindIDByEmail(ctx, login.Email)
	if err != nil {
		return nil, helper.ErrUserNotFound
	}

	accessToken, err := crypto.GenerateNewToken(id, []crypto.Role{crypto.USER})
	if err != nil {
		return nil, err
	}

	return &dto.AccessTokenDTO{
		ID:          user.ID,
		AccessToken: accessToken,
	}, nil
}

func NewUserService(redis *redis.Client, userRepo repository.UserRepo) UserService {
	return &userService{
		redis:    redis,
		userRepo: userRepo,
	}
}
