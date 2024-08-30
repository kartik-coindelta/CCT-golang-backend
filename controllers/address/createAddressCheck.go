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

// CreateAddressCheck creates a new EducationCheck document in the database
func CreateAddressCheck(c *gin.Context) {
	var addressCheck models.AddressCheck

	// Bind JSON body to the struct
	if err := c.ShouldBindJSON(&addressCheck); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Get the MongoDB collection
	collection := db.GetCollection("addresschecks")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Printf("AddressCheck")
	// Insert the document into the collection without transaction
	result, err := db.CreateItem(collection, addressCheck, ctx)
	if err != nil {
		log.Printf("Failed to create AddressCheck: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create AddressCheck"})
		return
	}

	// Return a success response
	c.JSON(http.StatusCreated, result)
}
