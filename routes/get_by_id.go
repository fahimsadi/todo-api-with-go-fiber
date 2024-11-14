package routes

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
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
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid ID"})
		}

		for _, todo := range todos {
			if todo.ID == id {
				return c.JSON(todo)
			}
		}
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "TODO not found"})
	})
}
