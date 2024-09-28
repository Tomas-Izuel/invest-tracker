package services

import (
	"invest/models"
	"invest/models/dto"
	"invest/repository"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// StockService define los métodos del servicio de stock
type StockService struct {
	Client *mongo.Client
}

// NewStockService crea una nueva instancia del servicio de stock
func NewStockService(client *mongo.Client) *StockService {
	return &StockService{Client: client}
}

// AddStockToAccount agrega una acción a una cuenta existente
func (s *StockService) AddStockToAccount(accountID string, createStockData dto.CreateStockDto) error {
	// Convertir el accountID a un ObjectID de MongoDB
	accountObjectID, err := primitive.ObjectIDFromHex(accountID)
	if err != nil {
		return err
	}

	// Crear una nueva acción
	stock := models.Stock{
		AccountID:  accountObjectID,
		StockName:  createStockData.StockName,
		IsBalanz:   createStockData.IsBalanz,
		Name:       createStockData.Name,
		Type:       createStockData.Type,
		Quotations: []primitive.ObjectID{}, // Inicializa con un array vacío de ObjectIDs para las cotizaciones
	}

	// Agregar la acción a la base de datos
	createdStock, err := repository.CreateStock(s.Client, stock)
	if err != nil {
		log.Println("Error creating stock:", err)
		return err
	}

	// Relacionar la nueva acción con la cuenta usando $push
	_, err = repository.UpdateAccount(s.Client, accountID, createdStock.InsertedID.(primitive.ObjectID))
	if err != nil {
		log.Println("Error updating account:", err)
		return err
	}

	return nil
}
