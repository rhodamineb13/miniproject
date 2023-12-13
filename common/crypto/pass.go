package crypto

import (
	"miniproject/common/helper"

	"golang.org/x/crypto/bcrypt"
)

func GeneratePassword(password string) (string, error) {
	passByte := []byte("password")

	passEncrypt, err := bcrypt.GenerateFromPassword(passByte, 10)
	if err != nil {
		return "", helper.ErrGeneratePassword
	}

	return string(passEncrypt), nil
}

func ComparePassword(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(hash))

	return err
}
