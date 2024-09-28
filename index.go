package main

import (
	"fmt"
	"invest/config"
	"invest/controllers"
	"invest/services"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	const accountId = "66f88a99cee86ff05e70a4ca"

	// Carga las variables de entorno desde el archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Obtiene las variables de entorno
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	// Construye el string de conexión usando las variables de entorno
	mongoURI := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority&appName=%s", username, password, host, dbName)

	// Inicializa la conexión con MongoDB
	client, err := config.ConnectMongo(mongoURI)
	if err != nil {
		log.Fatal("Error connecting to MongoDB: ", err)
	}

	// Inicializa los servicios
	stockService := services.NewStockService(client)

	// Inicializa los controladores
	stockController := controllers.NewStockController(stockService)

	// Ruta para agregar una acción
	http.HandleFunc("/add-stock", stockController.AddStock)

	// Iniciar el servidor
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
