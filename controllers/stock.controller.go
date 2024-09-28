package controllers

import (
	"encoding/json"
	"invest/models/dto"
	"invest/services"
	"net/http"
)

// StockController define los métodos del controlador de stock
type StockController struct {
	StockService *services.StockService
}

// NewStockController crea una nueva instancia del controlador de stock
func NewStockController(stockService *services.StockService) *StockController {
	return &StockController{StockService: stockService}
}

// AddStock maneja la solicitud HTTP para agregar una acción a una cuenta
func (sc *StockController) AddStock(w http.ResponseWriter, r *http.Request) {
	var requestBody dto.CreateStockDto

	// Decodificar el cuerpo de la solicitud
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Hardcodear el ID de la cuenta por el momento
	const accountID = "66f88a99cee86ff05e70a4ca"

	// Llamar al servicio para agregar la acción a la cuenta
	err = sc.StockService.AddStockToAccount(accountID, requestBody)
	if err != nil {
		http.Error(w, "Error adding stock", http.StatusInternalServerError)
		return
	}

	// Responder con un éxito
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Stock added successfully"))
}
