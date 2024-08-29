package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ProfessionalLicenseCheck represents the structure of the professionalLicenseCheck collection in MongoDB.
type ProfessionalLicenseCheck struct {
	ID                        primitive.ObjectID  `bson:"_id,omitempty"`                       // MongoDB ObjectID
	UserVerificationRequestID primitive.ObjectID  `bson:"userVerificationRequestId,omitempty"` // ObjectId for the user verification request
	InefficiencyID            *primitive.ObjectID `bson:"inefficiencyId,omitempty"`            // ObjectId for inefficiency (can be null)
	ProfessionalLicense       interface{}         `bson:"professionalLicense,omitempty"`       // Flexible field for professional license details
	CreatedAt                 time.Time           `bson:"createdAt,omitempty"`                 // Timestamp when the document was created
	UpdatedAt                 time.Time           `bson:"updatedAt,omitempty"`                 // Timestamp when the document was last updated
}
