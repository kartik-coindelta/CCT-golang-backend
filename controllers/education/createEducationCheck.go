package controllers

import (
	"CCT-GOLANG-BACKEND/db"
	"CCT-GOLANG-BACKEND/models"
	"encoding/json"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// Initialize your MongoDB collection
var educationCheckCollection *mongo.Collection // Initialize this with your MongoDB collection

// CreateEducationCheck handles the creation of a new EducationCheck document.
func CreateEducationCheck(w http.ResponseWriter, r *http.Request) {
	var data models.EducationCheck

	// Decode the request body into the EducationCheck struct
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Set timestamps if needed
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()

	// Call CreateItem to insert the data into MongoDB
	result, err := db.CreateItem(educationCheckCollection, data, nil)
	if err != nil {
		http.Error(w, "Failed to create the EducationCheck", http.StatusInternalServerError)
		return
	}

	// Respond with the inserted ID and a success message
	response := map[string]interface{}{
		"insertedID": result.InsertedID,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
