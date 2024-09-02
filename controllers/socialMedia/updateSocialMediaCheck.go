package controllers

import (
	"CCT-GOLANG-BACKEND/db"
	"CCT-GOLANG-BACKEND/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UpdateSocialMediaCheck handles the HTTP request to update an socialMedia check by ID
func UpdateSocialMediaCheck(c *gin.Context) {
	socialMediaID := c.Param("id")

	// Bind JSON body to a map (or a struct that matches the update fields)
	var updateData map[string]interface{}
	if err := c.BindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Get the socialMediaCheck collection
	collection := db.GetCollection("socialmediachecks")

	// Update the item in the database
	result, err := db.UpdateItem(socialMediaID, updateData, collection)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update the SocialMedia."})
		return
	}

	// Decode the updated item into a struct or map
	var updatedSocialMediaCheck models.SocialMediaCheck
	if err := result.Decode(&updatedSocialMediaCheck); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode updated SocialMedia."})
		return
	}

	c.JSON(http.StatusOK, updatedSocialMediaCheck)
}
