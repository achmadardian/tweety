package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiResponse struct {
	Code    int
	Message string
	Data    interface{}
}

func Ok(c *gin.Context, data interface{}, message ...string) {
	msg := "success"
	if len(message) > 0 {
		msg = message[0]
	}

	c.JSON(http.StatusOK, ApiResponse{
		Code:    http.StatusOK,
		Message: msg,
		Data:    data,
	})
}

func Created(c *gin.Context, data interface{}, message ...string) {
	msg := "created"
	if len(message) > 0 {
		msg = message[0]
	}

	c.JSON(http.StatusCreated, ApiResponse{
		Code:    http.StatusCreated,
		Message: msg,
		Data:    data,
	})
}

func Deleted(c *gin.Context, message ...string) {
	msg := "deleted"
	if len(message) > 0 {
		msg = message[0]
	}

	c.JSON(http.StatusOK, ApiResponse{
		Code:    http.StatusOK,
		Message: msg,
	})
}

func BadRequest(c *gin.Context, message ...string) {
	msg := "bad request"
	if len(message) > 0 {
		msg = message[0]
	}

	c.JSON(http.StatusBadRequest, ApiResponse{
		Code:    http.StatusBadRequest,
		Message: msg,
	})
}

func NotFound(c *gin.Context, message ...string) {
	msg := "not found"
	if len(message) > 0 {
		msg = message[0]
	}

	c.JSON(http.StatusNotFound, ApiResponse{
		Code:    http.StatusNotFound,
		Message: msg,
	})
}

func InternalServerError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, ApiResponse{
		Code:    http.StatusInternalServerError,
		Message: "internal server error",
	})
}
