package routes

import (
	"CCT-GOLANG-BACKEND/controllers/company"

	"github.com/gin-gonic/gin"
)

func CompanyRoutes(router *gin.Engine) {
	router.GET("company/:id", company.GetCompanyById)
	router.GET("company/getCompany", company.GetCompany)
}
