package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// EducationCheck represents the structure of the educationCheck collection in MongoDB
type EducationCheck struct {
	ID                        primitive.ObjectID `bson:"_id,omitempty"`                       // MongoDB ObjectID
	UserVerificationRequestID primitive.ObjectID `bson:"userVerificationRequestId,omitempty"` // ObjectId for the user verification request
	InefficiencyID            primitive.ObjectID `bson:"inefficiencyId,omitempty"`            // ObjectId for inefficiency (default null in Mongoose schema)
	Education                 interface{}        `bson:"education,omitempty"`                 // A map or any other data structure to represent the Object type
	CreatedAt                 time.Time          `bson:"createdAt,omitempty"`                 // Timestamp when the document was created
	UpdatedAt                 time.Time          `bson:"updatedAt,omitempty"`                 // Timestamp when the document was last updated
}
