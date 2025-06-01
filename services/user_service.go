package services

import (
	"errors"
	"fmt"
	"votes/models"
	"votes/repositories"
	"votes/requests"
	"votes/utils/errs"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (u *UserService) Create(req *requests.RegisterRequest) (*models.User, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("hash pasword: %w", err)
	}

	user := &models.User{
		ID:        uuid.New(),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Username:  req.Username,
		Email:     req.Email,
		Password:  string(hashedPass),
	}

	create, err := u.repo.Create(user)
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}

	return create, nil
}

func (u *UserService) GetByEmail(email string) (*models.User, error) {
	user, err := u.repo.GetByEmail(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.ErrDataNotFound
		}

		return nil, fmt.Errorf("get user by email: %w", err)
	}

	return user, nil
}

func (u *UserService) GetByUsername(username string) (*models.User, error) {
	user, err := u.repo.GetByUsername(username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.ErrDataNotFound
		}

		return nil, fmt.Errorf("get user by username: %w", err)
	}

	return user, nil
}

func (u *UserService) GetById(userId uuid.UUID) (*models.User, error) {
	user, err := u.repo.GetById(userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.ErrDataNotFound
		}

		return nil, fmt.Errorf("get user by id: %w", err)
	}

	return user, nil
}

func (u *UserService) Update(userId uuid.UUID, req *requests.UpdateMeRequest) (*models.User, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("hash password: %w", err)
	}

	user := &models.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Username:  req.Username,
		Email:     req.Email,
		Password:  string(hashedPass),
	}

	_, err = u.repo.Update(userId, user)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.ErrDataNotFound
		}

		return nil, fmt.Errorf("update user: %w", err)
	}

	newData, err := u.GetById(userId)
	if err != nil {
		return nil, err
	}

	return newData, nil
}
