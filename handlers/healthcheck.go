package handlers

import (
	"votes/responses"

	"github.com/gin-gonic/gin"
)

type Healthcheck struct{}

func NewHealthcheck() *Healthcheck {
	return &Healthcheck{}
}

func (h *Healthcheck) GetHealth(c *gin.Context) {
	responses.Ok(c, "app is running well")
}
