package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"CCT-GOLANG-BACKEND/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

var client *mongo.Client

func InitMongo() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(ctx, clientOptions)
}

func SignUp(c *gin.Context) {
	var req map[string]string
	if err := c.ShouldBindJSON(&req); err != nil {
		handleError(c, err)
		return
	}

	if email, ok := req["email"]; ok && email != "" {
		if userData, err := findUserByEmail(email); err == nil {
			if len(userData) != 0 {
				c.JSON(http.StatusOK, userData[0])
				return
			}
		} else {
			handleError(c, err)
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "This email is not registered with us."})
	} else if userName, ok := req["userName"]; ok && userName != "" {
		if userData, err := findUserByUserName(userName); err == nil {
			if len(userData) != 0 {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists."})
				return
			}
		} else {
			handleError(c, err)
			return
		}

		encryptedPassword, err := encryptPassword(req["password"])
		if err != nil {
			handleError(c, err)
			return
		}

		newUser := models.User{
			UserName: req["userName"],
			Email:    req["email"],
			Password: encryptedPassword,
		}

		collection := client.Database("yourdb").Collection("users") // Assuming you are inserting into the "users" collection
		_, err = collection.InsertOne(context.TODO(), newUser)
		if err != nil {
			handleError(c, err)
			return
		}
		c.JSON(http.StatusOK, newUser)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Email or userName is required."})
	}
}

func findUserByEmail(email string) ([]models.User, error) {
	var userData []models.User
	collections := []string{"users", "BCAs", "companies"}
	for _, collectionName := range collections {
		collection := client.Database("yourdb").Collection(collectionName)
		filter := bson.M{"email": email}
		cursor, err := collection.Find(context.TODO(), filter)
		if err != nil {
			return nil, err
		}
		if err = cursor.All(context.TODO(), &userData); err != nil {
			return nil, err
		}
		if len(userData) != 0 {
			return userData, nil
		}
	}
	return nil, nil
}

func findUserByUserName(userName string) ([]models.User, error) {
	var userData []models.User
	collections := []string{"users", "BCAs", "companies"}
	for _, collectionName := range collections {
		collection := client.Database("yourdb").Collection(collectionName)
		filter := bson.M{"userName": userName}
		cursor, err := collection.Find(context.TODO(), filter)
		if err != nil {
			return nil, err
		}
		if err = cursor.All(context.TODO(), &userData); err != nil {
			return nil, err
		}
		if len(userData) != 0 {
			return userData, nil
		}
	}
	return nil, nil
}

func encryptPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func handleError(c *gin.Context, err error) {
	log.Println(err)
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}
