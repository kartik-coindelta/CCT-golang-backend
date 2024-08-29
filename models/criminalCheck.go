package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CriminalCheck represents the structure of the criminalCheck collection in MongoDB.
type CriminalCheck struct {
	ID                        primitive.ObjectID  `bson:"_id,omitempty"`                       // MongoDB ObjectID
	UserVerificationRequestID primitive.ObjectID  `bson:"userVerificationRequestId,omitempty"` // ObjectId for the user verification request
	InefficiencyID            *primitive.ObjectID `bson:"inefficiencyId,omitempty"`            // ObjectId for inefficiency (can be null)
	Criminal                  interface{}         `bson:"criminal,omitempty"`                  // Flexible field for criminal details
	CreatedAt                 time.Time           `bson:"createdAt,omitempty"`                 // Timestamp when the document was created
	UpdatedAt                 time.Time           `bson:"updatedAt,omitempty"`                 // Timestamp when the document was last updated
}
