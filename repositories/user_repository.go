package repositories

import (
	"votes/config"
	"votes/models"
	"votes/requests"
	"votes/response"
	"votes/utils"

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

	query := u.ReadConnection.Model(&models.User{}).
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

func (u *UserRepository) Create(req *requests.UserRequest) (*models.User, error) {
	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	query := u.WriteConnection.Create(&user)
	if query.Error != nil {
		return nil, query.Error
	}

	return &user, nil
}
