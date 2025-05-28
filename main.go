package main

import (
	"log"
	"votes/config"
	"votes/routes"
	"votes/utils/logger"
	"votes/utils/validate"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// load .env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	// logger
	logger.Init()

	// init db
	DB := config.InitDB()

	// init router
	r := gin.Default()

	// init translation validator
	validate.InitTranslator()

	// init routes
	routes.InitRoutes(r, DB)

	log.Println("server running at http://localhost:8080")
	r.Run(":8080")
}
