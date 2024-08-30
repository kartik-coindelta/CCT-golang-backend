package routes

import (
	controllers "CCT-GOLANG-BACKEND/controllers/reference"

	"github.com/gin-gonic/gin"
)

// UserVerificationRequestRoutes sets up the routes for the user verification requests
func ReferenceCheck(router *gin.Engine) {
	router.POST("/referenceCheck/link", controllers.CreateReferenceCheck)
}
