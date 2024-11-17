package routes

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"todo-app/db"
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

		id := c.Params("id")

		todoJSON, err := db.Get("todo:" + id)
		if err != nil {
			// Check if the error is because the key doesn't exist
			if err.Error() == "redis: nil" {
				return c.Status(404).SendString("Todo not found")
			}
			return c.Status(500).SendString("Failed to fetch todo")
		}

		var existingTodo Todo
		if err := json.Unmarshal([]byte(todoJSON), &existingTodo); err != nil {
			return c.Status(500).SendString("Failed to process existing todo")
		}

		var updatedTodo Todo
		if err := c.BodyParser(&updatedTodo); err != nil {
			return c.Status(400).SendString("Invalid request body")
		}

		if updatedTodo.Task != "" {
			existingTodo.Task = updatedTodo.Task
		}
		existingTodo.Status = updatedTodo.Status // Update status

		updatedTodoJSON, err := json.Marshal(existingTodo)
		if err != nil {
			return c.Status(500).SendString("Failed to process updated todo")
		}

		if err := db.Set("todo:"+id, string(updatedTodoJSON)); err != nil {
			return c.Status(500).SendString("Failed to save updated todo")
		}

		return c.JSON(existingTodo)
	})
}
