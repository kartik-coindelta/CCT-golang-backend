package routes

import (
	"CCT-GOLANG-BACKEND/controllers/BCA"

	"github.com/gin-gonic/gin"
)

func SetupBCARoutes(router *gin.Engine) {
	// Define BCA-specific routes
	router.GET("/bca/:id", BCA.GetBCAByID)
	router.GET("/bca/getBCA", BCA.GetBCA)
}
