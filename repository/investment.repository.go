package repository

import (
	"context"
	"invest/config"
	"invest/errors"
	"invest/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsertInvestment inserts a new investment into the database
func InsertInvestment(ctx context.Context, investment *models.Investment) error {
	collection := config.GetCollection("investments")

	_, err := collection.InsertOne(ctx, investment)
	if err != nil {
		return errors.Wrap(500, "failed to insert investment", err)
	}

	return nil
}

// FindInvestmentByID finds an investment by its ID
func FindInvestmentByID(ctx context.Context, id string) (*models.Investment, error) {
	collection := config.GetCollection("investments")

	investmentID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.Wrap(400, "invalid investment ID format", err)
	}

	var investment models.Investment
	err = collection.FindOne(ctx, bson.M{"_id": investmentID}).Decode(&investment)
	if err != nil {
		return nil, errors.ErrNotFound
	}

	return &investment, nil
}

// UpdateInvestment updates an investment's data
func UpdateInvestment(ctx context.Context, id string, updateData bson.M) error {
	collection := config.GetCollection("investments")

	investmentID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.Wrap(400, "invalid investment ID format", err)
	}

	_, err = collection.UpdateOne(ctx, bson.M{"_id": investmentID}, bson.M{"$set": updateData})
	if err != nil {
		return errors.Wrap(500, "failed to update investment", err)
	}

	return nil
}

// DeleteInvestment deletes an investment by its ID
func DeleteInvestment(ctx context.Context, id string) error {
	collection := config.GetCollection("investments")

	investmentID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.Wrap(400, "invalid investment ID format", err)
	}

	_, err = collection.DeleteOne(ctx, bson.M{"_id": investmentID})
	if err != nil {
		return errors.Wrap(500, "failed to delete investment", err)
	}

	return nil
}

func GetAllInvestmentByAccountID(ctx context.Context, accountID string) ([]models.Investment, error) {
	collection := config.GetCollection("investments")

	accountObjectID, err := primitive.ObjectIDFromHex(accountID)
	if err != nil {
		return nil, errors.Wrap(400, "invalid account ID format", err)
	}

	cursor, err := collection.Find(ctx, bson.M{"account_id": accountObjectID})
	if err != nil {
		return nil, errors.Wrap(500, "failed to get investments", err)
	}

	var investments []models.Investment
	if err = cursor.All(ctx, &investments); err != nil {
		return nil, errors.Wrap(500, "failed to get investments", err)
	}

	return investments, nil
}
