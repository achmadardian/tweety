package routes

import (
	"votes/config"
	"votes/handlers"
	"votes/repositories"
	"votes/services"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine, DB *config.Database) {
	// injection
	healthcheckHandler := handlers.NewHealthcheck()
	userRepo := repositories.NewUserRepository(DB)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userRepo, userService)

	// routes
	api := r.Group("api")
	{
		// healthcheck
		api.GET("/", healthcheckHandler.GetHealth)

		// user
		users := api.Group("/users")
		{
			users.GET("/", userHandler.GetUserAll)
			users.GET("/:id", userHandler.GetUserById)
			users.POST("/", userHandler.CreateUser)
			users.PATCH("/:id", userHandler.UpdateUser)
			users.DELETE("/:id", userHandler.DeleteUser)
		}
	}
}
