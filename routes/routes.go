package routes

import (
	"CCT-GOLANG-BACKEND/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Define routes
	router.POST("/signup", controllers.SignUp)
	router.POST("/login", controllers.Login)
	router.POST("/loginCompany", controllers.Login)
	router.POST("/loginUser", controllers.Login)
	router.POST("/registerCompany", controllers.RegisterCompany)
	router.POST("/registerUser", controllers.RegisterUser)
}
