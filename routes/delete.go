package routes

import (
	"github.com/gofiber/fiber/v2"
	"todo-app/db"
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

		id := c.Params("id")

		err := db.Delete("todo:" + id)
		if err != nil {

			if err.Error() == "redis: nil" {
				return c.Status(404).SendString("Todo not found")
			}
			return c.Status(500).SendString("Failed to delete todo")
		}

		return c.SendString("Todo deleted successfully")
	})
}
