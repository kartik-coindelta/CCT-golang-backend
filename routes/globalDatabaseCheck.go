package routes

import (
	controllers "CCT-GOLANG-BACKEND/controllers/globalDatabase"

	"github.com/gin-gonic/gin"
)

// UserVerificationRequestRoutes sets up the routes for the user verification requests
func GlobalDatabaseCheck(router *gin.Engine) {
	router.POST("/globalDatabaseCheck/link", controllers.CreateGlobalDatabaseCheck)
}
