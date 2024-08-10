package controllers

import (
	"CCT-GOLANG-BACKEND/db"
	"CCT-GOLANG-BACKEND/middleware"
	"CCT-GOLANG-BACKEND/models"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Login(c *gin.Context) {
	var input struct {
		Email    *string `json:"email"`
		UserName *string `json:"userName"`
		Password *string `json:"password"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collectionNames := []string{"BCA", "company", "user"}

	var user models.User
	var company models.Company
	var bca models.BCA

	for _, collectionName := range collectionNames {
		collection := db.GetCollection(collectionName)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		filter := bson.M{}
		if input.Email != nil && *input.Email != "" {
			filter["email"] = input.Email
		} else if input.UserName != nil && *input.UserName != "" {
			filter["userName"] = input.UserName
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email or Username required"})
			return
		}

		var result interface{}
		switch collectionName {
		case "BCA":
			result = &bca
			log.Println("U....JUST...LOGIN....TO...BCA.....")
		case "company":
			result = &company
			log.Println("U....JUST...LOGIN....TO...COMPANY.....")
		case "user":
			result = &user
			log.Println("U....JUST...LOGIN....TO...USER.....")
		}

		err := collection.FindOne(ctx, filter).Decode(result)
		if err == nil {
			// Compare password
			var hashedPassword string
			var role string
			var id primitive.ObjectID

			switch collectionName {
			case "BCA":
				hashedPassword = *bca.Password
				role = *bca.Role
				id = bca.ID
			case "company":
				hashedPassword = *company.Password
				role = *company.Role
				id = company.ID
				log.Println(role)
			case "user":
				hashedPassword = *user.Password
				role = *user.Role
				id = user.ID
				log.Println(role)
			}

			err = middleware.ComparePassword(hashedPassword, *input.Password)
			if err != nil {
				continue
			}

			// Generate JWT Token
			token, err := middleware.GenerateToken(id.Hex(), role)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
				return
			}

			c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
			return
		}
	}

	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email/username or password"})
}
