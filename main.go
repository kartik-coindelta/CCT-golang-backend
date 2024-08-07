package main

import (
	controllers "CCT-GOLANG-BACKEND/controllers/auth"

	"github.com/gin-gonic/gin"
)

func main() {
	controllers.InitMongo() // Initialize MongoDB client

	router := gin.Default()
	router.POST("/signUp", controllers.SignUp)
	router.Run(":8080")
}
