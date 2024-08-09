package controllers

import (
	"context"
	"net/http"
	"strings"
	"time"

	"CCT-GOLANG-BACKEND/db"
	"CCT-GOLANG-BACKEND/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
	_, err := ValidateToken(tokenStr)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	collection := db.GetCollection("company")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	company := models.Company{
		ID:                        primitive.NewObjectID(),
		Name:                      input.Name,
		Email:                     input.Email,
		UserName:                  input.UserName,
		Password:                  input.Password, // You might want to hash this as well
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

	_, err = collection.InsertOne(ctx, company)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating company"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Company registered successfully"})
}
