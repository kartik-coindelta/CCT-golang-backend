package routes

import (
	controllers "CCT-GOLANG-BACKEND/controllers/uan"

	"github.com/gin-gonic/gin"
)

// UserVerificationRequestRoutes sets up the routes for the user verification requests
func UanCheck(router *gin.Engine) {
	router.POST("/uanCheck/link", controllers.CreateUanCheck)
}
