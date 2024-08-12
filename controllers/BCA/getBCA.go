package BCA

import (
	"CCT-GOLANG-BACKEND/db"
	"CCT-GOLANG-BACKEND/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetBCA(c *gin.Context) {
	collection := db.GetCollection("BCA")

	// Fetch all BCA records
	cursor, err := collection.Find(c, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching BCA records"})
		return
	}
	defer cursor.Close(c)

	var bcas []models.BCA
	if err = cursor.All(c, &bcas); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding BCA records"})
		return
	}

	c.JSON(http.StatusOK, bcas)
}
