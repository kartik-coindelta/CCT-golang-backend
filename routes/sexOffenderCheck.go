package routes

import (
	controllers "CCT-GOLANG-BACKEND/controllers/sexOffender"

	"github.com/gin-gonic/gin"
)

// UserVerificationRequestRoutes sets up the routes for the user verification requests
func SexOffenderCheck(router *gin.Engine) {
	router.POST("/sexOffenderCheck/link", controllers.CreateSexOffenderCheck)
}
