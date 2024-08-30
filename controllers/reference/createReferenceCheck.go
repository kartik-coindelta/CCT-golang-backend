package controllers

import (
	"CCT-GOLANG-BACKEND/db"
	"CCT-GOLANG-BACKEND/models"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateReferenceCheck creates a new EducationCheck document in the database
func CreateReferenceCheck(c *gin.Context) {
	var referenceCheck models.ReferenceCheck

	// Bind JSON body to the struct
	if err := c.ShouldBindJSON(&referenceCheck); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Get the MongoDB collection
	collection := db.GetCollection("referencechecks")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Insert the document into the collection without transaction
	result, err := db.CreateItem(collection, referenceCheck, ctx)
	if err != nil {
		log.Printf("Failed to create ReferenceCheck: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create ReferenceCheck"})
		return
	}

	// Return a success response
	c.JSON(http.StatusCreated, result)
}
