package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Document represents the document schema in the MongoDB database
type Document struct {
	DocumentHash string `bson:"documentHash" json:"documentHash"`
	URL          string `bson:"url" json:"url"`
}

// BCA represents the BCA schema in the MongoDB database
type BCA struct {
	ID                        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	BCAId                     primitive.ObjectID `bson:"bcaId,omitempty" json:"bcaId,omitempty"`
	Name                      *string            `bson:"name,omitempty" json:"name,omitempty"`
	FirstName                 *string            `bson:"firstName,omitempty" json:"firstName,omitempty"`
	LastName                  *string            `bson:"lastName,omitempty" json:"lastName,omitempty"`
	Email                     *string            `bson:"email,omitempty" json:"email,omitempty"`
	UserName                  *string            `bson:"userName,omitempty" json:"userName,omitempty"`
	Password                  *string            `bson:"password,omitempty" json:"password,omitempty"`
	PhoneNumber               *string            `bson:"phoneNumber,omitempty" json:"phoneNumber,omitempty"`
	Line2                     *string            `bson:"line2,omitempty" json:"line2,omitempty"`
	Line1                     *string            `bson:"line1,omitempty" json:"line1,omitempty"`
	Zipcode                   *int               `bson:"zipcode,omitempty" json:"zipcode,omitempty"`
	CompanyRegistrationNumber *int               `bson:"companyRegistrationNumber,omitempty" json:"companyRegistrationNumber,omitempty"`
	WebsiteLink               *string            `bson:"websiteLink,omitempty" json:"websiteLink,omitempty"`
	NoOfEmployees             *int               `bson:"noOfEmployees,omitempty" json:"noOfEmployees,omitempty"`
	UserWallet                *string            `bson:"userWallet,omitempty" json:"userWallet,omitempty"`
	SupportingDocuments       []Document         `bson:"supportingDocuments,omitempty" json:"supportingDocuments,omitempty"`
	LogoURL                   *string            `bson:"logoURL,omitempty" json:"logoURL,omitempty"`
	Status                    *string            `bson:"status,omitempty" json:"status,omitempty"`
	Role                      *string            `bson:"role,omitempty" json:"role,omitempty"`
	VendorName                *string            `bson:"vendorName,omitempty" json:"vendorName,omitempty"`
	ManagerName               *string            `bson:"managerName,omitempty" json:"managerName,omitempty"`
	Address                   *string            `bson:"address,omitempty" json:"address,omitempty"`
	City                      *string            `bson:"city,omitempty" json:"city,omitempty"`
	State                     *string            `bson:"state,omitempty" json:"state,omitempty"`
	GST                       *string            `bson:"gst,omitempty" json:"gst,omitempty"`
	Country                   *string            `bson:"country,omitempty" json:"country,omitempty"`
	AdditionalRemark          *string            `bson:"additionalRemark,omitempty" json:"additionalRemark,omitempty"`
	HasStaffAccess            *bool              `bson:"hasStaffAccess,omitempty" json:"hasStaffAccess,omitempty"`
	VerificationCode          *int               `bson:"verificationCode,omitempty" json:"verificationCode,omitempty"`
	VerificationCodeTimestamp *time.Time         `bson:"verificationCodeTimestamp,omitempty" json:"verificationCodeTimestamp,omitempty"`
	OtpBlockEndTime           *time.Time         `bson:"otpBlockEndTime,omitempty" json:"otpBlockEndTime,omitempty"`
	SES_URL                   *string            `bson:"ses_url,omitempty" json:"ses_url,omitempty"`                       // New field
	EmailBannerImage          *string            `bson:"email_banner_image,omitempty" json:"email_banner_image,omitempty"` // New field
	CompanyTitle              *string            `bson:"company_title,omitempty" json:"company_title,omitempty"`           // New field
	OtpAttempts               *int               `bson:"otpAttempts,omitempty" json:"otpAttempts,omitempty"`               // New field
	CreatedAt                 time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt                 time.Time          `bson:"updatedAt" json:"updatedAt"`
}
