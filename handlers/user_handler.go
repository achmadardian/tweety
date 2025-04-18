package handlers

import (
	"log"
	"votes/repositories"
	"votes/response"

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
	users, err := u.userRepo.GetAll()
	if err != nil {
		log.Printf("[UserHandler.GetUserAll] failed to get user data: %v", err)
		response.InternalServerError(c)
		return
	}

	response.Ok(c, users, "user data")
}
