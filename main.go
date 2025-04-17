package main

import (
	"log"
	"votes/config"
	"votes/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// init db
	DB := config.InitDB()

	// init router
	r := gin.Default()

	// init routes
	routes.InitRoutes(r, DB)

	log.Println("server running at http://localhost:8080")
	r.Run()
}
