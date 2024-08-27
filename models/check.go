package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Check represents the schema for a check in MongoDB.
type Check struct {
	ID                        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserVerificationRequestID primitive.ObjectID `bson:"userVerificationRequestId,omitempty" json:"userVerificationRequestId,omitempty"`
	CheckList                 bson.M             `bson:"checkList,omitempty" json:"checkList,omitempty"`
	CreatedAt                 time.Time          `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt                 time.Time          `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}
