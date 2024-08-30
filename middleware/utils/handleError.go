package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

// HandleError handles error by printing to console in development env
// and builds and sends an error response.
func HandleError(w http.ResponseWriter, err error, code int) {
	if os.Getenv("ENV") != "test" {
		log.Println(err)
	}

	// Define default error message
	errorMessage := "An unhandled error occurred."

	// Check if err is a custom error type and extract message if available
	if customErr, ok := err.(CustomError); ok {
		errorMessage = customErr.Message
		code = customErr.Code
	}

	// Set response headers and status code
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	// Write JSON response
	response := map[string]interface{}{
		"errors": map[string]string{
			"msg": errorMessage,
		},
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Println("Error encoding response:", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// CustomError defines a custom error type with a code and message.
type CustomError struct {
	Code    int
	Message string
}

func (e CustomError) Error() string {
	return e.Message
}
