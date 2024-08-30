package db

import "go.mongodb.org/mongo-driver/bson"

// CleanPaginationID removes the "id" field from each document in the result set
func CleanPaginationID(result []bson.M) []bson.M {
	for _, element := range result {
		delete(element, "id")
	}
	return result
}
