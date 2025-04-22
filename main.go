package main

import (
	"log"
	"votes/config"
	"votes/routes"
	"votes/utils"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// load .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("error load .env:", err)
	}

	// init db
	DB := config.InitDB()

	// init router
	r := gin.Default()

	// init translation validator
	utils.InitTransValidator()

	// init routes
	routes.InitRoutes(r, DB)

	log.Println("server running at http://localhost:8080")
	r.Run(":8080")
}
