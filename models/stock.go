package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Stock struct {
	ID         primitive.ObjectID   `bson:"_id,omitempty"`
	AccountID  primitive.ObjectID   `bson:"account_id"`           // Relación con la cuenta
	StockName  string               `bson:"stock_name"`           // Nombre de la acción (ej. SPY)
	Name       string               `bson:"name"`                 // Nombre de la acción (ej. Standard & Poor's Depositary Receipts)
	IsBalanz   bool                 `bson:"is_balanz,omitempty"`  // Indica si es una acción de Balanz
	Type       string               `bson:"type,omitempty"`       // Tipo de acción (ej. ETF)
	Quotations []primitive.ObjectID `bson:"quotations,omitempty"` // Lista de IDs de cotizaciones
}
