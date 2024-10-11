package repository

import (
	"context"
	"invest/config"
	"invest/errors"
	"invest/models"
)

// InsertInvestment inserts a new investment into the database
func InsertInvestmentPrice(ctx context.Context, investment *models.Investment) error {
	collection := config.GetCollection("investment_prices")

	_, err := collection.InsertOne(ctx, investment)
	if err != nil {
		return errors.Wrap(500, "failed to insert investment price", err)
	}

	return nil
}
