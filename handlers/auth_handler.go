package handlers

import (
	"errors"
	"log"
	"votes/requests"
	"votes/responses"
	"votes/services"
	"votes/utils/errs"
	"votes/utils/validate"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authSvc *services.AuthService
}

func NewAuthHandler(authSvc *services.AuthService) *AuthHandler {
	return &AuthHandler{
		authSvc: authSvc,
	}
}

func (a *AuthHandler) Register(c *gin.Context) {
	var req requests.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		errVal := validate.ExtractValidationErrors(req, err)
		responses.UnprocessableEntity(c, errVal)
		return
	}

	register, err := a.authSvc.Register(&req)
	if err != nil {
		if errors.Is(err, errs.ErrEmailAlreadyExist) {
			responses.UnprocessableEntity(c, gin.H{"email": errs.ErrEmailAlreadyExist.Error()})
			return
		}

		if errors.Is(err, errs.ErrUsernameAlreadyExist) {
			responses.UnprocessableEntity(c, gin.H{"username": errs.ErrUsernameAlreadyExist.Error()})
			return
		}

		log.Printf("failed to register new user: %s", err)
		responses.InternalServerError(c)
		return
	}

	res := responses.RegisterResponse{
		ID:        register.ID,
		FirstName: register.FirstName,
		LastName:  *register.LastName,
		Username:  register.Username,
		Email:     register.Email,
	}

	responses.Created(c, res)
}
