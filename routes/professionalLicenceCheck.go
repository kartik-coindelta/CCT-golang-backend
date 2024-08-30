package routes

import (
	controllers "CCT-GOLANG-BACKEND/controllers/professionalLicence"

	"github.com/gin-gonic/gin"
)

// UserVerificationRequestRoutes sets up the routes for the user verification requests
func ProfessionalLicenceCheck(router *gin.Engine) {
	router.POST("/professionalLicenceCheck/link", controllers.CreateProfessionalLicenceCheck)
}
