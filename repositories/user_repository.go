package repositories

import (
	"errors"
	"votes/config"
	"votes/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *config.Database
}

func NewUserRepository(db *config.Database) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) Create(user *models.User) (*models.User, error) {
	err := u.db.Create(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User

	err := u.db.Select("id, first_name, last_name, username, email").Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}

		return nil, err
	}

	return &user, nil
}

func (u *UserRepository) GetByUsername(username string) (*models.User, error) {
	var user models.User

	err := u.db.Select("id, first_name, last_name, username, email").Where("username = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}

		return nil, err
	}

	return &user, nil
}
