package controllers

import (
	"CCT-GOLANG-BACKEND/db"
	"CCT-GOLANG-BACKEND/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UpdateCertificatesCheck handles the HTTP request to update an certificates check by ID
func UpdateCertificatesCheck(c *gin.Context) {
	certificatesID := c.Param("id")

	// Bind JSON body to a map (or a struct that matches the update fields)
	var updateData map[string]interface{}
	if err := c.BindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Get the certificatesCheck collection
	collection := db.GetCollection("certificateschecks")

	// Update the item in the database
	result, err := db.UpdateItem(certificatesID, updateData, collection)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update the Certificate."})
		return
	}

	// Decode the updated item into a struct or map
	var updatedCertificatesCheck models.CertificatesCheck
	if err := result.Decode(&updatedCertificatesCheck); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode updated Certificate."})
		return
	}

	c.JSON(http.StatusOK, updatedCertificatesCheck)
}
