package routes

import (
	controllers "CCT-GOLANG-BACKEND/controllers/criminal"

	"github.com/gin-gonic/gin"
)

// UserVerificationRequestRoutes sets up the routes for the user verification requests
func CriminalCheck(router *gin.Engine) {
	router.POST("/criminalChecks/link", controllers.CreateCriminalCheck)
}
