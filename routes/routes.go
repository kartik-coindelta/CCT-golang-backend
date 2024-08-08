package routes

import (
	"CCT-GOLANG-BACKEND/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/signup", controllers.Signup)

	return r
}
