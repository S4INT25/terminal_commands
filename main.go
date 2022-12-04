package main

import (
	"github.com/gofiber/fiber/v2"
	"terminal_commands/database"
	"terminal_commands/routes"
)

func main() {

	database.InitializeDb()

	app := routes.FiberApp{App: fiber.New()}

	app.UseUserRoutes()

	app.UseCommandRoutes()

	err := app.Listen(":3000")

	if err != nil {
		return
	}

}
