package routes

import (
	controllers "CCT-GOLANG-BACKEND/controllers/experience"

	"github.com/gin-gonic/gin"
)

// UserVerificationRequestRoutes sets up the routes for the user verification requests
func ExperienceCheck(router *gin.Engine) {
	router.POST("/experienceCheck/link", controllers.CreateExperienceCheck)
}
