package db

import (
	"context"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	clientInstance      *mongo.Client
	clientInstanceError error
	mongoOnce           sync.Once
)

const (
	CONNECTIONSTRING = "mongodb://localhost:27017"
	DB               = "CCTdb"
)

func init() {
	clientOptions := options.Client().ApplyURI(CONNECTIONSTRING)

	mongoOnce.Do(func() {
		clientInstance, clientInstanceError = mongo.Connect(context.TODO(), clientOptions)
		if clientInstanceError != nil {
			log.Fatal(clientInstanceError)
		}

		err := clientInstance.Ping(context.TODO(), nil)
		if err != nil {
			log.Fatal(err)
		}
	})
}

func GetMongoClient() *mongo.Client {
	return clientInstance
}

func GetCollection(collectionName string) *mongo.Collection {
	client := GetMongoClient()
	return client.Database(DB).Collection(collectionName)
}
