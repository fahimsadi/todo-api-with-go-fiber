package routes

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
)

// ToDoUpdate godoc
// @Summary Update a TODO by ID
// @Description Update the title and status of a TODO item by its ID
// @Tags todos
// @Accept json
// @Produce json
// @Param id path int true "Todo ID"
// @Param todo body Todo true "Updated Todo object"
// @Success 200 {object} Todo
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /todos/{id} [put]
func ToDoUpdate(app *fiber.App) {

	app.Put("/todos/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid ID"})
		}

		var updatedTodo Todo
		if err := c.BodyParser(&updatedTodo); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
		}

		for i, todo := range todos {
			if todo.ID == id {
				todos[i].Title = updatedTodo.Title
				todos[i].Status = updatedTodo.Status
				return c.JSON(todos[i])
			}
		}
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "TODO not found"})
	})

}
