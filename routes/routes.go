package routes

import (
	"votes/config"
	"votes/handlers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine, DB *config.Database) {
	// injection
	healthcheckHandler := handlers.NewHealthcheck()

	// routes
	api := r.Group("api")
	{
		// healthcheck
		api.GET("/", healthcheckHandler.GetHealth)
	}
}
