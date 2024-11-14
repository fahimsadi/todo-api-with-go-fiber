package routes

import "github.com/gofiber/fiber/v2"

// TodoGet godoc
// @Summary Get all TODOs
// @Description Retrieve a list of all TODO items
// @Tags todos
// @Produce json
// @Success 200 {array} Todo
// @Router /todos [get]
func TodoGet(app *fiber.App) {
	app.Get("/todos", func(c *fiber.Ctx) error {
		return c.JSON(todos)
	})
}
