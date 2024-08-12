package routes

import (
	"CCT-GOLANG-BACKEND/controllers/user"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.GET("user/:id", user.GetUserById)
	router.GET("user/getUser", user.GetUser)
}
