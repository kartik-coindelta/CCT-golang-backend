package db

import (
    "context"
    "go.mongodb.org/mongo-driver/mongo"
)

// CreateItem inserts a new item into the given MongoDB collection without using transactions
func CreateItem(collection *mongo.Collection, item interface{}, ctx context.Context) (*mongo.InsertOneResult, error) {
    result, err := collection.InsertOne(ctx, item)
    if err != nil {
        return nil, err
    }
    return result, nil
}
