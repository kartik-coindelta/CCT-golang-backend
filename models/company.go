package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Company represents the company schema in the MongoDB database
type Company struct {
	ID                        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name                      *string            `bson:"name,omitempty" json:"name,omitempty"`
	Email                     *string            `bson:"email,omitempty" json:"email,omitempty"`
	UserName                  *string            `bson:"userName,omitempty" json:"userName,omitempty"`
	Password                  *string            `bson:"password,omitempty" json:"password,omitempty"`
	PhoneNumber               *int               `bson:"phoneNumber,omitempty" json:"phoneNumber,omitempty"`
	Address                   *string            `bson:"address,omitempty" json:"address,omitempty"`
	Line2                     *string            `bson:"line2,omitempty" json:"line2,omitempty"`
	Line1                     *string            `bson:"line1,omitempty" json:"line1,omitempty"`
	Zipcode                   *int               `bson:"zipcode,omitempty" json:"zipcode,omitempty"`
	State                     *string            `bson:"state,omitempty" json:"state,omitempty"`
	City                      *string            `bson:"city,omitempty" json:"city,omitempty"`
	Country                   *string            `bson:"country,omitempty" json:"country,omitempty"`
	BCAID                     primitive.ObjectID `bson:"bcaId" json:"userId"`
	WebsiteLink               *string            `bson:"websiteLink,omitempty" json:"websiteLink,omitempty"`
	NoOfEmployees             *int               `bson:"noOfEmployees,omitempty" json:"noOfEmployees,omitempty"`
	CompanyRegistrationNumber *int               `bson:"companyRegistrationNumber,omitempty" json:"companyRegistrationNumber,omitempty"`
	UserWallet                *string            `bson:"userWallet,omitempty" json:"userWallet,omitempty"`
	DiscountPrice             *string            `bson:"discountPrice,omitempty" json:"discountPrice,omitempty"`
	LogoURL                   *string            `bson:"logoURL,omitempty" json:"logoURL,omitempty"`
	Role                      *string            `bson:"role,omitempty" json:"role,omitempty"`
	AvailableChecks           []string           `bson:"availableChecks,omitempty" json:"availableChecks,omitempty"`
	VerificationCode          *int               `bson:"verificationCode,omitempty" json:"verificationCode,omitempty"`
	VerificationCodeTimestamp *time.Time         `bson:"verificationCodeTimestamp,omitempty" json:"verificationCodeTimestamp,omitempty"`
	PrePayment                *bool              `bson:"prePayment,omitempty" json:"prePayment,omitempty"`
	CreatedAt                 time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt                 time.Time          `bson:"updatedAt" json:"updatedAt"`
}
