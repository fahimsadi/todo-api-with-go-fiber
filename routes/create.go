package routes

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"log"
	"todo-app/db"
)

func TodoCreate(app *fiber.App) {
	app.Post("/todos", func(c *fiber.Ctx) error {
		log.Println("Request Body:", string(c.Body()))
		var todo Todo

		if err := c.BodyParser(&todo); err != nil {
			return c.Status(400).SendString("Invalid request body")
		}

		todo.ID = uuid.New().String()

		todoJSON, err := json.Marshal(todo)
		if err != nil {
			return c.Status(500).SendString("Failed to process todo")
		}

		if err := db.Set("todo:"+todo.ID, string(todoJSON)); err != nil {
			return c.Status(500).SendString("Failed to save todo")
		}

		return c.Status(201).JSON(todo)
	})
}
