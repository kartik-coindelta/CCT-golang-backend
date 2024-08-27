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

func RegisterUser(c *gin.Context) {
	// Extract JWT token from Authorization header
	tokenStr := c.GetHeader("Authorization")
	if tokenStr == "" || !strings.HasPrefix(tokenStr, "Bearer ") {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing or malformed"})
		return
	}
	tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")

	// Validate the token and get claims
	claims, err := middleware.ValidateToken(tokenStr)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	// Check if the role is "company"
	role, ok := claims["role"].(string)
	if !ok || role != "company" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		return
	}

	// Extract Company ID from token claims
	companyIDStr, ok := claims["userID"].(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Company ID is missing in token"})
		return
	}

	userID, err := primitive.ObjectIDFromHex(companyIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Company ID"})
		return
	}

	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collection := db.GetCollection("users") // Assuming User collection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var existingUser models.User

	// Check if email already exists
	err = collection.FindOne(ctx, bson.M{"email": input.Email}).Decode(&existingUser)
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

	// Create new user
	companyID := &userID
	newUser := models.User{
		ID:                        primitive.NewObjectID(),
		FirstName:                 input.FirstName,
		LastName:                  input.LastName,
		CaseNumber:                input.CaseNumber,
		ClientName:                input.ClientName,
		Email:                     input.Email,
		UserName:                  input.UserName,
		Password:                  &hashedPassword,
		FatherName:                input.FatherName,
		MotherName:                input.MotherName,
		DateOfBirth:               input.DateOfBirth,
		Gender:                    input.Gender,
		PhoneNumber:               input.PhoneNumber,
		ImgURL:                    input.ImgURL,
		UserWallet:                input.UserWallet,
		Role:                      input.Role,
		CompanyID:                 companyID,
		CaseID:                    input.CaseID,
		VerificationCode:          input.VerificationCode,
		Address:                   input.Address,
		Address1:                  input.Address1,
		Address2:                  input.Address2,
		City:                      input.City,
		State:                     input.State,
		Pincode:                   input.Pincode,
		Country:                   input.Country,
		VerificationCodeTimestamp: input.VerificationCodeTimestamp,
		OTPBlockEndTime:           input.OTPBlockEndTime,
		CreatedAt:                 time.Now(),
		UpdatedAt:                 time.Now(),
	}

	_, err = collection.InsertOne(ctx, newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}
