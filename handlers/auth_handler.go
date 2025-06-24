package handlers

import (
	"errors"
	"votes/requests"
	"votes/responses"
	"votes/services"
	"votes/utils/errs"
	"votes/utils/validate"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
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
	z := zerolog.Ctx(c.Request.Context())

	var req requests.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		errVal := validate.ExtractValidationErrors(req, err)
		responses.UnprocessableEntity(c, errVal)
		return
	}

	register, err := a.authSvc.Register(&req)
	if err != nil {
		if errors.Is(err, errs.ErrEmailAlreadyExist) {
			z.Warn().
				Str("event", "auth.register").
				Str("email", req.Email).
				Str("reason", errs.ErrEmailAlreadyExist.Error()).
				Msg("failed to register")

			responses.UnprocessableEntity(c, gin.H{"email": errs.ErrEmailAlreadyExist.Error()})
			return
		}

		if errors.Is(err, errs.ErrUsernameAlreadyExist) {
			z.Warn().
				Str("event", "auth.register").
				Str("username", req.Username).
				Str("reason", errs.ErrUsernameAlreadyExist.Error()).
				Msg("failed to register")

			responses.UnprocessableEntity(c, gin.H{"username": errs.ErrUsernameAlreadyExist.Error()})
			return
		}

		z.Error().
			Str("event", "auth.register").
			Str("email", req.Email).
			Str("username", req.Username).
			Err(err).
			Msg("failed to register")

		responses.InternalServerError(c)
		return
	}

	res := responses.RegisterResponse{
		ID:        register.ID,
		FirstName: register.FirstName,
		LastName:  register.LastName,
		Username:  register.Username,
		Email:     register.Email,
		Role: responses.RoleUserResponse{
			ID:   register.Role.ID,
			Name: register.Role.Name,
		},
	}

	responses.Created(c, res)
}

func (a *AuthHandler) Login(c *gin.Context) {
	z := zerolog.Ctx(c.Request.Context())

	var req requests.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		errVal := validate.ExtractValidationErrors(req, err)
		responses.UnprocessableEntity(c, errVal)
		return
	}

	acc, ref, err := a.authSvc.Login(&req)
	if err != nil {
		if errors.Is(err, errs.ErrInvalidLogin) {
			z.Warn().
				Str("event", "auth.login").
				Str("email", req.Email).
				Str("reason", errs.ErrInvalidLogin.Error()).
				Msg("failed to login")

			responses.Unauthorized(c, errs.ErrInvalidLogin.Error())
			return
		}

		z.Error().
			Str("event", "auth.login").
			Str("email", req.Email).
			Err(err).
			Msg("failed to login")

		responses.InternalServerError(c)
		return
	}

	res := responses.LoginResponse{
		AccessToken:  acc,
		RefreshToken: ref,
	}

	responses.Ok(c, res)
}

func (a *AuthHandler) RefreshToken(c *gin.Context) {
	z := zerolog.Ctx(c.Request.Context())
	var req requests.RefreshTokenRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		errVal := validate.ExtractValidationErrors(req, err)
		responses.UnprocessableEntity(c, errVal)
		return
	}

	newToken, err := a.authSvc.RefreshToken(&req)
	if err != nil {
		if errors.Is(err, errs.ErrInvalidToken) {
			z.Warn().
				Str("event", "auth.refresh_token").
				Str("reason", errs.ErrInvalidToken.Error()).
				Msg("failed to generate new access token")

			responses.Unauthorized(c)
			return
		}

		z.Error().
			Str("event", "auth.refresh_token").
			Err(err).
			Msg("failed to generate new access token")

		responses.InternalServerError(c)
		return
	}

	res := responses.RefreshTokenResponse{
		AccessToken: newToken,
	}

	responses.Ok(c, res)
}
