package controllers

import (
	"CCT-GOLANG-BACKEND/db"
	"CCT-GOLANG-BACKEND/middleware"
	"CCT-GOLANG-BACKEND/models"
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// RegisterCompany handles company registration
func RegisterCompany(c *gin.Context) {
	var input models.Company
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate JWT token
	tokenStr := c.GetHeader("Authorization")
	if tokenStr == "" || !strings.HasPrefix(tokenStr, "Bearer ") {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing or malformed"})
		return
	}

	tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")
	claims, err := middleware.ValidateToken(tokenStr)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	role, ok := claims["role"].(string)
	if !ok || role != "BCA" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		return
	}

	// Extract BCA ID from token claims
	bcaIDStr, ok := claims["userID"].(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "BCA ID is missing in token"})
		return
	}

	userID, err := primitive.ObjectIDFromHex(bcaIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid BCA ID"})
		return
	}

	collection := db.GetCollection("companies")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Check if company email already exists
	var existingCompany models.Company
	err = collection.FindOne(ctx, bson.M{"email": input.Email}).Decode(&existingCompany)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already registered"})
		return
	}

	// Check if company username already exists
	err = collection.FindOne(ctx, bson.M{"userName": input.UserName}).Decode(&existingCompany)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Username already taken"})
		return
	}

	// Hash the password using bcrypt
	hashedPassword, err := middleware.HashPassword(*input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error encrypting password"})
		return
	}

	// Create new company with hashed password and BCA ID
	now := time.Now()
	newCompany := models.Company{
		ID:                        primitive.NewObjectID(),         // Generate a new ObjectID for the company
		BCAId:                     userID,                          // Store BCA ID from token
		Name:                      input.Name,                      // Company name
		Email:                     input.Email,                     // Email address
		UserName:                  input.UserName,                  // Username
		Password:                  &hashedPassword,                 // Store hashed password
		PhoneNumber:               input.PhoneNumber,               // Phone number
		Address:                   input.Address,                   // Address
		Line1:                     input.Line1,                     // Line 1 of address
		Line2:                     input.Line2,                     // Line 2 of address
		Zipcode:                   input.Zipcode,                   // Zip code (postal code)
		State:                     input.State,                     // State
		City:                      input.City,                      // City
		Country:                   input.Country,                   // Country
		WebsiteLink:               input.WebsiteLink,               // Website link
		NoOfEmployees:             input.NoOfEmployees,             // Number of employees
		CompanyRegistrationNumber: input.CompanyRegistrationNumber, // Company registration number
		UserWallet:                input.UserWallet,                // User wallet address
		DiscountPrice:             input.DiscountPrice,             // Discount price
		LogoURL:                   input.LogoURL,                   // Logo URL
		Role:                      input.Role,                      // Role in the system (default: "company")
		AvailableChecks:           input.AvailableChecks,           // List of available checks
		VerificationCode:          input.VerificationCode,          // Verification code
		VerificationCodeTimestamp: input.VerificationCodeTimestamp, // Timestamp for verification code
		PrePayment:                input.PrePayment,                // Pre-payment status
		OtpBlockEndTime:           input.OtpBlockEndTime,           // OTP block end time
		ApiKeyRequest: models.ApiKeyRequest{ // Initialize nested struct for API key request
			ApiKey:        input.ApiKeyRequest.ApiKey,
			RequestStatus: input.ApiKeyRequest.RequestStatus,
		},
		CreatedAt: &now, // Timestamp for creation
		UpdatedAt: &now, // Timestamp for last update
	}

	_, err = collection.InsertOne(ctx, newCompany)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating company"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Company registered successfully"})
}
