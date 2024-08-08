package controllers

import (
	"CCT-GOLANG-BACKEND/db"
	"CCT-GOLANG-BACKEND/models"
	"context"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func hashPasswordWithSecret(password, secretKey string) (string, error) {
	combined := password + secretKey
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(combined), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

func Signup(c *gin.Context) {
	var input models.BCA
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	secretKey := os.Getenv("SECRET_KEY")

	collection := db.GetCollection("BCA")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var existingUser models.BCA

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

	hashedPassword, err := hashPasswordWithSecret(*input.Password, secretKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error encrypting password"})
		return
	}

	// Register new user with role set to "BCA"
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
