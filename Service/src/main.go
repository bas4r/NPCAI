package main

import (
	"log"

	"github.com/basarrcan/NPCAI/routes"
	"github.com/basarrcan/NPCAI/services"
	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	r := gin.Default()

	config, err := LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}
	conn := services.ConnectDB(&config)

	// Define your API endpoints
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// Create New User Endpoint
	r.POST("/new-user", InjectDBToContextMiddleware(conn), routes.NewUserHandler)

	// Start the server
	r.Run()
}
