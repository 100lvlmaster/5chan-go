package routes

import (
	"github.com/100lvlmaster/5-chan-go/controllers"
	"github.com/gofiber/fiber/v2"
)

func RoutesInit(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hellso, World 👋!")
	})
	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Get("/posts", controllers.GetPosts)
	v1.Post("/posts", controllers.CreatePost)
	v1.Get("/posts/:id", controllers.GetPostById)
	v1.Get("/replies", controllers.GetReplies)
	v1.Get("/replies/:id", controllers.GetRepliesByPostId)
	v1.Post("/replies", controllers.CreateReply)
}
