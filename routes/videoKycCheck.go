package routes

import (
	controllers "CCT-GOLANG-BACKEND/controllers/videoKyc"

	"github.com/gin-gonic/gin"
)

// UserVerificationRequestRoutes sets up the routes for the user verification requests
func VideoKycCheck(router *gin.Engine) {
	router.POST("/videoKycCheck/link", controllers.CreateVideoKycCheck)
}
