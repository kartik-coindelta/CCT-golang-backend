package db

import (
	"net/http"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ListInitOptions constructs initial options for a MongoDB query based on request parameters.
func ListInitOptions(req *http.Request) (*options.FindOptions, error) {
	order := -1 // default order
	if req.URL.Query().Get("order") != "" {
		parsedOrder, err := strconv.Atoi(req.URL.Query().Get("order"))
		if err != nil {
			return nil, err
		}
		order = parsedOrder
	}

	sort := req.URL.Query().Get("sort")
	if sort == "" {
		sort = "createdAt" // default sort field
	}
	sortBy := BuildSort(sort, order)

	page := 1 // default page
	if req.URL.Query().Get("page") != "" {
		parsedPage, err := strconv.Atoi(req.URL.Query().Get("page"))
		if err != nil {
			return nil, err
		}
		page = parsedPage
	}

	limit := 5 // default limit
	if req.URL.Query().Get("limit") != "" {
		parsedLimit, err := strconv.Atoi(req.URL.Query().Get("limit"))
		if err != nil {
			return nil, err
		}
		limit = parsedLimit
	}

	populateArray := make([]bson.M, 0)
	populateFields := req.URL.Query().Get("populate")
	if populateFields != "" {
		fields := strings.Split(populateFields, ",")
		selectPopulate := req.URL.Query().Get("selectPopulate")
		selectPopulateFields := strings.Split(selectPopulate, ",")

		for _, field := range fields {
			populate := bson.M{
				"path":   field,
				"select": selectPopulateFields,
			}
			populateArray = append(populateArray, populate)
		}
	}

	// Create FindOptions with sorting and pagination
	findOptions := options.Find()
	findOptions.SetSort(sortBy)
	findOptions.SetSkip(int64((page - 1) * limit))
	findOptions.SetLimit(int64(limit))
	// Add any additional options such as populate here if required

	return findOptions, nil
}

// BuildSort creates a sort object for MongoDB queries
// sortField: the field to sort by
// order: the sort order (1 for ascending, -1 for descending)
