package controllers

import (
	"CCT-GOLANG-BACKEND/db"
	"CCT-GOLANG-BACKEND/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UpdateSexOffenderCheck handles the HTTP request to update an sexOffender check by ID
func UpdateSexOffenderCheck(c *gin.Context) {
	sexOffenderID := c.Param("id")

	// Bind JSON body to a map (or a struct that matches the update fields)
	var updateData map[string]interface{}
	if err := c.BindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Get the sexOffenderCheck collection
	collection := db.GetCollection("sexoffenderchecks")

	// Update the item in the database
	result, err := db.UpdateItem(sexOffenderID, updateData, collection)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update the SexOffender."})
		return
	}

	// Decode the updated item into a struct or map
	var updatedSexOffenderCheck models.SexOffenderCheck
	if err := result.Decode(&updatedSexOffenderCheck); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode updated SexOffender."})
		return
	}

	c.JSON(http.StatusOK, updatedSexOffenderCheck)
}
