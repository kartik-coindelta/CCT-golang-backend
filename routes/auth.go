package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// Mockup of controller functions
// func login(c *gin.Context)                {}
// func getRefreshToken(c *gin.Context)      {}
// func refreshToken(c *gin.Context)         {}
// func fileUpload(c *gin.Context)           {}
// func deleteFile(c *gin.Context)           {}
// func verifyLoginOtp(c *gin.Context)       {}
// func sendInvite(c *gin.Context)           {}
func signUp(c *gin.Context) {}

// func sendVerificationLink(c *gin.Context) {}
// func generateLink(c *gin.Context)         {}
// func multiUploadFile(c *gin.Context)      {}

// Mockup of validator functions
// func validateLogin() gin.HandlerFunc {
// 	return func(c *gin.Context) {}
// }

// func validateRefreshToken() gin.HandlerFunc {
// 	return func(c *gin.Context) {}
// }

func main() {
	router := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	// Auth middleware placeholder
	// requireAuth := func(c *gin.Context) {
	// 	// Implement JWT authentication logic here
	// 	c.Next()
	// }

	api := router.Group("/api")
	{
		// api.POST("/sendInvite", requireAuth, sendInvite)
		// api.POST("/verifyLogin", verifyLoginOtp)
		// api.POST("/sendVerificationLink", sendVerificationLink)
		// api.POST("/generateLink", generateLink)
		// api.GET("/token", requireAuth, getRefreshToken)
		// api.POST("/refreshToken", validateRefreshToken(), refreshToken)
		// api.POST("/login", login)
		api.POST("/signUp", signUp)
		// api.POST("/upload", fileUpload)
		// api.POST("/multiUpload", multiUploadFile)
		// api.POST("/deleteFile", deleteFile)
	}

	router.Run(":8080")
}
