package routes

import (
	controllers "CCT-GOLANG-BACKEND/controllers/drug"

	"github.com/gin-gonic/gin"
)

// UserVerificationRequestRoutes sets up the routes for the user verification requests
func DrugCheck(router *gin.Engine) {
	router.POST("/drugCheck/link", controllers.CreateDrugCheck)
}
