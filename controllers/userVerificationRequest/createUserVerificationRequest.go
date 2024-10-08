package controllers

import (
	"CCT-GOLANG-BACKEND/db"
	"CCT-GOLANG-BACKEND/middleware"
	"CCT-GOLANG-BACKEND/models"
	"context"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateUserVerificationRequest(c *gin.Context) {
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

	companyID, err := primitive.ObjectIDFromHex(companyIDStr)
	if err != nil {
		log.Println("Error converting company ID to ObjectID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Company ID"})
		return
	}

	var request models.UserVerificationRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Fetch specific company fields from the company collection using companyID
	companyCollection := db.GetCollection("companies")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var companyInfo struct {
		CompanyName            string             `bson:"name" json:"name"`
		CompanyEmail           string             `bson:"email" json:"email"`
		CompanyAddress         string             `bson:"address" json:"address"`
		CompanyWebsiteLink     string             `bson:"websiteLink" json:"websiteLink"`
		CompanyLogoUrl         string             `bson:"logoURL" json:"logoURL"`
		CompanyAvailableChecks []string           `bson:"availableChecks" json:"availableChecks"`
		BCAId                  primitive.ObjectID `bson:"bcaId" json:"bcaId"`
	}
	projection := bson.M{"name": 1, "email": 1, "address": 1, "websiteLink": 1, "logoURL": 1, "availableChecks": 1, "bcaId": 1}

	err = companyCollection.FindOne(ctx, bson.M{"_id": companyID}, options.FindOne().SetProjection(projection)).Decode(&companyInfo)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Company not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching company information"})
		return
	}

	*request.BCAId = companyInfo.BCAId // Include bcaId directly in the request

	// Populate request.CompanyInfo with fetched company info
	request.CompanyInfo = map[string]interface{}{
		"name":            companyInfo.CompanyName,
		"email":           companyInfo.CompanyEmail,
		"address":         companyInfo.CompanyAddress,
		"websiteLink":     companyInfo.CompanyWebsiteLink,
		"logoURL":         nil,
		"availableChecks": companyInfo.CompanyAvailableChecks,
	}

	bcaCollection := db.GetCollection("bcas")

	var bcaInfo struct {
		BCAName    string `bson:"name" json:"name"`
		BCAEmail   string `bson:"email" json:"email"`
		BCACountry string `bson:"country" json:"country"`
		BCACity    string `bson:"city" json:"city"`
		BCAAddress string `bson:"address" json:"address"`
	}
	projection = bson.M{"name": 1, "email": 1, "country": 1, "city": 1, "address": 1}

	err = bcaCollection.FindOne(ctx, bson.M{"_id": companyInfo.BCAId}, options.FindOne().SetProjection(projection)).Decode(&bcaInfo)
	if err != nil {
		log.Println("MongoDB Query Error:", err) // Log the exact error
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusBadRequest, gin.H{"error": "BCA not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching BCA information"})
		return
	}

	// Populate request.CompanyInfo with fetched companyName and companyAddress
	request.BCAInfo = map[string]interface{}{
		"_id":     companyInfo.BCAId,
		"name":    bcaInfo.BCAName,
		"email":   bcaInfo.BCAEmail,
		"country": bcaInfo.BCACountry,
		"city":    bcaInfo.BCACity,
		"address": bcaInfo.BCAAddress,
	}

	// Assign a new ObjectID to the UserVerificationRequest
	request.ID = primitive.NewObjectID()
	request.CompanyId = companyID

	// Insert the UserVerificationRequest into the database
	verificationRequestCollection := db.GetCollection("userverificationrequest")
	_, err = verificationRequestCollection.InsertOne(ctx, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user verification request"})
		return
	}

	// Generate a new JWT token with no expiration
	tokenString, err := middleware.GenerateTokenWithoutExpiry(request.ID.Hex())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User verification request created successfully",
		"token":   tokenString,
	})
}
