package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PaginatedResponse struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"page_size" form:"page_size"`
}

type ApiResponse struct {
	Code       int               `json:"code"`
	Message    string            `json:"message"`
	Data       interface{}       `json:"data,omitempty"`
	Errors     interface{}       `json:"errors,omitempty"`
	Pagination PaginatedResponse `json:"pagination"`
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

func OkPaginate(c *gin.Context, data interface{}, pagination *PaginatedResponse, message ...string) {
	msg := "success"
	if len(message) > 0 {
		msg = message[0]
	}

	c.JSON(http.StatusOK, ApiResponse{
		Code:    http.StatusOK,
		Message: msg,
		Data:    data,
		Pagination: PaginatedResponse{
			Page:     pagination.Page,
			PageSize: pagination.PageSize,
		},
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

func UnprocessableEntity(c *gin.Context, errors interface{}, message ...string) {
	msg := "unprocessable entity"
	if len(message) > 0 {
		msg = message[0]
	}

	c.JSON(http.StatusUnprocessableEntity, ApiResponse{
		Code:    http.StatusUnprocessableEntity,
		Message: msg,
		Errors:  errors,
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
