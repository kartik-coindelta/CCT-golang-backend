package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ReputationalCheck represents the structure of the reputationalCheck collection in MongoDB.
type ReputationalCheck struct {
	ID                        primitive.ObjectID  `bson:"_id,omitempty"`                       // MongoDB ObjectID
	UserVerificationRequestID primitive.ObjectID  `bson:"userVerificationRequestId,omitempty"` // ObjectId for the user verification request
	InefficiencyID            *primitive.ObjectID `bson:"inefficiencyId,omitempty"`            // ObjectId for inefficiency (can be null)
	Reputational              interface{}         `bson:"reputational,omitempty"`              // Flexible field for reputational details
	CreatedAt                 time.Time           `bson:"createdAt,omitempty"`                 // Timestamp when the document was created
	UpdatedAt                 time.Time           `bson:"updatedAt,omitempty"`                 // Timestamp when the document was last updated
}
