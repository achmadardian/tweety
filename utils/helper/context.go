package helper

import (
	"votes/responses"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

func GetUserIdFromContext(c *gin.Context, z *zerolog.Logger, logEvent string) (uuid.UUID, bool) {
	userId, ok := c.Get("user_id")
	if !ok {
		z.Error().
			Str("event", logEvent).
			Msg("missing user_id from context")

		responses.InternalServerError(c)
		return uuid.Nil, false
	}

	userIdStr := userId.(string)
	userUUID, err := uuid.Parse(userIdStr)
	if err != nil {
		z.Error().
			Str("event", logEvent).
			Msg("failed to parse user_id ")

		responses.InternalServerError(c)
		return uuid.Nil, false
	}

	return userUUID, true
}
