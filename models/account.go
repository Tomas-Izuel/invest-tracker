package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Account struct {
	ID     primitive.ObjectID   `bson:"_id,omitempty"`
	Name   string               `bson:"name"`
	Stocks []primitive.ObjectID `bson:"stocks,omitempty"` // Lista de IDs de acciones
}
