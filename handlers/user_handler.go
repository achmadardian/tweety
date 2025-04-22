package handlers

import (
	"log"
	"strings"
	"votes/repositories"
	"votes/requests"
	"votes/response"
	"votes/utils"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userRepo *repositories.UserRepository
}

func NewUserHandler(userRepo *repositories.UserRepository) *UserHandler {
	return &UserHandler{
		userRepo: userRepo,
	}
}

func (u *UserHandler) GetUserAll(c *gin.Context) {
	page, err := utils.GetQueryParamPagination(c)
	if err != nil {
		response.BadRequest(c, "invalid query params")
		return
	}

	keyword := c.Query("name")
	keyword = strings.TrimSpace(keyword)
	users, err := u.userRepo.GetAll(page, keyword)
	if err != nil {
		log.Printf("[UserHandler.GetUserAll] failed to get user data: %v", err)
		response.InternalServerError(c)
		return
	}

	userMaps := make([]response.UserResponse, 0, len(users))
	for _, user := range users {
		userMaps = append(userMaps, response.UserResponse{
			Id:   user.Id,
			Name: user.Name,
		})
	}

	response.OkPaginate(c, userMaps, page, "user data")
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req requests.UserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		errFields := utils.ExtractValidationErrors(err)
		response.UnprocessableEntity(c, errFields)
		return
	}

	create, err := h.userRepo.Create(&req)
	if err != nil {
		log.Printf("[UserHandler.CreateUser] failed to create user: %v", err)
		response.InternalServerError(c)
		return
	}

	res := response.UserResponse{
		Id:   create.Id,
		Name: req.Name,
	}

	response.Created(c, res, "create user data")
}
