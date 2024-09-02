package controllers

import (
	"CCT-GOLANG-BACKEND/db"
	"CCT-GOLANG-BACKEND/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UpdateReputationalCheck handles the HTTP request to update an reputational check by ID
func UpdateReputationalCheck(c *gin.Context) {
	reputationalID := c.Param("id")

	// Bind JSON body to a map (or a struct that matches the update fields)
	var updateData map[string]interface{}
	if err := c.BindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Get the reputationalCheck collection
	collection := db.GetCollection("reputationalchecks")

	// Update the item in the database
	result, err := db.UpdateItem(reputationalID, updateData, collection)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update the Reputational."})
		return
	}

	// Decode the updated item into a struct or map
	var updatedReputationalCheck models.ReputationalCheck
	if err := result.Decode(&updatedReputationalCheck); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode updated Reputational."})
		return
	}

	c.JSON(http.StatusOK, updatedReputationalCheck)
}
