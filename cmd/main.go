package main

import (
	"fmt"
	"nutrimama/internal/config"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	if (os.Getenv("APP_ENV")) != "production" {
		godotenvErr := godotenv.Load("./.env")
		if godotenvErr != nil {
			panic("Failed to load .env file: " + godotenvErr.Error())
		}
	}

	db, err := config.ConnectDB()
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET, POST, PUT, DELETE, PATCH, OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization, X-Requested-With, Access-Control-Request-Method, Access-Control-Request-Headers, ngrok-skip-browser-warning",
		ExposeHeaders:    "Content-Length, Content-Type",
		AllowCredentials: false,
	}))
	config.Bootstrap(&config.BootstrapConfig{
		DB:  db,
		App: app,
	})

	webPort := os.Getenv("APPPORT")
	err = app.Listen(":" + webPort)
	fmt.Printf("Server running on port %s\n", webPort)
	if err != nil {
		panic("Failed to start server: " + err.Error())
	}
}
