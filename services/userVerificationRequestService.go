package services

import (
	"context"
	"time"

	"CCT-GOLANG-BACKEND/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserVerificationRequestService struct {
	collection *mongo.Collection
}

// NewUserVerificationRequestService creates a new service instance
func NewUserVerificationRequestService(db *mongo.Database) *UserVerificationRequestService {
	return &UserVerificationRequestService{
		collection: db.Collection("userVerificationRequests"),
	}
}

// Create creates a new UserVerificationRequest
func (s *UserVerificationRequestService) Create(ctx context.Context, request *models.UserVerificationRequest) (*mongo.InsertOneResult, error) {
	request.CreatedAt = time.Now()
	request.UpdatedAt = time.Now()

	// Handle the pre-save hook
	err := request.BeforeSave(ctx, s.collection.Database())
	if err != nil {
		return nil, err
	}

	return s.collection.InsertOne(ctx, request)
}

// GetByID retrieves a UserVerificationRequest by ID
func (s *UserVerificationRequestService) GetByID(ctx context.Context, id primitive.ObjectID) (*models.UserVerificationRequest, error) {
	var request models.UserVerificationRequest
	err := s.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&request)
	if err != nil {
		return nil, err
	}
	return &request, nil
}

// GetAll retrieves all UserVerificationRequests with optional pagination
func (s *UserVerificationRequestService) GetAll(ctx context.Context, page, limit int64) ([]*models.UserVerificationRequest, error) {
	var requests []*models.UserVerificationRequest

	opts := options.Find().
		SetSkip((page - 1) * limit).
		SetLimit(limit)

	cursor, err := s.collection.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var request models.UserVerificationRequest
		if err := cursor.Decode(&request); err != nil {
			return nil, err
		}
		requests = append(requests, &request)
	}

	return requests, cursor.Err()
}

// Update updates a UserVerificationRequest by ID
func (s *UserVerificationRequestService) Update(ctx context.Context, id primitive.ObjectID, updateData bson.M) (*mongo.UpdateResult, error) {
	updateData["updatedAt"] = time.Now() // Update the timestamp
	return s.collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": updateData})
}

// Delete deletes a UserVerificationRequest by ID
func (s *UserVerificationRequestService) Delete(ctx context.Context, id primitive.ObjectID) (*mongo.DeleteResult, error) {
	return s.collection.DeleteOne(ctx, bson.M{"_id": id})
}
