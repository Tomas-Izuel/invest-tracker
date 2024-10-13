package repository

import (
	"context"
	"invest/config"
	"invest/errors"
	"invest/models"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// InsertUser inserts a new user into the MongoDB database
func InsertUser(ctx context.Context, user *models.User) error {
	collection := config.GetCollection("users")

	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		return errors.Wrap(500, "failed to insert user", err)
	}

	return nil
}

// FindUserByID finds a user by their ID
func FindUserByID(ctx context.Context, id string) (*models.User, error) {
	collection := config.GetCollection("users")

	userID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.Wrap(400, "invalid user ID format", err)
	}

	var user models.User
	err = collection.FindOne(ctx, bson.M{"_id": userID}).Decode(&user)
	if err != nil {
		return nil, errors.ErrNotFound
	}

	return &user, nil
}

func FindUserByName(ctx context.Context, name string) (*models.User, error) {
	collection := config.GetCollection("users")

	var user models.User
	err := collection.FindOne(ctx, bson.M{"name": name}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		log.Println(err)
		return nil, errors.ErrNotFound
	}

	return &user, nil
}
