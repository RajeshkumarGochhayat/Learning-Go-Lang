package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rajesh/personal-skill-tracker/database"
	"github.com/rajesh/personal-skill-tracker/routes"
)

func main() {
	// Connect to DB
	database.Connect()

	// Initialize Gin router
	r := gin.Default()

	// Serve static files for Swagger UI
	r.Static("/swagger-ui", "./swagger-ui") // Serve Swagger UI static files

	// API Routes
	routes.RegisterRoutes(r)

	// Serve the Swagger YAML file
	r.GET("/swagger.yaml", func(c *gin.Context) {
		c.File("./docs/swagger.yaml")
	})

	// Run server
	r.Run(":8080")
}
