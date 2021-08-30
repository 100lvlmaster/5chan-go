package controllers

import (
	"errors"

	"github.com/100lvlmaster/5-chan-go/database"
	"github.com/100lvlmaster/5-chan-go/models"
	"github.com/gofiber/fiber/v2"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

func GetReplies(c *fiber.Ctx) error {
	var replies []models.Reply
	database.DBConn.Find(&replies)
	return c.JSON(replies)
}

///
func GetRepliesByPostId(c *fiber.Ctx) error {
	id := c.Params("id")
	var replies []models.Reply
	database.DBConn.Where(&models.Reply{PostID: id}).Find(&replies)
	return c.JSON(replies)
}

///
func CreateReply(c *fiber.Ctx) error {
	var reply models.Reply
	if err := c.BodyParser(&reply); err != nil {
		return c.Status(400).SendString("Incorrect input body, please chcek input convention and try again")
	}
	/// DB contains post
	var result models.Post
	err := database.DBConn.First(&result, reply.PostID).Error
	hasRecord := !errors.Is(err, gorm.ErrRecordNotFound)
	if !hasRecord {
		return c.Status(400).SendString("Post does not exist")
	}

	///
	id, _ := gonanoid.Generate("0123456789", 10)
	reply.ID = id
	database.DBConn.Create(&reply)
	return c.Status(201).JSON(reply)
}
