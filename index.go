package main

import (
	"fmt"
	"invest/config"
	"invest/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
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

	app := fiber.New()

	routes.RouteHandler(app)

	log.Fatal(app.Listen(":3000"))
}
