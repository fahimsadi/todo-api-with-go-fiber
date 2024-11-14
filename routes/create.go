package routes

import "github.com/gofiber/fiber/v2"

// TodoCreate godoc
// @Description Create a new TODO item with a title and set status to pending
// @Tags todos
// @Accept json
// @Produce json
// @Param todo body Todo true "Todo object"
// @Success 201 {object} Todo
// @Failure 400 {object} map[string]interface{}
// @Router /todos [post]
func TodoCreate(app *fiber.App) {

	app.Post("/todos", func(c *fiber.Ctx) error {
		var newTodo Todo
		if err := c.BodyParser(&newTodo); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
		}

		newTodo.ID = idCounter
		idCounter++
		newTodo.Status = "pending"

		todos = append(todos, newTodo)
		return c.Status(fiber.StatusCreated).JSON(newTodo)
	})
}
