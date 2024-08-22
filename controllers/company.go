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

	collection := db.GetCollection("company")
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
	newCompany := models.Company{
		ID:                        primitive.NewObjectID(),
		Name:                      input.Name,
		Email:                     input.Email,
		UserName:                  input.UserName,
		Password:                  &hashedPassword, // Store hashed password
		PhoneNumber:               input.PhoneNumber,
		Address:                   input.Address,
		Line2:                     input.Line2,
		Line1:                     input.Line1,
		Zipcode:                   input.Zipcode,
		State:                     input.State,
		City:                      input.City,
		Country:                   input.Country,
		WebsiteLink:               input.WebsiteLink,
		NoOfEmployees:             input.NoOfEmployees,
		CompanyRegistrationNumber: input.CompanyRegistrationNumber,
		UserWallet:                input.UserWallet,
		DiscountPrice:             input.DiscountPrice,
		LogoURL:                   input.LogoURL,
		Role:                      input.Role,
		AvailableChecks:           input.AvailableChecks,
		VerificationCode:          input.VerificationCode,
		VerificationCodeTimestamp: input.VerificationCodeTimestamp,
		PrePayment:                input.PrePayment,
		BCAID:                     userID, // Store BCA ID from token
		CreatedAt:                 time.Now(),
		UpdatedAt:                 time.Now(),
	}

	_, err = collection.InsertOne(ctx, newCompany)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating company"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Company registered successfully"})
}
