package routes

import (
	"github.com/achmadardian/tweety/config"
	"github.com/achmadardian/tweety/handlers"
	"github.com/achmadardian/tweety/middlewares"
	"github.com/achmadardian/tweety/repositories"
	"github.com/achmadardian/tweety/services"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine, DB *config.Database) {
	// injection
	healthcheckHandler := handlers.NewHealthcheck()

	// user
	userRepo := repositories.NewUserRepository(DB)
	userSvc := services.NewUserService(userRepo)
	userHandl := handlers.NewUserHandler(userSvc)

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
			auth.POST("/refresh-token", authHandl.RefreshToken)
		}

		api.Use(middlewares.Auth(authSvc))
		// user
		user := api.Group("/users")
		{
			user.GET("/me", userHandl.Me)
			user.PATCH("/me", userHandl.UpdateMe)
		}
	}
}
