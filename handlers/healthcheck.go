package handlers

import (
	"votes/response"

	"github.com/gin-gonic/gin"
)

type Healthcheck struct{}

func NewHealthcheck() *Healthcheck {
	return &Healthcheck{}
}

func (h *Healthcheck) GetHealth(c *gin.Context) {
	response.Ok(c, "app is running well")
}
