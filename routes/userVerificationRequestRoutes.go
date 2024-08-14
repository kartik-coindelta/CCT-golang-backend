package routes

import (
	"CCT-GOLANG-BACKEND/controllers"

	"github.com/gin-gonic/gin"
)

// UserVerificationRequestRoutes sets up the routes for the user verification requests
func UserVerificationRequestRoutes(router *gin.Engine) {
	userVerificationRequestGroup := router.Group("/userVerificationRequest")
	{
		userVerificationRequestGroup.POST("/", controllers.CreateUserVerificationRequest)
		userVerificationRequestGroup.GET("/:id", controllers.GetUserVerificationRequest)
		userVerificationRequestGroup.PUT("/:id", controllers.UpdateUserVerificationRequest)
		userVerificationRequestGroup.DELETE("/:id", controllers.DeleteUserVerificationRequest)
		userVerificationRequestGroup.GET("/", controllers.ListUserVerificationRequests)
	}
}
