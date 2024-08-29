package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SocialMediaCheck represents the structure of the socialMediaCheck collection in MongoDB.
type SocialMediaCheck struct {
	ID                        primitive.ObjectID  `bson:"_id,omitempty"`                       // MongoDB ObjectID
	UserVerificationRequestID primitive.ObjectID  `bson:"userVerificationRequestId,omitempty"` // ObjectId for the user verification request
	InefficiencyID            *primitive.ObjectID `bson:"inefficiencyId,omitempty"`            // ObjectId for inefficiency (can be null)
	SocialMedia               interface{}         `bson:"socialMedia,omitempty"`               // Flexible field for social media details
	CreatedAt                 time.Time           `bson:"createdAt,omitempty"`                 // Timestamp when the document was created
	UpdatedAt                 time.Time           `bson:"updatedAt,omitempty"`                 // Timestamp when the document was last updated
}
