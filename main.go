package main

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

var todos = []Todo{} // Slices as test database for now. Will replace with redis or other db
var idCounter = 1

func main() {
	app := fiber.New()

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

	app.Get("/todos", func(c *fiber.Ctx) error {
		return c.JSON(todos)
	})

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

	app.Listen(":3000")
}
