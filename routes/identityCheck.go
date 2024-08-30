package routes

import (
	controllers "CCT-GOLANG-BACKEND/controllers/identity"

	"github.com/gin-gonic/gin"
)

// UserVerificationRequestRoutes sets up the routes for the user verification requests
func IdentityCheck(router *gin.Engine) {
	router.POST("/identityCheck/link", controllers.CreateIdentityCheck)
}
