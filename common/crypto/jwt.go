package crypto

import (
	"miniproject/common/config"
	"miniproject/common/helper"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Role string

var (
	ADMIN Role = "admin"
	USER  Role = "user"
)

type JWTClaims struct {
	ID   uint
	Role []Role
	*jwt.RegisteredClaims
}

func NewClaims(id uint, roles []Role) *JWTClaims {

	claims := &jwt.RegisteredClaims{
		Issuer: config.Config.Issuer,
		ExpiresAt: &jwt.NumericDate{
			Time: time.Now().Add(time.Duration(config.Config.Duration)),
		},
	}

	return &JWTClaims{
		ID:               id,
		Role:             roles,
		RegisteredClaims: claims,
	}
}

func NewToken(id uint, roles []Role) (string, error) {
	claims := NewClaims(id, roles)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessToken, err := token.SignedString(config.Config.LibSecretKey)
	if err != nil {
		return "", helper.ErrCreateToken
	}

	return accessToken, nil
}
