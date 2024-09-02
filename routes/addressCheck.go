package routes

import (
	controllers "CCT-GOLANG-BACKEND/controllers/address"

	"github.com/gin-gonic/gin"
)

// UserVerificationRequestRoutes sets up the routes for the user verification requests
func AddressCheck(router *gin.Engine) {
	router.POST("/addressCheck/link", controllers.CreateAddressCheck)
	router.GET("/addressCheck/link/:id", controllers.GetAddressCheckByVerificationID)
	router.PUT("/addressCheck/link/:id", controllers.UpdateAddressCheck)
	router.PUT("/address/status/:id", controllers.UpdateAddressCheckStatus)
}
