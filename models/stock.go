package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Stock struct {
	ID         primitive.ObjectID   `bson:"_id,omitempty"`
	AccountID  primitive.ObjectID   `bson:"account_id"`           // Relación con la cuenta
	StockName  string               `bson:"stock_name"`           // Nombre de la acción (ej. SPY)
	EndPoint   string               `bson:"end_point"`            // Punto final de la API (ej. https://api.iextrading.com/1.0/stock/spy/quote)
	Name       string               `bson:"name"`                 // Nombre de la acción (ej. Standard & Poor's Depositary Receipts)
	Quotations []primitive.ObjectID `bson:"quotations,omitempty"` // Lista de IDs de cotizaciones
}
