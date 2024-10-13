package repository

import (
	"context"
	"invest/config"
	"invest/errors"
	"invest/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertAccountType(ctx context.Context, accountType *models.AccountType) (*mongo.InsertOneResult, error) {
	collection := config.GetCollection("account_type")

	created, err := collection.InsertOne(ctx, accountType)
	if err != nil {
		return nil, errors.Wrap(500, "failed to insert account type", err)
	}

	return created, nil
}

func FindAccountTypes(ctx context.Context) ([]models.AccountType, error) {
	collection := config.GetCollection("account_type")

	cursor, err := collection.Find(ctx, nil)
	if err != nil {
		return nil, errors.Wrap(500, "failed to get account types", err)
	}

	var accountTypes []models.AccountType
	if err = cursor.All(ctx, &accountTypes); err != nil {
		return nil, errors.Wrap(500, "failed to get account types", err)
	}

	return accountTypes, nil
}

func FindAccountTypeByID(ctx context.Context, id string) (*models.AccountType, error) {
	collection := config.GetCollection("account_type")

	accountTypeID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.Wrap(400, "invalid account type ID format", err)
	}

	var accountType models.AccountType
	err = collection.FindOne(ctx, bson.M{"_id": accountTypeID}).Decode(&accountType)
	if err != nil {
		return nil, errors.ErrNotFound
	}

	return &accountType, nil
}

func UpdateAccountType(ctx context.Context, id string, accountType *models.AccountType) (*mongo.UpdateResult, error) {
	collection := config.GetCollection("account_type")

	accountTypeID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.Wrap(400, "invalid account type ID format", err)
	}

	updated, err := collection.ReplaceOne(ctx, bson.M{"_id": accountTypeID}, accountType)
	if err != nil {
		return nil, errors.Wrap(500, "failed to update account type", err)
	}

	return updated, nil
}
