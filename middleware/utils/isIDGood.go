package utils

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// IsIDGood checks if the given ID is a valid MongoDB ObjectId.
func IsIDGood(id string) (primitive.ObjectID, map[string]interface{}) {
	// Try to parse the ID as a MongoDB ObjectId
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return primitive.ObjectID{}, BuildErrObject(422, "ID_MALFORMED")
	}
	return objectID, nil
}
