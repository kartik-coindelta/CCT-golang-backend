package routes

import (
	"CCT-GOLANG-BACKEND/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Define routes
	router.POST("/signup", controllers.Signup)
	router.POST("/login", controllers.Login)
	router.POST("/register-company", controllers.RegisterCompany) // Ensure this route is defined
}
