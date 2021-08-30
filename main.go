package main

import (
	"github.com/100lvlmaster/5-chan-go/database"
	"github.com/100lvlmaster/5-chan-go/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	database.DBConn = database.InitDb()
	routes.RoutesInit(app)
	app.Listen(":8080")
}
