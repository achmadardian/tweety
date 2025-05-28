package routes

import (
	"votes/config"
	"votes/handlers"
	"votes/middlewares"
	"votes/repositories"
	"votes/services"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine, DB *config.Database) {
	// injection
	healthcheckHandler := handlers.NewHealthcheck()

	// user
	userRepo := repositories.NewUserRepository(DB)
	userSvc := services.NewUserService(userRepo)

	// auth
	authSvc := services.NewAuthService(userSvc)
	authHandl := handlers.NewAuthHandler(authSvc)

	// routes
	api := r.Group("api")
	{
		// middleware
		// logger
		api.Use(middlewares.Logger())

		// healthcheck
		api.GET("/", healthcheckHandler.GetHealth)

		// auth
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandl.Register)
			auth.POST("/login", authHandl.Login)
		}
	}
}
