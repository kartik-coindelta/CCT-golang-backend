package controllers

import (
	"CCT-GOLANG-BACKEND/db"
	"CCT-GOLANG-BACKEND/models"
	"context"
	"errors"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UpdateUserVerificationRequest handles updating an existing user verification request by its ID
func UpdateUserVerificationRequest(c *gin.Context) {
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var updateRequest map[string]interface{}
	if err := c.ShouldBindJSON(&updateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate that the input keys match the UserVerificationRequest model
	if err := validateInputKeys(updateRequest, models.UserVerificationRequest{}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collection := db.GetCollection("userVerificationRequest")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	update := bson.M{
		"$set": updateRequest,
	}

	_, err = collection.UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User verification request updated successfully"})
}

// validateInputKeys checks that all keys in the input map match the struct fields
func validateInputKeys(input map[string]interface{}, model interface{}) error {
	modelType := reflect.TypeOf(model)
	validFields := make(map[string]bool)

	// Collect all valid field names (JSON tags) from the model
	for i := 0; i < modelType.NumField(); i++ {
		field := modelType.Field(i)
		jsonTag := field.Tag.Get("json")
		if jsonTag == "" {
			jsonTag = strings.ToLower(field.Name)
		} else {
			jsonTag = strings.Split(jsonTag, ",")[0]
		}
		validFields[jsonTag] = true
	}

	// Check that all input keys are valid
	for key := range input {
		if !validFields[key] {
			return errors.New("invalid key: " + key)
		}
	}

	return nil
}
