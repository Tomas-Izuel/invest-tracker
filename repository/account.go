package repository

import (
	"context"
	"invest/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateAccount inserta una nueva cuenta en la base de datos
func CreateAccount(client *mongo.Client, account models.Account) (*mongo.InsertOneResult, error) {
	collection := client.Database("investmentdb").Collection("accounts")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, account)
	if err != nil {
		return nil, err
	}

	log.Println("Inserted a new account: ", result.InsertedID)
	return result, nil
}

// GetAccountByID obtiene una cuenta por su ID
func GetAccountByID(client *mongo.Client, id string) (models.Account, error) {
	collection := client.Database("investmentdb").Collection("accounts")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var account models.Account
	objID, _ := primitive.ObjectIDFromHex(id)
	err := collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&account)
	if err != nil {
		return models.Account{}, err
	}

	return account, nil
}

// UpdateAccount actualiza una cuenta existente
func UpdateAccount(client *mongo.Client, accountID string, stockID primitive.ObjectID) (*mongo.UpdateResult, error) {
	collection := client.Database("investmentdb").Collection("accounts")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Convertir el accountID a un ObjectID de MongoDB
	accountObjectID, err := primitive.ObjectIDFromHex(accountID)
	if err != nil {
		return nil, err
	}

	// Actualizar usando $push para agregar el nuevo stock al array de stocks
	result, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": accountObjectID}, // Filtro por el ID de la cuenta
		bson.M{"$push": bson.M{"stocks": stockID}}, // Usa $push para agregar el stock al array de stocks
	)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// DeleteAccount elimina una cuenta por su ID
func DeleteAccount(client *mongo.Client, id string) (*mongo.DeleteResult, error) {
	collection := client.Database("investmentdb").Collection("accounts")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, _ := primitive.ObjectIDFromHex(id)
	result, err := collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return nil, err
	}

	return result, nil
}
