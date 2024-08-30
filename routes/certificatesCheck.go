package routes

import (
	controllers "CCT-GOLANG-BACKEND/controllers/certificates"

	"github.com/gin-gonic/gin"
)

// UserVerificationRequestRoutes sets up the routes for the user verification requests
func CertificatesCheck(router *gin.Engine) {
	router.POST("/certificatesCheck/link", controllers.CreateCertificatesCheck)
}
