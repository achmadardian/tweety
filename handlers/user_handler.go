package handlers

import (
	"errors"
	"votes/responses"
	"votes/services"
	"votes/utils/errs"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

type UserHandler struct {
	userSvc *services.UserService
}

func NewUserHandler(userSvc *services.UserService) *UserHandler {
	return &UserHandler{
		userSvc: userSvc,
	}
}

func (u *UserHandler) Me(c *gin.Context) {
	z := zerolog.Ctx(c.Request.Context())
	userId := uuid.New() // it should be call helper to get userid from context

	user, err := u.userSvc.GetById(userId)
	if err != nil {
		if errors.Is(err, errs.ErrDataNotFound) {
			z.Warn().
				Str("user_id", userId.String()).
				Msg("failed to fetch user by id: data not found")

			responses.NotFound(c)
			return
		}

		z.Error().
			Str("user_id", userId.String()).
			Err(err).
			Msg("failed to fetch user by id")

		responses.InternalServerError(c)
		return
	}

	res := responses.Me{
		Id:        user.ID,
		FirstName: user.FirstName,
		LastName:  *user.LastName,
		Username:  user.Username,
		Email:     user.Email,
	}

	responses.Ok(c, res)
}
