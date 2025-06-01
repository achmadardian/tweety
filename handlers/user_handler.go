package handlers

import (
	"errors"
	"votes/responses"
	"votes/services"
	"votes/utils/errs"
	"votes/utils/helper"

	"github.com/gin-gonic/gin"
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
	userId, ok := helper.GetUserIdFromContext(c, z, "user.me")
	if !ok {
		return
	}

	user, err := u.userSvc.GetById(userId)
	if err != nil {
		if errors.Is(err, errs.ErrDataNotFound) {
			z.Warn().
				Str("user_id", userId.String()).
				Str("reason", "data not found").
				Msg("failed to fetch user by id")

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
		LastName:  user.LastName,
		Username:  user.Username,
		Email:     user.Email,
	}

	responses.Ok(c, res)
}
