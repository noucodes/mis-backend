package main

import (
	"log"
	"os"

	"github.com/noucodes/mis-backend/config"
	"github.com/noucodes/mis-backend/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDB()

	app := fiber.New()
	routes.SetupRoutes(app)

	log.Println("🚀 Server running on port", os.Getenv("PORT"))
	app.Listen(":" + os.Getenv("PORT"))
}