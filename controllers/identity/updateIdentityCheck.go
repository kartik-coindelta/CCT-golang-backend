package controllers

import (
	"CCT-GOLANG-BACKEND/db"
	"CCT-GOLANG-BACKEND/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UpdateIdentityCheck handles the HTTP request to update an identity check by ID
func UpdateIdentityCheck(c *gin.Context) {
	identityID := c.Param("id")

	// Bind JSON body to a map (or a struct that matches the update fields)
	var updateData map[string]interface{}
	if err := c.BindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Get the identityCheck collection
	collection := db.GetCollection("identitychecks")

	// Update the item in the database
	result, err := db.UpdateItem(identityID, updateData, collection)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update the Identity."})
		return
	}

	// Decode the updated item into a struct or map
	var updatedIdentityCheck models.IdentityCheck
	if err := result.Decode(&updatedIdentityCheck); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode updated Identity."})
		return
	}

	c.JSON(http.StatusOK, updatedIdentityCheck)
}
