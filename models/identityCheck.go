package models

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
    "time"
)

// IdentityCheck represents the structure of the identityCheck collection in MongoDB.
type IdentityCheck struct {
    ID              primitive.ObjectID  `bson:"_id,omitempty"`               // MongoDB ObjectID
    UserVerificationRequestID primitive.ObjectID `bson:"userVerificationRequestId,omitempty"` // ObjectId for the user verification request
    InefficiencyID  *primitive.ObjectID `bson:"inefficiencyId,omitempty"`    // ObjectId for inefficiency (can be null)
    AdhaarCard      interface{}         `bson:"adhaarCard,omitempty"`        // Flexible field for Aadhaar card details
    PanCard         interface{}         `bson:"panCard,omitempty"`           // Flexible field for PAN card details
    DrivingLicence  interface{}         `bson:"drivingLicence,omitempty"`    // Flexible field for Driving Licence details
    Passport        interface{}         `bson:"passport,omitempty"`          // Flexible field for Passport details
    CreatedAt       time.Time           `bson:"createdAt,omitempty"`         // Timestamp when the document was created
    UpdatedAt       time.Time           `bson:"updatedAt,omitempty"`         // Timestamp when the document was last updated
}
