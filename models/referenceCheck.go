package models

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
    "time"
)

// ReferenceCheck represents the structure of the referenceCheck collection in MongoDB.
type ReferenceCheck struct {
    ID                      primitive.ObjectID  `bson:"_id,omitempty"`               // MongoDB ObjectID
    UserVerificationRequestID primitive.ObjectID `bson:"userVerificationRequestId,omitempty"` // ObjectId for the user verification request
    InefficiencyID          *primitive.ObjectID `bson:"inefficiencyId,omitempty"`    // ObjectId for inefficiency (can be null)
    Reference               interface{}         `bson:"reference,omitempty"`        // Flexible field for reference details
    CreatedAt               time.Time           `bson:"createdAt,omitempty"`         // Timestamp when the document was created
    UpdatedAt               time.Time           `bson:"updatedAt,omitempty"`         // Timestamp when the document was last updated
}
