package main

import (
	"os"

	"github.com/100lvlmaster/5chan-go/database"
	"github.com/100lvlmaster/5chan-go/middleware"
	"github.com/100lvlmaster/5chan-go/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()
	godotenv.Load()
	app.Use(cors.New())
	// app.Use(cache.New())
	app.Use(middleware.CustomeMiddleware())
	database.DBConn = database.InitDb()
	routes.RoutesInit(app)
	port, envExists := os.LookupEnv("PORT")
	if !envExists {
		port = "8080"
	}
	app.Listen(":" + port)
}
