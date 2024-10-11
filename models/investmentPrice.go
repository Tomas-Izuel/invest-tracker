package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InvestmentPrice struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	TotalPrice   float64            `bson:"total_price"`
	Date         time.Time          `bson:"date"`
	InvestmentID primitive.ObjectID `bson:"investment_id"`
}
