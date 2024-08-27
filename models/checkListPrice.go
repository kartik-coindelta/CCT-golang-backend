package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CheckListPrice represents the schema for checklist prices in MongoDB.
type CheckListPrice struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Education      *int               `bson:"Education,omitempty" json:"Education,omitempty"`
	Experience     *int               `bson:"Experience,omitempty" json:"Experience,omitempty"`
	Certificate    *int               `bson:"Certificate,omitempty" json:"Certificate,omitempty"`
	Address        *int               `bson:"Address,omitempty" json:"Address,omitempty"`
	Identity       *int               `bson:"Identity,omitempty" json:"Identity,omitempty"`
	Court          *int               `bson:"Court,omitempty" json:"Court,omitempty"`
	Reference      *int               `bson:"Reference,omitempty" json:"Reference,omitempty"`
	Police         *int               `bson:"Police,omitempty" json:"Police,omitempty"`
	Credit         *int               `bson:"Credit,omitempty" json:"Credit,omitempty"`
	Drug           *int               `bson:"Drug,omitempty" json:"Drug,omitempty"`
	VideoKyc       *int               `bson:"VideoKyc,omitempty" json:"VideoKyc,omitempty"`
	GlobalDatabase *int               `bson:"GlobalDatabase,omitempty" json:"GlobalDatabase,omitempty"`
	SexOffender    *int               `bson:"SexOffender,omitempty" json:"SexOffender,omitempty"`
	LastCompany    *int               `bson:"LastCompany,omitempty" json:"LastCompany,omitempty"`
	Last2Company   *int               `bson:"Last_2_Company,omitempty" json:"Last_2_Company,omitempty"`
	Last3Company   *int               `bson:"Last_3_Company,omitempty" json:"Last_3_Company,omitempty"`
	CreatedAt      time.Time          `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt      time.Time          `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}
