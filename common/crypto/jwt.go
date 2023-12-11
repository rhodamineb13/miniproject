package crypto

import (
	"log"
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
	Email string
	Role  []Role
	*jwt.RegisteredClaims
}

func NewClaims(email string, roles []Role) jwt.Claims {

	claims := &jwt.RegisteredClaims{
		Issuer: config.Config.Issuer,
		ExpiresAt: &jwt.NumericDate{
			Time: time.Now().Add(time.Duration(config.Config.Duration)),
		},
	}

	return &JWTClaims{
		Email:            email,
		Role:             roles,
		RegisteredClaims: claims,
	}
}

func GenerateNewToken(email string, roles []Role) (string, error) {
	claims := NewClaims(email, roles)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessToken, err := token.SignedString([]byte(config.Config.LibSecretKey))
	if err != nil {
		log.Println(err)
		return "", helper.ErrCreateToken
	}

	return accessToken, nil
}

func ValidateToken(accessToken string) (*JWTClaims, error) {
	var claims *JWTClaims
	token, err := jwt.ParseWithClaims(accessToken, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, helper.ErrParseToken
		}
		return config.Config.LibSecretKey, nil
	})

	if err != nil || !token.Valid {
		return nil, helper.ErrInvalidToken
	}

	return claims, nil
}
