package repository

import (
	"context"
	"invest/config"
	"invest/errors"
	"invest/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// InsertAccount inserts a new account into the database
func InsertAccount(ctx context.Context, account *models.Account) (*mongo.InsertOneResult, error) {
	collection := config.GetCollection("accounts")

	created, err := collection.InsertOne(ctx, account)
	if err != nil {
		return nil, errors.Wrap(500, "failed to insert account", err)
	}

	return created, nil
}

// FindAccountByID finds an account by its ID
func FindAccountByID(ctx context.Context, id string) (*models.Account, error) {
	collection := config.GetCollection("accounts")

	accountID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.Wrap(400, "invalid account ID format", err)
	}

	// Pipeline de agregación con $lookup para obtener las inversiones y las cotizaciones
	pipeline := mongo.Pipeline{
		bson.D{{"$match", bson.D{{"_id", accountID}}}}, // Filtrar por account ID
		bson.D{{"$lookup", bson.D{
			{"from", "investments"},        // Colección de inversiones
			{"localField", "_id"},          // Campo local (account._id)
			{"foreignField", "account_id"}, // Campo de la inversión que se relaciona con account_id
			{"as", "investments"},          // El resultado será una lista de investments en cada account
			{"pipeline", mongo.Pipeline{ // Subpipeline para las inversiones
				bson.D{{"$lookup", bson.D{
					{"from", "investment_prices"},     // Colección de precios
					{"localField", "_id"},             // Campo local (investment._id)
					{"foreignField", "investment_id"}, // Campo de la cotización relacionado con la inversión
					{"as", "prices"},                  // El resultado será una lista de precios
					{"pipeline", mongo.Pipeline{ // Subpipeline para traer las últimas 5 cotizaciones
						bson.D{{"$sort", bson.D{{"date", -1}}}}, // Ordenar por fecha descendente
						bson.D{{"$limit", 5}},                   // Limitar a las últimas 5 cotizaciones
					}},
				}}},
			}},
		}}},
	}

	// Ejecutar la agregación
	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, errors.Wrap(500, "failed to aggregate account with investments and prices", err)
	}

	var accounts []models.Account
	if err = cursor.All(ctx, &accounts); err != nil {
		return nil, errors.Wrap(500, "failed to decode account", err)
	}

	// Asegurarse de que solo devolvemos una cuenta (porque _id es único)
	if len(accounts) == 0 {
		return nil, errors.ErrNotFound
	}

	return &accounts[0], nil
}

// DeleteAccount deletes an account by its ID
func DeleteAccount(ctx context.Context, id string) error {
	collection := config.GetCollection("accounts")

	accountID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.Wrap(400, "invalid account ID format", err)
	}

	_, err = collection.DeleteOne(ctx, bson.M{"_id": accountID})
	if err != nil {
		return errors.Wrap(500, "failed to delete account", err)
	}

	return nil
}

// GetAllAccountsByUserID returns all accounts related to a user
func GetAllAccountsByUserID(ctx context.Context, userID string) ([]models.Account, error) {
	collection := config.GetCollection("accounts")

	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, errors.Wrap(400, "invalid user ID format", err)
	}

	// Pipeline de agregación con $lookup para obtener las inversiones relacionadas
	pipeline := mongo.Pipeline{
		bson.D{{"$match", bson.D{{"user_id", userObjectID}}}}, // Filtrar por user_id
		bson.D{{"$lookup", bson.D{
			{"from", "investments"},        // Colección de inversiones
			{"localField", "_id"},          // Campo local (account._id)
			{"foreignField", "account_id"}, // Campo de la inversión que se relaciona con account_id
			{"as", "investments"},          // El resultado será una lista de investments en cada account
		}}},
	}

	// Ejecutar la agregación
	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, errors.Wrap(500, "failed to aggregate accounts", err)
	}

	var accounts []models.Account
	if err = cursor.All(ctx, &accounts); err != nil {
		return nil, errors.Wrap(500, "failed to decode accounts", err)
	}

	return accounts, nil
}

// UpdateAccount updates an account's data
func UpdateAccount(ctx context.Context, id string, updateData bson.M) (*mongo.UpdateResult, error) {
	collection := config.GetCollection("accounts")

	accountID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.Wrap(400, "invalid account ID format", err)
	}

	updated, err := collection.UpdateOne(ctx, bson.M{"_id": accountID}, bson.M{"$set": updateData})
	if err != nil {
		return nil, errors.Wrap(500, "failed to update account", err)
	}

	return updated, nil
}

func GetAllAccounts(ctx context.Context) ([]models.Account, error) {
	collection := config.GetCollection("accounts")

	pipeline := mongo.Pipeline{
		bson.D{{"$lookup", bson.D{
			{"from", "account_types"},
			{"localField", "type"},
			{"foreignField", "_id"},
			{"as", "type"},
		}}},
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, errors.Wrap(500, "failed to aggregate accounts", err)
	}

	var accounts []models.Account
	if err = cursor.All(ctx, &accounts); err != nil {
		return nil, errors.Wrap(500, "failed to decode accounts", err)
	}


	return accounts, nil
}