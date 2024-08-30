package db

import (
	"go.mongodb.org/mongo-driver/bson"
)

// BuildSort creates a sort object for MongoDB queries
// sortField: the field to sort by
// order: the sort order (1 for ascending, -1 for descending)
func BuildSort(sortField string, order int) bson.D {
	sortBy := bson.D{}
	if sortField != "" {
		sortBy = append(sortBy, bson.E{Key: sortField, Value: order})
	}
	return sortBy
}
