package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Account struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Period      string             `bson:"period"`
	EndpointURL string             `bson:"endpoint_url"`
	HasLogin    bool               `bson:"has_login"`
	Username    *string            `bson:"username,omitempty"`
	Password    *string            `bson:"password,omitempty"`
	UserID      primitive.ObjectID `bson:"user_id"`
}
