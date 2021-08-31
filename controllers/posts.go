package controllers

import (
	"github.com/100lvlmaster/5chan-go/database"
	"github.com/100lvlmaster/5chan-go/models"
	"github.com/gofiber/fiber/v2"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func GetPosts(c *fiber.Ctx) error {
	var posts []models.Post
	database.DBConn.Find(&posts)
	return c.JSON(posts)
}

func GetPostById(c *fiber.Ctx) error {
	id := c.Params("id")
	var post models.Post
	database.DBConn.Preload("Reply").First(&post, id)
	return c.JSON(post)
}

func CreatePost(c *fiber.Ctx) error {
	var post models.Post
	if err := c.BodyParser(&post); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	id, _ := gonanoid.Generate("0123456789", 10)
	post.ID = id
	database.DBConn.Create(&post)
	return c.Status(201).JSON(post)
}

func DeletePost(c *fiber.Ctx) error {
	id := c.Params("id")
	var post models.Post
	database.DBConn.Delete(&post, id)
	return c.Status(200).SendString("Post deleted successfully")
}
