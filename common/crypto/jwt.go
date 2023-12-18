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
	ID   uint
	Role []Role
	*jwt.RegisteredClaims
}

func NewClaims(id uint, roles []Role) jwt.Claims {

	claims := &jwt.RegisteredClaims{
		Issuer: config.Config.Issuer,
		ExpiresAt: &jwt.NumericDate{
			Time: time.Now().Add(time.Hour * time.Duration(config.Config.Duration)),
		},
	}

	return &JWTClaims{
		ID:               id,
		Role:             roles,
		RegisteredClaims: claims,
	}
}

func GenerateNewToken(id uint, roles []Role) (string, error) {
	claims := NewClaims(id, roles)
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
	token, err := jwt.ParseWithClaims(accessToken, &JWTClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, helper.ErrParseToken
		}
		return []byte(config.Config.LibSecretKey), nil
	})

	if err != nil || !token.Valid {
		return nil, helper.ErrInvalidToken
	}

	if claims, ok := token.Claims.(*JWTClaims); ok {
		return claims, nil
	}

	return claims, err
}
