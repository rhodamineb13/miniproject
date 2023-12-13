package helper

import "net/http"

var (
	ErrRegisterFailed   = NewCustomError(http.StatusBadRequest, "error in registering user")
	ErrParseTimeFormat  = NewCustomError(http.StatusBadRequest, "error parsing time format")
	ErrLogin            = NewCustomError(http.StatusBadRequest, "email or password is incorrect")
	ErrAddBook          = NewCustomError(http.StatusBadRequest, "error adding new book")
	ErrBooksEmpty       = NewCustomError(http.StatusOK, "there are no books yet")
	ErrBookNotFound     = NewCustomError(http.StatusNotFound, "book not found")
	ErrCreateToken      = NewCustomError(http.StatusInternalServerError, "create token failed")
	ErrParseToken       = NewCustomError(http.StatusInternalServerError, "error parse token")
	ErrInvalidToken     = NewCustomError(http.StatusForbidden, "invalid token, access denied")
	ErrDuplicatedBook   = NewCustomError(http.StatusForbidden, "error book already added")
	ErrUserExists       = NewCustomError(http.StatusBadRequest, "user already exists")
	ErrGeneratePassword = NewCustomError(http.StatusInternalServerError, "error generating password")
)

const (
	TimeParseLayout = "2006-01-02"
)
