package controllers

import (
	"CCT-GOLANG-BACKEND/db"
	"CCT-GOLANG-BACKEND/middleware"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// SendInvite handles sending the invite email using the user ID
func SendInvite(c *gin.Context) {
	// Extract userVerificationRequestID from the request query
	userIDStr := c.Query("userVerificationRequestID")
	if userIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "userVerificationRequestID is required"})
		return
	}
	log.Printf("userId:%s", userIDStr)
	// Convert userVerificationRequestID to ObjectID
	userVerificationRequestID, err := primitive.ObjectIDFromHex(userIDStr)
	if err != nil {

		log.Println("Error converting userVerificationRequestID to ObjectID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid userVerificationRequestID"})
		return
	}

	// Fetch user verification request details from the database
	userVerificationRequestCollection := db.GetCollection("userverificationrequest")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var userVerificationRequestInfo struct {
		UserInfo struct {
			Email string `bson:"email" json:"email"`
		} `bson:"userInfo"`
		CompanyInfo struct {
			Name string `bson:"name" json:"name"`
		} `bson:"companyInfo"`
	}
	err = userVerificationRequestCollection.FindOne(ctx, bson.M{"_id": userVerificationRequestID}).Decode(&userVerificationRequestInfo)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Details not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching information"})
		return
	}

	// Extract email and companyName
	email := userVerificationRequestInfo.UserInfo.Email
	companyName := userVerificationRequestInfo.CompanyInfo.Name

	log.Printf("email: %s\n", email)
	log.Printf("companyName: %s", companyName)

	// // Construct the invite URL
	url := fmt.Sprintf("https://yourdomain.com/redirectToForm?id=%s", userIDStr)


	// Send invite email
	err = middleware.SendInvitation(email, url, companyName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error sending invite email"})
		return 
	}

	c.JSON(http.StatusOK, gin.H{"message": "Invite sent successfully"})
}
