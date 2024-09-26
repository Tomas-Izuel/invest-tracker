package main

import (
	"fmt"
	"invest/config"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	const accountId = "66f4a80755c8e29142521d4b"

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

	print(client)
}
