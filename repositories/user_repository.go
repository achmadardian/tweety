package repositories

import (
	"votes/config"
	"votes/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	ReadConnection  *gorm.DB
	WriteConnection *gorm.DB
}

func NewUserRepository(DB *config.Database) *UserRepository {
	return &UserRepository{
		ReadConnection:  DB.ReadConnection,
		WriteConnection: DB.WriteConnection,
	}
}

func (u *UserRepository) GetAll() ([]models.User, error) {
	var users []models.User

	query := u.ReadConnection.Model(&models.User{}).Find(&users)
	if query.Error != nil {
		return nil, query.Error
	}

	return users, nil
}
