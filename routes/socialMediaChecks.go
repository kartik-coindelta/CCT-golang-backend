package routes

import (
	controllers "CCT-GOLANG-BACKEND/controllers/socialMedia"

	"github.com/gin-gonic/gin"
)

// UserVerificationRequestRoutes sets up the routes for the user verification requests
func SocialMediaCheck(router *gin.Engine) {
	router.POST("/socialMediaCheck/link", controllers.CreateSocialMediaCheck)
}
