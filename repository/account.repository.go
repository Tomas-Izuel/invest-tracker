package repository

import (
	"context"
	"invest/config"
	"invest/errors"
	"invest/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsertAccount inserts a new account into the database
func InsertAccount(ctx context.Context, account *models.Account) error {
	collection := config.GetCollection("accounts")

	_, err := collection.InsertOne(ctx, account)
	if err != nil {
		return errors.Wrap(500, "failed to insert account", err)
	}

	return nil
}

// FindAccountByID finds an account by its ID
func FindAccountByID(ctx context.Context, id string) (*models.Account, error) {
	collection := config.GetCollection("accounts")

	accountID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.Wrap(400, "invalid account ID format", err)
	}

	var account models.Account
	err = collection.FindOne(ctx, bson.M{"_id": accountID}).Decode(&account)
	if err != nil {
		return nil, errors.ErrNotFound
	}

	return &account, nil
}

// UpdateAccount updates an account's data
func UpdateAccount(ctx context.Context, id string, updateData bson.M) error {
	collection := config.GetCollection("accounts")

	accountID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.Wrap(400, "invalid account ID format", err)
	}

	_, err = collection.UpdateOne(ctx, bson.M{"_id": accountID}, bson.M{"$set": updateData})
	if err != nil {
		return errors.Wrap(500, "failed to update account", err)
	}

	return nil
}

// DeleteAccount deletes an account by its ID
func DeleteAccount(ctx context.Context, id string) error {
	collection := config.GetCollection("accounts")

	accountID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.Wrap(400, "invalid account ID format", err)
	}

	_, err = collection.DeleteOne(ctx, bson.M{"_id": accountID})
	if err != nil {
		return errors.Wrap(500, "failed to delete account", err)
	}

	return nil
}
