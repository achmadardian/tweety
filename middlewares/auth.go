package middlewares

import (
	"strings"
	"votes/responses"
	"votes/services"
	"votes/utils/errs"
	z "votes/utils/logger"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Auth(a *services.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		prefix := "Bearer "

		if !strings.HasPrefix(authorization, prefix) {
			responses.Unauthorized(c)
			unauthorized(c, "missing or invalid Authorization header")
			return
		}

		accToken := strings.TrimPrefix(authorization, prefix)
		if accToken == " " {
			responses.Unauthorized(c)
			unauthorized(c, "empty access token")
			return
		}

		claim, err := a.ValidateToken(accToken)
		if err != nil {
			responses.Unauthorized(c)
			unauthorized(c, errs.ErrInvalidToken.Error())
			return
		}

		if claim.TokenType != services.TokenTypeAccess {
			responses.Unauthorized(c)
			unauthorized(c, "token is not access token type")
			return
		}

		userId, err := uuid.Parse(claim.Subject)
		if err != nil {
			responses.InternalServerError(c)
			z.Log.Error().
				Str("event", "middleware.auth").
				Err(err).
				Msg("failed to parse user_id")
			c.Abort()
			return
		}

		c.Set("user_id", userId)
		c.Next()
	}
}

func unauthorized(c *gin.Context, reason string) {
	z.Log.Warn().
		Str("event", "middleware.auth").
		Str("reason", reason).
		Msg("failed to authorize")
	c.Abort()
}
