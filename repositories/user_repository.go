package repositories

import (
	"errors"
	"votes/config"
	"votes/models"

	"github.com/google/uuid"
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

	err := u.db.Select("id, first_name, last_name, username, password, email, role_id").Preload("Role", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, name")
	}).Where("email = ?", email).First(&user).Error
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

	err := u.db.Select("id, first_name, last_name, username, email, role_id").Preload("Role", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, name")
	}).Where("username = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}

		return nil, err
	}

	return &user, nil
}

func (u *UserRepository) GetById(userId uuid.UUID) (*models.User, error) {
	var user models.User

	err := u.db.Select("id, first_name, last_name, username, email, role_id").Preload("Role", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, name")
	}).First(&user, userId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}

		return nil, err
	}

	return &user, nil
}

func (u *UserRepository) Update(userId uuid.UUID, user *models.User) (*models.User, error) {
	result := u.db.Where("id = ?", userId).Updates(user)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return user, nil
}
