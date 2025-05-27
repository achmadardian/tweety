package services

import (
	"errors"
	"fmt"
	"votes/models"
	"votes/requests"
	"votes/utils/errs"
)

type AuthService struct {
	userSvc *UserService
}

func NewAuthService(userSvc *UserService) *AuthService {
	return &AuthService{
		userSvc: userSvc,
	}
}

func (a *AuthService) Register(req *requests.RegisterRequest) (*models.User, error) {
	if _, err := a.userSvc.GetByEmail(req.Email); err == nil {
		return nil, errs.ErrEmailAlreadyExist
	} else if !errors.Is(err, errs.ErrDataNotFound) {
		return nil, fmt.Errorf("check existing email: %w", err)
	}

	if _, err := a.userSvc.GetByUsername(req.Username); err == nil {
		return nil, errs.ErrUsernameAlreadyExist
	} else if !errors.Is(err, errs.ErrDataNotFound) {
		return nil, fmt.Errorf("check existing username: %w", err)
	}

	register, err := a.userSvc.Create(req)
	if err != nil {
		return nil, fmt.Errorf("register: %w", err)
	}

	return register, nil
}
