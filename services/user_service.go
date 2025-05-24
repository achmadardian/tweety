package services

import (
	"errors"
	"fmt"
	"votes/models"
	"votes/repositories"
	"votes/requests"
	"votes/response"
	"votes/utils/errs"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (u *UserService) GetAll(page *response.PaginatedResponse, keyword string) ([]models.User, error) {
	users, err := u.userRepo.GetAll(page, keyword)
	if err != nil {
		return nil, fmt.Errorf("get all user: %w", err)
	}

	return users, nil
}

func (u *UserService) Create(req *requests.UserRequest) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("hashed password: %w", err)
	}

	user := &models.User{
		Id:       uuid.New(),
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	_, err = u.GetByEmail(req.Email)
	if err != nil {
		if errors.Is(err, errs.ErrDataNotFound) {
			return nil, errs.ErrEmailAlreadyExist
		}
	}

	save, err := u.userRepo.Create(user)
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}

	return save, nil
}

func (u *UserService) GetById(id uuid.UUID) (*models.User, error) {
	user, err := u.userRepo.GetById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.ErrDataNotFound
		}

		return nil, fmt.Errorf("get user by id: %w", err)
	}

	return user, err
}

func (u *UserService) GetByEmail(email string) (*models.User, error) {
	user, err := u.userRepo.GetByEmail(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.ErrDataNotFound
		}

		return nil, fmt.Errorf("get user by email: %w", err)
	}

	return user, err
}

func (u *UserService) Update(req *requests.UserRequestUpdate, id uuid.UUID) error {
	_, err := u.userRepo.GetById(id)
	if err != nil {
		return err
	}

	if err = u.userRepo.Update(req, id); err != nil {
		return fmt.Errorf("update user: %w", err)
	}

	return nil
}

func (u *UserService) Delete(id uuid.UUID) error {
	_, err := u.userRepo.GetById(id)
	if err != nil {
		return err
	}

	if err := u.userRepo.Delete(id); err != nil {
		return fmt.Errorf("delete user: %", err)
	}

	return nil
}
