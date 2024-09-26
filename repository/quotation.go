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

// CreateQuotation inserta una nueva cotización en la base de datos
func CreateQuotation(client *mongo.Client, quotation models.Quotation) (*mongo.InsertOneResult, error) {
	collection := client.Database("investmentdb").Collection("quotations")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, quotation)
	if err != nil {
		return nil, err
	}

	log.Println("Inserted a new quotation: ", result.InsertedID)
	return result, nil
}

// GetQuotationsByStockID obtiene todas las cotizaciones de una acción
func GetQuotationsByStockID(client *mongo.Client, stockID string) ([]models.Quotation, error) {
	collection := client.Database("investmentdb").Collection("quotations")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var quotations []models.Quotation
	objID, _ := primitive.ObjectIDFromHex(stockID)
	cursor, err := collection.Find(ctx, bson.M{"stock_id": objID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var quotation models.Quotation
		if err = cursor.Decode(&quotation); err != nil {
			return nil, err
		}
		quotations = append(quotations, quotation)
	}

	return quotations, nil
}

// UpdateQuotation actualiza una cotización existente
func UpdateQuotation(client *mongo.Client, id string, update bson.M) (*mongo.UpdateResult, error) {
	collection := client.Database("investmentdb").Collection("quotations")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, _ := primitive.ObjectIDFromHex(id)
	result, err := collection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": update})
	if err != nil {
		return nil, err
	}

	return result, nil
}

// DeleteQuotation elimina una cotización por su ID
func DeleteQuotation(client *mongo.Client, id string) (*mongo.DeleteResult, error) {
	collection := client.Database("investmentdb").Collection("quotations")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, _ := primitive.ObjectIDFromHex(id)
	result, err := collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return nil, err
	}

	return result, nil
}
