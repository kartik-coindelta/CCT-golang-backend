package controllers

import (
	"context"
	"net/http"
	"time"

	"CCT-GOLANG-BACKEND/db"
	"CCT-GOLANG-BACKEND/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// Concatenate password with secret key and hash it
func hashPassword(password string) (string, error) {
	// Concatenate password with secret key
	concatenatedPassword := password + string(bcryptSecretKey)

	bytes, err := bcrypt.GenerateFromPassword([]byte(concatenatedPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// Signup handles user registration
func Signup(c *gin.Context) {
	var input models.BCA
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collection := db.GetCollection("BCA")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var existingUser models.BCA

	// Check if email or username already exists
	err := collection.FindOne(ctx, bson.M{"email": input.Email}).Decode(&existingUser)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already registered"})
		return
	}

	err = collection.FindOne(ctx, bson.M{"userName": input.UserName}).Decode(&existingUser)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Username already taken"})
		return
	}

	// Encrypt the password with secret key concatenation
	hashedPassword, err := hashPassword(*input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error encrypting password"})
		return
	}

	newUser := models.BCA{
		ID:                        primitive.NewObjectID(),
		UserName:                  input.UserName,
		Email:                     input.Email,
		Password:                  &hashedPassword,
		FirstName:                 input.FirstName,
		LastName:                  input.LastName,
		PhoneNumber:               input.PhoneNumber,
		Line1:                     input.Line1,
		Line2:                     input.Line2,
		Zipcode:                   input.Zipcode,
		CompanyRegistrationNumber: input.CompanyRegistrationNumber,
		WebsiteLink:               input.WebsiteLink,
		NoOfEmployees:             input.NoOfEmployees,
		UserWallet:                input.UserWallet,
		SupportingDocuments:       input.SupportingDocuments,
		LogoURL:                   input.LogoURL,
		Status:                    input.Status,
		Role:                      input.Role,
		VendorName:                input.VendorName,
		ManagerName:               input.ManagerName,
		Address:                   input.Address,
		City:                      input.City,
		State:                     input.State,
		GST:                       input.GST,
		Country:                   input.Country,
		AdditionalRemark:          input.AdditionalRemark,
		HasStaffAccess:            input.HasStaffAccess,
		VerificationCode:          input.VerificationCode,
		VerificationCodeTimestamp: input.VerificationCodeTimestamp,
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
