package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User represents the schema for users in MongoDB.
type User struct {
	ID                        primitive.ObjectID  `bson:"_id,omitempty" json:"_id,omitempty"`
	FirstName                 *string             `bson:"firstName,omitempty" json:"firstName,omitempty"`
	LastName                  *string             `bson:"lastName,omitempty" json:"lastName,omitempty"`
	CaseNumber                *int                `bson:"caseNumber,omitempty" json:"caseNumber,omitempty"`
	ClientName                *string             `bson:"clientName,omitempty" json:"clientName,omitempty"`
	Email                     *string             `bson:"email,omitempty" json:"email,omitempty"`
	UserName                  *string             `bson:"userName,omitempty" json:"userName,omitempty"`
	Password                  *string             `bson:"password,omitempty" json:"password,omitempty"`
	FatherName                *string             `bson:"fatherName,omitempty" json:"fatherName,omitempty"`
	MotherName                *string             `bson:"motherName,omitempty" json:"motherName,omitempty"`
	DateOfBirth               *time.Time          `bson:"dateOfBirth,omitempty" json:"dateOfBirth,omitempty"`
	Gender                    *string             `bson:"gender,omitempty" json:"gender,omitempty"`
	PhoneNumber               *int                `bson:"phoneNumber,omitempty" json:"phoneNumber,omitempty"`
	ImgURL                    *string             `bson:"imgURL,omitempty" json:"imgURL,omitempty"`
	UserWallet                *string             `bson:"userWallet,omitempty" json:"userWallet,omitempty"`
	Role                      *string             `bson:"role,omitempty" json:"role,omitempty"`
	CompanyID                 *primitive.ObjectID `bson:"companyId,omitempty" json:"companyId,omitempty"`
	CaseID                    *primitive.ObjectID `bson:"caseId,omitempty" json:"caseId,omitempty"`
	VerificationCode          *int                `bson:"verificationCode,omitempty" json:"verificationCode,omitempty"`
	Address                   *string             `bson:"address,omitempty" json:"address,omitempty"`
	Address1                  *string             `bson:"address1,omitempty" json:"address1,omitempty"`
	Address2                  *string             `bson:"address2,omitempty" json:"address2,omitempty"`
	City                      *string             `bson:"city,omitempty" json:"city,omitempty"`
	State                     *string             `bson:"state,omitempty" json:"state,omitempty"`
	Pincode                   *int                `bson:"pincode,omitempty" json:"pincode,omitempty"`
	Country                   *string             `bson:"country,omitempty" json:"country,omitempty"`
	VerificationCodeTimestamp *time.Time          `bson:"verificationCodeTimestamp,omitempty" json:"verificationCodeTimestamp,omitempty"`
	OTPBlockEndTime           *time.Time          `bson:"otpBlockEndTime,omitempty" json:"otpBlockEndTime,omitempty"`
	CreatedAt                 time.Time           `bson:"createdAt" json:"createdAt"`
	UpdatedAt                 time.Time           `bson:"updatedAt" json:"updatedAt"`
}
