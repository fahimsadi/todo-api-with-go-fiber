package routes

import (
	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

// SwaggerInit godoc
// @title TODO API
// @version 1.0
// @description This is a simple TODO API with Fiber and Swagger.
// @host localhost:3000
// @BasePath /
func SwaggerInit(app *fiber.App) {

	app.Get("/swagger/*", fiberSwagger.WrapHandler)

}
