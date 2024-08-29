package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// DrugCheck represents the structure of the drugCheck collection in MongoDB
type DrugCheck struct {
	ID                        primitive.ObjectID  `bson:"_id,omitempty"`                       // MongoDB ObjectID
	UserVerificationRequestID primitive.ObjectID  `bson:"userVerificationRequestId,omitempty"` // ObjectId for the user verification request
	InefficiencyID            *primitive.ObjectID `bson:"inefficiencyId,omitempty"`            // ObjectId for inefficiency (can be null)
	Drug                      interface{}         `bson:"drug,omitempty"`                      // Use interface{} for flexibility with various data types
	CreatedAt                 time.Time           `bson:"createdAt,omitempty"`                 // Timestamp when the document was created
	UpdatedAt                 time.Time           `bson:"updatedAt,omitempty"`                 // Timestamp when the document was last updated
}
