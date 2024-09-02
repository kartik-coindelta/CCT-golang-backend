package controllers

import (
	"CCT-GOLANG-BACKEND/db"
	"CCT-GOLANG-BACKEND/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UpdateProfessionalLicenseCheck handles the HTTP request to update an professionalLicense check by ID
func UpdateProfessionalLicenseCheck(c *gin.Context) {
	professionalLicenseID := c.Param("id")

	// Bind JSON body to a map (or a struct that matches the update fields)
	var updateData map[string]interface{}
	if err := c.BindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Get the professionalLicenseCheck collection
	collection := db.GetCollection("professionallicensechecks")

	// Update the item in the database
	result, err := db.UpdateItem(professionalLicenseID, updateData, collection)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update the ProfessionalLicense."})
		return
	}

	// Decode the updated item into a struct or map
	var updatedProfessionalLicenseCheck models.ProfessionalLicenseCheck
	if err := result.Decode(&updatedProfessionalLicenseCheck); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode updated ProfessionalLicense."})
		return
	}

	c.JSON(http.StatusOK, updatedProfessionalLicenseCheck)
}
