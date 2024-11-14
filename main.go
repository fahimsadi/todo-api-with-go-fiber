package main

import (
	"github.com/gofiber/fiber/v2"
	_ "todo-app/docs"
	"todo-app/routes"
)

func main() {
	app := fiber.New()

	routes.SwaggerInit(app)
	routes.TodoCreate(app)
	routes.TodoGet(app)
	routes.TodoGetByID(app)
	routes.ToDoDelete(app)
	routes.ToDoUpdate(app)

	err := app.Listen(":3000")
	if err != nil {
		return
	}
}
