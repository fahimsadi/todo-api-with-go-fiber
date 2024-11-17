package routes

import "github.com/gofiber/fiber/v2"

func InitializeRoutes(app *fiber.App) {
	SwaggerInit(app)
	TodoCreate(app)
	TodoGet(app)
	TodoGetByID(app)
	ToDoDelete(app)
	ToDoUpdate(app)
}
