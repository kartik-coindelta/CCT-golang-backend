package controllers

import (
	"CCT-GOLANG-BACKEND/db"
	"CCT-GOLANG-BACKEND/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UpdateDrugCheck handles the HTTP request to update an drug check by ID
func UpdateDrugCheck(c *gin.Context) {
	drugID := c.Param("id")

	// Bind JSON body to a map (or a struct that matches the update fields)
	var updateData map[string]interface{}
	if err := c.BindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Get the drugCheck collection
	collection := db.GetCollection("drugchecks")

	// Update the item in the database
	result, err := db.UpdateItem(drugID, updateData, collection)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update the Drug."})
		return
	}

	// Decode the updated item into a struct or map
	var updatedDrugCheck models.DrugCheck
	if err := result.Decode(&updatedDrugCheck); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode updated Drug."})
		return
	}

	c.JSON(http.StatusOK, updatedDrugCheck)
}
