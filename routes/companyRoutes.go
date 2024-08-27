package routes

import (
	"CCT-GOLANG-BACKEND/controllers/company"

	"github.com/gin-gonic/gin"
)

func CompanyRoutes(router *gin.Engine) {
	router.GET("/:id", company.GetCompanyById)
	router.GET("/", company.GetCompany)
	router.GET("/BCAId", company.GetCompanyByBCAId)
}
