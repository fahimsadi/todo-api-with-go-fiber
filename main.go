package main

import (
	"github.com/gofiber/fiber/v2"
	"os"
	"todo-app/db"
	_ "todo-app/docs"
	"todo-app/routes"
)

func main() {

	db.Initialize()
	//defer db.Close()

	app := fiber.New()
	routes.InitializeRoutes(app)

	err := app.Listen(os.Getenv("APP_PORT"))
	if err != nil {
		return
	}
}
