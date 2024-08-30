package routes

import (
	"CCT-GOLANG-BACKEND/controllers"
	"CCT-GOLANG-BACKEND/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Define routes
	router.POST("/bca", controllers.SignUp)
	router.POST("/login", controllers.Login)
	router.POST("/company", controllers.RegisterCompany)
	router.POST("/user", controllers.RegisterUser)

	router.POST("/verifyOTP", middleware.VerifyOTP)
	router.POST("/userVerificationRequestID", controllers.SendInvite)
	

	

	SetupBCARoutes(router)
	CompanyRoutes(router)
	UserRoutes(router)

	UserVerificationRequestRoutes(router)
	AddressCheck(router)
	CertificatesCheck(router)
	CriminalCheck(router)
	DrugCheck(router)
	EducationCheck(router)
	ExperienceCheck(router)
	GlobalDatabaseCheck(router)
	IdentityCheck(router)
	ProfessionalLicenceCheck(router)
	ReferenceCheck(router)
	ReputationalCheck(router)
	SexOffenderCheck(router)
	SocialMediaCheck(router)
	UanCheck(router)
	VideoKycCheck(router)
	

}
