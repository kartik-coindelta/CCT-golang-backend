package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// CreateItem inserts a new document into the MongoDB collection.
func CreateItem(collection *mongo.Collection, data interface{}, session mongo.SessionContext) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if session != nil {
		return collection.InsertOne(session, data) // Use session directly
	}
	return collection.InsertOne(ctx, data) // Use default context
}
