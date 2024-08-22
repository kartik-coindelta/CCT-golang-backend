package controllers

import (
	"CCT-GOLANG-BACKEND/db"
	"CCT-GOLANG-BACKEND/models"
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ListUserVerificationRequests handles listing all user verification requests with pagination
func ListUserVerificationRequests(c *gin.Context) {
	collection := db.GetCollection("userVerificationRequest")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	limitParam := c.DefaultQuery("limit", "10")
	pageParam := c.DefaultQuery("page", "1")

	limit, err := strconv.ParseInt(limitParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit value"})
		return
	}

	page, err := strconv.ParseInt(pageParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page value"})
		return
	}

	skip := (page - 1) * limit

	cursor, err := collection.Find(ctx, bson.M{}, options.Find().SetLimit(limit).SetSkip(skip))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(ctx)

	var requests []models.UserVerificationRequest
	if err := cursor.All(ctx, &requests); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, requests)
}
