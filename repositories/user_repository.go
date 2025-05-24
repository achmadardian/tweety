package repositories

import (
	"errors"
	"votes/config"
	"votes/models"
	"votes/requests"
	"votes/response"
	"votes/utils"

	"github.com/google/uuid"
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

func (u *UserRepository) GetAll(page *response.PaginatedResponse, keyword string) ([]models.User, error) {
	var users []models.User

	query := u.ReadConnection.
		Select("id", "name", "email").
		Scopes(utils.Paginate(page.Page, page.PageSize)).
		Order("id DESC")

	if keyword != "" {
		query = query.Where("name LIKE ?", "%"+keyword+"%")
	}

	if err := query.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (u *UserRepository) Create(user *models.User) (*models.User, error) {
	query := u.WriteConnection.Create(&user)
	if query.Error != nil {
		return nil, query.Error
	}

	return user, nil
}

func (u *UserRepository) GetById(id uuid.UUID) (*models.User, error) {
	var user models.User

	err := u.WriteConnection.Select("id", "name", "email").First(&user, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}

		return nil, err
	}

	return &user, nil
}

func (u *UserRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User

	err := u.WriteConnection.Select("id", "name", "email").Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}

		return nil, err
	}

	return &user, nil
}

func (u *UserRepository) Update(req *requests.UserRequestUpdate, id uuid.UUID) error {
	user := models.User{}

	return u.WriteConnection.Model(&user).Where("id = ?", id).Updates(req).Error
}

func (u *UserRepository) Delete(id uuid.UUID) error {
	return u.WriteConnection.Delete(&models.User{}, id).Error
}
