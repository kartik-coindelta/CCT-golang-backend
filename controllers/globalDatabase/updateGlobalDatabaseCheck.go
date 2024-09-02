package controllers

import (
	"CCT-GOLANG-BACKEND/db"
	"CCT-GOLANG-BACKEND/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UpdateGlobalDatabaseCheck handles the HTTP request to update an globalDatabase check by ID
func UpdateGlobalDatabaseCheck(c *gin.Context) {
	globalDatabaseID := c.Param("id")

	// Bind JSON body to a map (or a struct that matches the update fields)
	var updateData map[string]interface{}
	if err := c.BindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Get the globalDatabaseCheck collection
	collection := db.GetCollection("globaldatabasechecks")

	// Update the item in the database
	result, err := db.UpdateItem(globalDatabaseID, updateData, collection)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update the GlobalDatabase."})
		return
	}

	// Decode the updated item into a struct or map
	var updatedGlobalDatabaseCheck models.GlobalDatabaseCheck
	if err := result.Decode(&updatedGlobalDatabaseCheck); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode updated GlobalDatabase."})
		return
	}

	c.JSON(http.StatusOK, updatedGlobalDatabaseCheck)
}
