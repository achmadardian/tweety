package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "ok")
	})

	log.Printf("server running at http://localhost:8080 port")
	r.Run()
}
