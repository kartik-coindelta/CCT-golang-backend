package controllers

import (
	"CCT-GOLANG-BACKEND/db"
	"CCT-GOLANG-BACKEND/middleware"
	"CCT-GOLANG-BACKEND/models"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SignUp(c *gin.Context) {
	var input models.BCA
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collection := db.GetCollection("bcas")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var existingUser models.BCA

	// Check if email already exists
	err := collection.FindOne(ctx, bson.M{"email": input.Email}).Decode(&existingUser)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already registered"})
		return
	}

	// Check if username already exists
	err = collection.FindOne(ctx, bson.M{"userName": input.UserName}).Decode(&existingUser)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Username already taken"})
		return
	}

	// Encrypt the password
	hashedPassword, err := middleware.HashPassword(*input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error encrypting password"})
		return
	}

	// Create new user with role set to "bca"
	now := time.Now()

	newUser := models.BCA{
		ID:                        primitive.NewObjectID(),         // New ObjectID for the BCA
		BCAId:                     input.BCAId,                     // BCA ID reference if needed
		Name:                      input.Name,                      // Name of the BCA
		FirstName:                 input.FirstName,                 // First Name of the BCA user
		LastName:                  input.LastName,                  // Last Name of the BCA user
		UserName:                  input.UserName,                  // Username
		Email:                     input.Email,                     // Email
		Password:                  &hashedPassword,                 // Hashed password
		PhoneNumber:               input.PhoneNumber,               // Phone number
		Line1:                     input.Line1,                     // Line 1 of the address
		Line2:                     input.Line2,                     // Line 2 of the address
		Zipcode:                   input.Zipcode,                   // Zipcode
		CompanyRegistrationNumber: input.CompanyRegistrationNumber, // Company registration number
		WebsiteLink:               input.WebsiteLink,               // Website link
		NoOfEmployees:             input.NoOfEmployees,             // Number of employees
		UserWallet:                input.UserWallet,                // User wallet
		SupportingDocuments:       input.SupportingDocuments,       // Supporting documents
		LogoURL:                   input.LogoURL,                   // Logo URL
		Status:                    input.Status,                    // Status
		Role:                      input.Role,                      // Role (BCA, BCAStaff, etc.)
		VendorName:                input.VendorName,                // Vendor name
		ManagerName:               input.ManagerName,               // Manager name
		Address:                   input.Address,                   // Address
		City:                      input.City,                      // City
		State:                     input.State,                     // State
		GST:                       input.GST,                       // GST
		Country:                   input.Country,                   // Country
		AdditionalRemark:          input.AdditionalRemark,          // Additional remark
		HasStaffAccess:            input.HasStaffAccess,            // Has staff access
		VerificationCode:          input.VerificationCode,          // Verification code
		VerificationCodeTimestamp: input.VerificationCodeTimestamp, // Verification code timestamp
		OtpBlockEndTime:           input.OtpBlockEndTime,           // OTP block end time
		SES_URL:                   input.SES_URL,                   // SES URL
		EmailBannerImage:          input.EmailBannerImage,          // Email banner image
		CompanyTitle:              input.CompanyTitle,              // Company title
		OtpAttempts:               input.OtpAttempts,               // OTP attempts
		CreatedAt:                 now,                             // Creation time
		UpdatedAt:                 now,                             // Last updated time
	}

	_, err = collection.InsertOne(ctx, newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}
