package routes

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
)

// ToDoDelete godoc
// @Summary Delete a TODO by ID
// @Description Delete a specific TODO item by its ID
// @Tags todos
// @Param id path int true "Todo ID"
// @Success 204 "No Content"
// @Failure 400 {object} map[string]interface{}
// @Router /todos/{id} [delete]
func ToDoDelete(app *fiber.App) {

	app.Delete("/todos/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid ID"})
		}

		for i, todo := range todos {
			if todo.ID == id {
				todos = append(todos[:i], todos[i+1:]...)
				return c.SendStatus(fiber.StatusNoContent)
			}
		}
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "TODO not found"})
	})

}
