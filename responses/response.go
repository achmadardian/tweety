package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Errors  any    `json:"errors,omitempty"`
}

const (
	MsgSuccess             = "ok"
	MsgCreated             = "created"
	MsgBadRequest          = "bad request"
	MsgNotFound            = "not found"
	MsgUnprocessableEntity = "unprocessable entity"
	MsgUnauthorized        = "unauthorized"
	MsgDeleted             = "deleted"
	MsgInternalServerError = "internal server error"
)

func Ok(c *gin.Context, data any, message ...string) {
	msg := MsgSuccess
	if len(message) > 0 {
		msg = message[0]
	}

	c.JSON(http.StatusOK, ApiResponse{
		Code:    http.StatusOK,
		Message: msg,
		Data:    data,
	})
}

func Created(c *gin.Context, data any, message ...string) {
	msg := MsgCreated
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
	msg := MsgDeleted
	if len(message) > 0 {
		msg = message[0]
	}

	c.JSON(http.StatusOK, ApiResponse{
		Code:    http.StatusOK,
		Message: msg,
	})
}

func BadRequest(c *gin.Context, message ...string) {
	msg := MsgBadRequest
	if len(message) > 0 {
		msg = message[0]
	}

	c.JSON(http.StatusBadRequest, ApiResponse{
		Code:    http.StatusBadRequest,
		Message: msg,
	})
}

func Unauthorized(c *gin.Context, message ...string) {
	msg := MsgUnauthorized
	if len(message) > 0 {
		msg = message[0]
	}

	c.JSON(http.StatusUnauthorized, ApiResponse{
		Code:    http.StatusUnauthorized,
		Message: msg,
	})
}

func UnprocessableEntity(c *gin.Context, errors any, message ...string) {
	msg := MsgUnprocessableEntity
	if len(message) > 0 {
		msg = message[0]
	}

	c.JSON(http.StatusUnprocessableEntity, ApiResponse{
		Code:    http.StatusUnprocessableEntity,
		Message: msg,
		Errors:  errors,
	})
}

func UnprocessableEntityEmpty(c *gin.Context, message ...string) {
	msg := MsgUnprocessableEntity
	if len(message) > 0 {
		msg = message[0]
	}

	c.JSON(http.StatusUnprocessableEntity, ApiResponse{
		Code:    http.StatusUnprocessableEntity,
		Message: msg,
		Errors: map[string]string{
			"_": "at least one field must be filled",
		},
	})
}

func UnprocessableEntityMalformedJSON(c *gin.Context, message ...string) {
	msg := MsgUnprocessableEntity
	if len(message) > 0 {
		msg = message[0]
	}

	c.JSON(http.StatusUnprocessableEntity, ApiResponse{
		Code:    http.StatusUnprocessableEntity,
		Message: msg,
		Errors: map[string]string{
			"body": "invalid or malformed JSON",
		},
	})
}

func NotFound(c *gin.Context, message ...string) {
	msg := MsgNotFound
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
		Message: MsgInternalServerError,
	})
}
