package routes

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"todo-app/db"
)

// TodoGet godoc
// @Summary Get all TODOs
// @Description Retrieve a list of all TODO items
// @Tags todos
// @Produce json
// @Success 200 {array} Todo
// @Router /todos [get]
func TodoGet(app *fiber.App) {
	app.Get("/todos", func(c *fiber.Ctx) error {
		// Get all keys matching "todo:*"
		keys, err := db.Keys("todo:*")
		if err != nil {
			return c.Status(500).SendString("Failed to retrieve todos")
		}

		todos := []Todo{}

		// Fetch and deserialize each todo
		for _, key := range keys {
			todoJSON, err := db.Get(key)
			if err != nil {
				continue // Skip if there's an error fetching the todo
			}

			var todo Todo
			if err := json.Unmarshal([]byte(todoJSON), &todo); err == nil {
				todos = append(todos, todo)
			}
		}

		// Return the list of todos
		return c.JSON(todos)
	})
}
