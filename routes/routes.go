package routes

import (
	"votes/config"
	"votes/handlers"
	"votes/repositories"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine, DB *config.Database) {
	// injection
	healthcheckHandler := handlers.NewHealthcheck()
	userRepo := repositories.NewUserRepository(DB)
	userHandler := handlers.NewUserHandler(userRepo)

	// routes
	api := r.Group("api")
	{
		// healthcheck
		api.GET("/", healthcheckHandler.GetHealth)

		// user
		users := api.Group("/users")
		{
			users.GET("/", userHandler.GetUserAll)
		}
	}
}
