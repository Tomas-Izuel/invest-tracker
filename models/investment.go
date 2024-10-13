package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Investment struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name"`
	Code      string             `bson:"code"`
	Stock     int                `bson:"stock"`
	AccountID primitive.ObjectID `bson:"account_id"`
	Prices    []InvestmentPrice  `bson:"prices,omitempty"` // Lista de precios (cotizaciones)
}
