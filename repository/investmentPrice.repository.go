package repository

import (
	"context"
	"invest/config"
	"invest/errors"
	"invest/models"
)

// InsertInvestment inserts a new investment into the database
func InsertInvestmentPrice(ctx context.Context, investmentPrice *models.InvestmentPrice) error {
	collection := config.GetCollection("investment_prices")

	_, err := collection.InsertOne(ctx, investmentPrice)
	if err != nil {
		return errors.Wrap(500, "failed to insert investment price", err)
	}

	return nil
}
