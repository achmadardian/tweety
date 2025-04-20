package utils

import (
	"votes/response"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetQueryParamPagination(c *gin.Context) (*response.PaginatedResponse, error) {
	var page response.PaginatedResponse
	err := c.BindQuery(&page)
	if err != nil {
		return nil, err
	}

	return &page, nil
}

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
