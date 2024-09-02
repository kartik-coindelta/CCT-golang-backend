package db

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// UpdateItem updates an item in the database by ID without using a session.
func UpdateItem(id string, updateData interface{}, collection *mongo.Collection) (*mongo.SingleResult, error) {
	// Convert the ID string to a MongoDB ObjectID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid ID format")
	}

	// Set options for the update operation
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After).SetUpsert(false)

	// Filter and update data
	filter := bson.M{"_id": objID}
	update := bson.M{"$set": updateData}

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Update the document in the collection without using a session
	result := collection.FindOneAndUpdate(ctx, filter, update, opts)
	if result.Err() != nil {
		return nil, result.Err()
	}

	return result, nil
}
