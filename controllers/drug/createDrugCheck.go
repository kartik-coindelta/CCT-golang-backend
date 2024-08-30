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

// CreateDrugCheck creates a new EducationCheck document in the database
func CreateDrugCheck(c *gin.Context) {
	var drugCheck models.DrugCheck

	// Bind JSON body to the struct
	if err := c.ShouldBindJSON(&drugCheck); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Get the MongoDB collection
	collection := db.GetCollection("drugchecks")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Insert the document into the collection without transaction
	result, err := db.CreateItem(collection, drugCheck, ctx)
	if err != nil {
		log.Printf("Failed to create DrugCheck: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create DrugCheck"})
		return
	}

	// Return a success response
	c.JSON(http.StatusCreated, result)
}
