package routes

import (
	controllers "CCT-GOLANG-BACKEND/controllers/education"

	"github.com/gin-gonic/gin"
)

// UserVerificationRequestRoutes sets up the routes for the user verification requests
func EducationCheck(router *gin.Engine) {
	router.POST("/educationCheck/link", controllers.CreateEducationCheck)
}
