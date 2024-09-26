package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Quotation struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	StockID primitive.ObjectID `bson:"stock_id"` // Relación con la acción
	Price   float64            `bson:"price"`    // Precio de la cotización
	Date    time.Time          `bson:"date"`     // Fecha de la cotización
}
