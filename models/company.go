package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Company represents the company schema in MongoDB.
type Company struct {
	ID                        primitive.ObjectID `bson:"_id,omitempty"`                       // MongoDB Object ID
	BCAId                     primitive.ObjectID `bson:"bcaId,omitempty"`                     // Reference to BCA ID
	Name                      string             `bson:"name,omitempty"`                      // Company name
	Email                     *string            `bson:"email,omitempty"`                     // Email address
	UserName                  *string            `bson:"userName,omitempty"`                  // Username
	Password                  *string            `bson:"password,omitempty"`                  // Password (hashed)
	PhoneNumber               *string            `bson:"phoneNumber,omitempty"`               // Phone number (encrypted)
	Address                   *string            `bson:"address,omitempty"`                   // Address
	Line1                     *string            `bson:"line1,omitempty"`                     // Line 1 of address
	Line2                     *string            `bson:"line2,omitempty"`                     // Line 2 of address
	Zipcode                   *int               `bson:"zipcode,omitempty"`                   // Zip code (postal code)
	State                     *string            `bson:"state,omitempty"`                     // State
	City                      *string            `bson:"city,omitempty"`                      // City
	Country                   *string            `bson:"country,omitempty"`                   // Country
	WebsiteLink               *string            `bson:"websiteLink,omitempty"`               // Website link
	NoOfEmployees             *string            `bson:"noOfEmployees,omitempty"`             // Number of employees
	CompanyRegistrationNumber *int               `bson:"companyRegistrationNumber,omitempty"` // Company registration number
	UserWallet                *string            `bson:"userWallet,omitempty"`                // User wallet address
	DiscountPrice             *string            `bson:"discountPrice,omitempty"`             // Discount price
	LogoURL                   *string            `bson:"logoURL,omitempty"`                   // Logo URL
	Role                      *string            `bson:"role,omitempty"`                      // Role in the system (default: "company")
	AvailableChecks           []string           `bson:"availableChecks,omitempty"`           // List of available checks
	VerificationCode          *int               `bson:"verificationCode,omitempty"`          // Verification code
	VerificationCodeTimestamp *time.Time         `bson:"verificationCodeTimestamp,omitempty"` // Timestamp for verification code
	PrePayment                *bool              `bson:"prePayment,omitempty"`                // Pre-payment status
	OtpBlockEndTime           *time.Time         `bson:"otpBlockEndTime,omitempty"`           // OTP block end time
	ApiKeyRequest             ApiKeyRequest      `bson:"apiKeyRequest,omitempty"`             // API key request details
	CreatedAt                 *time.Time         `bson:"createdAt,omitempty"`                 // Timestamp for creation
	UpdatedAt                 *time.Time         `bson:"updatedAt,omitempty"`                 // Timestamp for last update
}

// ApiKeyRequest represents the nested structure for API key request details.
type ApiKeyRequest struct {
	ApiKey        *string `bson:"apiKey,omitempty"`        // API key
	RequestStatus *string `bson:"requestStatus,omitempty"` // Request status (null, "Pending", "Generated")
}
