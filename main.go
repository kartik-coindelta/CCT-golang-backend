package main

import (
	"CCT-GOLANG-BACKEND/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Set up routes
	routes.SetupRoutes(router)

	// Start the server
	router.Run(":8080")
}
