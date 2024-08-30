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

// CreateProfessionalLicenceCheck creates a new EducationCheck document in the database
func CreateProfessionalLicenceCheck(c *gin.Context) {
	var professionalLicenceCheck models.ProfessionalLicenseCheck

	// Bind JSON body to the struct
	if err := c.ShouldBindJSON(&professionalLicenceCheck); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Get the MongoDB collection
	collection := db.GetCollection("professionallicencechecks")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Insert the document into the collection without transaction
	result, err := db.CreateItem(collection, professionalLicenceCheck, ctx)
	if err != nil {
		log.Printf("Failed to create ProfessionalLicenceCheck: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create ProfessionalLicenceCheck"})
		return
	}

	// Return a success response
	c.JSON(http.StatusCreated, result)
}
