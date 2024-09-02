package routes

import (
	controllers "CCT-GOLANG-BACKEND/controllers/reputational"

	"github.com/gin-gonic/gin"
)

// UserVerificationRequestRoutes sets up the routes for the user verification requests
func ReputationalCheck(router *gin.Engine) {
	router.POST("/reputationalCheck/link", controllers.CreateReputationalCheck)
	router.GET("/reputationalCheck/:id", controllers.GetReputationalCheckByVerificationID)
}
