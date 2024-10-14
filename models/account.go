package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Account struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Period string             `bson:"period"`
	UserID primitive.ObjectID `bson:"user_id"`
	Type   primitive.ObjectID `bson:"account_type_id"`
}
