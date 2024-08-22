package controllers

import (
	"CCT-GOLANG-BACKEND/db"
	"CCT-GOLANG-BACKEND/models"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateUserVerificationRequest handles the creation of a new user verification request
func CreateUserVerificationRequest(c *gin.Context) {
	var request models.UserVerificationRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collection := db.GetCollection("userVerificationRequest")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	request.ID = primitive.NewObjectID()
	_, err := collection.InsertOne(ctx, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User verification request created successfully"})
}
