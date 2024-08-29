package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SexOffenderCheck represents the structure of the sexOffenderCheck collection in MongoDB.
type SexOffenderCheck struct {
	ID                        primitive.ObjectID  `bson:"_id,omitempty"`                       // MongoDB ObjectID
	UserVerificationRequestID primitive.ObjectID  `bson:"userVerificationRequestId,omitempty"` // ObjectId for the user verification request
	InefficiencyID            *primitive.ObjectID `bson:"inefficiencyId,omitempty"`            // ObjectId for inefficiency (can be null)
	SexOffender               interface{}         `bson:"sexOffender,omitempty"`               // Flexible field for sex offender details
	CreatedAt                 time.Time           `bson:"createdAt,omitempty"`                 // Timestamp when the document was created
	UpdatedAt                 time.Time           `bson:"updatedAt,omitempty"`                 // Timestamp when the document was last updated
}
