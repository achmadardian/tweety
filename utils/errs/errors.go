package errs

import "errors"

var (
	ErrDataNotFound      = errors.New("data not found")
	ErrEmailAlreadyExist = errors.New("email already exist")
)
