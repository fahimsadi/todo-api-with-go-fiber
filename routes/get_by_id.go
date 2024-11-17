package routes

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"todo-app/db"
)

// TodoGetByID
// @Summary Get a TODO by ID
// @Description Retrieve a specific TODO item by its ID
// @Tags todos
// @Produce json
// @Param id path int true "Todo ID"
// @Success 200 {object} Todo
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /todos/{id} [get]
func TodoGetByID(app *fiber.App) {
	app.Get("/todos/:id", func(c *fiber.Ctx) error {

		id := c.Params("id")

		todoJSON, err := db.Get("todo:" + id)

		if err != nil {

			if err.Error() == "redis: nil" {
				return c.Status(404).SendString("Todo not found")
			}
			return c.Status(500).SendString("Failed to fetch todo")
		}

		var todo Todo
		if err := json.Unmarshal([]byte(todoJSON), &todo); err != nil {
			return c.Status(500).SendString("Failed to process todo")
		}

		return c.JSON(todo)
	})
}
