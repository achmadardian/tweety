package errs

import "errors"

var (
	ErrDataNotFound         = errors.New("data not found")
	ErrEmailAlreadyExist    = errors.New("email already exist")
	ErrUsernameAlreadyExist = errors.New("username already exist")
	ErrInvalidLogin         = errors.New("invalid login")
)
