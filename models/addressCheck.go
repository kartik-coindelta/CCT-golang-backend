package models

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
    "time"
)

// AddressCheck represents the structure of the addressCheck collection in MongoDB.
type AddressCheck struct {
    ID                       primitive.ObjectID  `bson:"_id,omitempty"`                        // MongoDB ObjectID
    UserVerificationRequestID primitive.ObjectID `bson:"userVerificationRequestId,omitempty"`  // ObjectId for the user verification request
    InefficiencyID           *primitive.ObjectID `bson:"inefficiencyId,omitempty"`             // ObjectId for inefficiency (can be null)
    Address                  interface{}         `bson:"address,omitempty"`                    // Flexible field for various address formats
    CreatedAt                time.Time           `bson:"createdAt,omitempty"`                  // Timestamp when the document was created
    UpdatedAt                time.Time           `bson:"updatedAt,omitempty"`                  // Timestamp when the document was last updated
}
