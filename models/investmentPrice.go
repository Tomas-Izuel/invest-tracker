package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InvestmentPrice struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	TotalPrice   float64            `bson:"total_price"`
	Date         time.Time          `bson:"date"`
	ThisStock   int                `bson:"this_stock"`
	InvestmentID primitive.ObjectID `bson:"investment_id"`
}
