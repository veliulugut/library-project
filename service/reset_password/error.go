package reset_password

import "errors"

var (
	ErrExpiredValidationCode = errors.New("expired validation code")
	ErrPasswordNotMatch      = errors.New("password not match")
)
