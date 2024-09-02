package controllers

import (
	"CCT-GOLANG-BACKEND/db"
	"CCT-GOLANG-BACKEND/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UpdateAddressCheck handles the HTTP request to update an address check by ID
func UpdateAddressCheck(c *gin.Context) {
	addressID := c.Param("id")

	// Bind JSON body to a map (or a struct that matches the update fields)
	var updateData map[string]interface{}
	if err := c.BindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	

	// Get the addressCheck collection
	collection := db.GetCollection("addresschecks")

	// Update the item in the database
	result, err := db.UpdateItem(addressID, updateData, collection)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update the Address."})
		return
	}

	// Decode the updated item into a struct or map
	var updatedAddressCheck models.AddressCheck
	if err := result.Decode(&updatedAddressCheck); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode updated Address."})
		return
	}

	c.JSON(http.StatusOK, updatedAddressCheck)
}
