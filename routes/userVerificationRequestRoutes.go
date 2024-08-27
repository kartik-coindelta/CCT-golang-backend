package routes

import (
	controllers "CCT-GOLANG-BACKEND/controllers/userVerificationRequest"

	"github.com/gin-gonic/gin"
)

// UserVerificationRequestRoutes sets up the routes for the user verification requests
func UserVerificationRequestRoutes(router *gin.Engine) {
	userVerificationRequestGroup := router.Group("ss")
	{
		userVerificationRequestGroup.POST("/", controllers.CreateUserVerificationRequest)
		userVerificationRequestGroup.GET("/:id", controllers.GetUserVerificationRequest)
		userVerificationRequestGroup.PUT("/:id", controllers.UpdateUserVerificationRequest)
		userVerificationRequestGroup.DELETE("/:id", controllers.DeleteUserVerificationRequest)
		userVerificationRequestGroup.GET("/", controllers.ListUserVerificationRequests)

		userVerificationRequestGroup.GET("/CompanyId", controllers.GetCompanyById)
	}
}
