package services

import (
	"votes/models"
	"votes/repositories"
	"votes/requests"
	"votes/response"
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
		return nil, err
	}

	return users, nil
}

func (u *UserService) Create(req *requests.UserRequest) (*models.User, error) {
	user, err := u.userRepo.Create(req)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserService) GetById(id int) (*models.User, error) {
	user, err := u.userRepo.GetById(id)
	if err != nil {
		return nil, err
	}

	return user, err
}

func (u *UserService) Update(req *requests.UserRequestUpdate, id int) error {
	_, err := u.userRepo.GetById(id)
	if err != nil {
		return err
	}

	if err = u.userRepo.Update(req, id); err != nil {
		return err
	}

	return nil
}

func (u *UserService) Delete(id int) error {
	_, err := u.userRepo.GetById(id)
	if err != nil {
		return err
	}

	if err := u.userRepo.Delete(id); err != nil {
		return err
	}

	return nil
}
