package company

import (
	"CCT-GOLANG-BACKEND/db"
	"CCT-GOLANG-BACKEND/models"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetCompany(c *gin.Context) {
	collection := db.GetCollection("company")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var companies []models.Company

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching companies"})
		return
	}

	if err := cursor.All(ctx, &companies); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding companies"})
		return
	}

	c.JSON(http.StatusOK, companies)
}
