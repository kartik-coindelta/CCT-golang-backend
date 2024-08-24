package controllers

import (
	"context"
	"net/http"
	"strings"
	"time"

	"CCT-GOLANG-BACKEND/db"
	"CCT-GOLANG-BACKEND/middleware"
	"CCT-GOLANG-BACKEND/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Register(c *gin.Context) {
	var input map[string]interface{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	role, ok := input["role"].(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Role is required"})
		return
	}

	switch role {
	case "BCA":
		var bcaInput models.BCA
		if err := c.ShouldBindJSON(&bcaInput); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		hashedPassword, err := middleware.HashPassword(*bcaInput.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error encrypting password"})
			return
		}

		now := time.Now()
		newBCA := models.BCA{
			ID:                        primitive.NewObjectID(),            // New ObjectID for the BCA
			BCAId:                     bcaInput.BCAId,                     // BCA ID reference if needed
			Name:                      bcaInput.Name,                      // Name of the BCA
			FirstName:                 bcaInput.FirstName,                 // First Name of the BCA user
			LastName:                  bcaInput.LastName,                  // Last Name of the BCA user
			UserName:                  bcaInput.UserName,                  // Username
			Email:                     bcaInput.Email,                     // Email
			Password:                  &hashedPassword,                    // Hashed password
			PhoneNumber:               bcaInput.PhoneNumber,               // Phone number
			Line1:                     bcaInput.Line1,                     // Line 1 of the address
			Line2:                     bcaInput.Line2,                     // Line 2 of the address
			Zipcode:                   bcaInput.Zipcode,                   // Zipcode
			CompanyRegistrationNumber: bcaInput.CompanyRegistrationNumber, // Company registration number
			WebsiteLink:               bcaInput.WebsiteLink,               // Website link
			NoOfEmployees:             bcaInput.NoOfEmployees,             // Number of employees
			UserWallet:                bcaInput.UserWallet,                // User wallet
			SupportingDocuments:       bcaInput.SupportingDocuments,       // Supporting documents
			LogoURL:                   bcaInput.LogoURL,                   // Logo URL
			Status:                    bcaInput.Status,                    // Status
			Role:                      bcaInput.Role,                      // Role (BCA, BCAStaff, etc.)
			VendorName:                bcaInput.VendorName,                // Vendor name
			ManagerName:               bcaInput.ManagerName,               // Manager name
			Address:                   bcaInput.Address,                   // Address
			City:                      bcaInput.City,                      // City
			State:                     bcaInput.State,                     // State
			GST:                       bcaInput.GST,                       // GST
			Country:                   bcaInput.Country,                   // Country
			AdditionalRemark:          bcaInput.AdditionalRemark,          // Additional remark
			HasStaffAccess:            bcaInput.HasStaffAccess,            // Has staff access
			VerificationCode:          bcaInput.VerificationCode,          // Verification code
			VerificationCodeTimestamp: bcaInput.VerificationCodeTimestamp, // Verification code timestamp
			OtpBlockEndTime:           bcaInput.OtpBlockEndTime,           // OTP block end time
			SES_URL:                   bcaInput.SES_URL,                   // SES URL
			EmailBannerImage:          bcaInput.EmailBannerImage,          // Email banner image
			CompanyTitle:              bcaInput.CompanyTitle,              // Company title
			OtpAttempts:               bcaInput.OtpAttempts,               // OTP attempts
			CreatedAt:                 now,                                // Creation time
			UpdatedAt:                 now,                                // Last updated time
		}

		collection := db.GetCollection("bcas")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		_, err = collection.InsertOne(ctx, newBCA)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating BCA"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "BCA registered successfully"})

	case "Company":
		var companyInput models.Company
		if err := c.ShouldBindJSON(&companyInput); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

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

		bcaIDStr, ok := claims["userID"].(string)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "BCA ID is missing in token"})
			return
		}

		bcaID, err := primitive.ObjectIDFromHex(bcaIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid BCA ID"})
			return
		}

		collection := db.GetCollection("companies")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var existingCompany models.Company
		err = collection.FindOne(ctx, bson.M{"email": companyInput.Email}).Decode(&existingCompany)
		if err == nil {
			c.JSON(http.StatusConflict, gin.H{"error": "Email already registered"})
			return
		}

		err = collection.FindOne(ctx, bson.M{"userName": companyInput.UserName}).Decode(&existingCompany)
		if err == nil {
			c.JSON(http.StatusConflict, gin.H{"error": "Username already taken"})
			return
		}

		hashedPassword, err := middleware.HashPassword(*companyInput.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error encrypting password"})
			return
		}

		now := time.Now()
		newCompany := models.Company{
			ID:                        primitive.NewObjectID(),
			BCAId:                     bcaID,
			Name:                      companyInput.Name,
			Email:                     companyInput.Email,
			UserName:                  companyInput.UserName,
			Password:                  &hashedPassword,
			PhoneNumber:               companyInput.PhoneNumber,
			Address:                   companyInput.Address,
			Line1:                     companyInput.Line1,
			Line2:                     companyInput.Line2,
			Zipcode:                   companyInput.Zipcode,
			State:                     companyInput.State,
			City:                      companyInput.City,
			Country:                   companyInput.Country,
			WebsiteLink:               companyInput.WebsiteLink,
			NoOfEmployees:             companyInput.NoOfEmployees,
			CompanyRegistrationNumber: companyInput.CompanyRegistrationNumber,
			UserWallet:                companyInput.UserWallet,
			DiscountPrice:             companyInput.DiscountPrice,
			LogoURL:                   companyInput.LogoURL,
			Role:                      companyInput.Role,
			AvailableChecks:           companyInput.AvailableChecks,
			VerificationCode:          companyInput.VerificationCode,
			VerificationCodeTimestamp: companyInput.VerificationCodeTimestamp,
			PrePayment:                companyInput.PrePayment,
			OtpBlockEndTime:           companyInput.OtpBlockEndTime,
			ApiKeyRequest: models.ApiKeyRequest{
				ApiKey:        companyInput.ApiKeyRequest.ApiKey,
				RequestStatus: companyInput.ApiKeyRequest.RequestStatus,
			},
			CreatedAt: &now,
			UpdatedAt: &now,
		}

		_, err = collection.InsertOne(ctx, newCompany)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating company"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Company registered successfully"})

	case "User":
		var userInput models.User
		if err := c.ShouldBindJSON(&userInput); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

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
		if !ok || role != "company" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
			return
		}

		companyIDStr, ok := claims["userID"].(string)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Company ID is missing in token"})
			return
		}

		companyID, err := primitive.ObjectIDFromHex(companyIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Company ID"})
			return
		}
		companyIDPtr := &companyID

		collection := db.GetCollection("users")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var existingUser models.User
		err = collection.FindOne(ctx, bson.M{"email": userInput.Email}).Decode(&existingUser)
		if err == nil {
			c.JSON(http.StatusConflict, gin.H{"error": "Email already registered"})
			return
		}

		err = collection.FindOne(ctx, bson.M{"userName": userInput.UserName}).Decode(&existingUser)
		if err == nil {
			c.JSON(http.StatusConflict, gin.H{"error": "Username already taken"})
			return
		}

		hashedPassword, err := middleware.HashPassword(*userInput.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error encrypting password"})
			return
		}

		newUser := models.User{
			ID:                        primitive.NewObjectID(),
			FirstName:                 userInput.FirstName,
			LastName:                  userInput.LastName,
			CaseNumber:                userInput.CaseNumber,
			ClientName:                userInput.ClientName,
			Email:                     userInput.Email,
			UserName:                  userInput.UserName,
			Password:                  &hashedPassword,
			FatherName:                userInput.FatherName,
			MotherName:                userInput.MotherName,
			DateOfBirth:               userInput.DateOfBirth,
			Gender:                    userInput.Gender,
			PhoneNumber:               userInput.PhoneNumber,
			ImgURL:                    userInput.ImgURL,
			UserWallet:                userInput.UserWallet,
			Role:                      userInput.Role,
			CompanyID:                 companyIDPtr,
			CaseID:                    userInput.CaseID,
			VerificationCode:          userInput.VerificationCode,
			Address:                   userInput.Address,
			Address1:                  userInput.Address1,
			Address2:                  userInput.Address2,
			City:                      userInput.City,
			State:                     userInput.State,
			Pincode:                   userInput.Pincode,
			Country:                   userInput.Country,
			VerificationCodeTimestamp: userInput.VerificationCodeTimestamp,
			OTPBlockEndTime:           userInput.OTPBlockEndTime,
			CreatedAt:                 time.Now(),
			UpdatedAt:                 time.Now(),
		}

		_, err = collection.InsertOne(ctx, newUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})

	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role"})
	}
}
