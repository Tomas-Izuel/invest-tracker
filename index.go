package main

import (
	"fmt"
	"invest/config"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
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
	config.InitializeDB(mongoURI, dbName)

	// Iniciar el servidor
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
