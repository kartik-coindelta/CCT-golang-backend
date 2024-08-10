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
	_, err := middleware.ValidateToken(tokenStr)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
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

	// Create new company with hashed password
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
