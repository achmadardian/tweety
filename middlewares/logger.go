package middlewares

import (
	"votes/utils/logger"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const RequestId = "request_id"

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestId := c.GetHeader("X-Request-ID")
		if requestId == "" {
			requestId = uuid.New().String()
		}

		c.Writer.Header().Set("X-Request-ID", requestId)

		loggerWithReqID := logger.Log.With().Str(RequestId, requestId).Logger()
		ctx := loggerWithReqID.WithContext(c.Request.Context())
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
