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

// CreateStock inserta una nueva acción en la base de datos
func CreateStock(client *mongo.Client, stock models.Stock) (*mongo.InsertOneResult, error) {
	collection := client.Database("investmentdb").Collection("stocks")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, stock)
	if err != nil {
		return nil, err
	}

	log.Println("Inserted a new stock: ", result.InsertedID)
	return result, nil
}

// GetStockByID obtiene una acción por su ID
func GetStockByID(client *mongo.Client, id string) (models.Stock, error) {
	collection := client.Database("investmentdb").Collection("stocks")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var stock models.Stock
	objID, _ := primitive.ObjectIDFromHex(id)
	err := collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&stock)
	if err != nil {
		return models.Stock{}, err
	}

	return stock, nil
}

// UpdateStock actualiza una acción existente
func UpdateStock(client *mongo.Client, id string, update bson.M) (*mongo.UpdateResult, error) {
	collection := client.Database("investmentdb").Collection("stocks")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, _ := primitive.ObjectIDFromHex(id)
	result, err := collection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": update})
	if err != nil {
		return nil, err
	}

	return result, nil
}

// DeleteStock elimina una acción por su ID
func DeleteStock(client *mongo.Client, id string) (*mongo.DeleteResult, error) {
	collection := client.Database("investmentdb").Collection("stocks")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, _ := primitive.ObjectIDFromHex(id)
	result, err := collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return nil, err
	}

	return result, nil
}
