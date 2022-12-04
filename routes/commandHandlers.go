package routes

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"terminal_commands/database"
	"terminal_commands/middlewares"
	"terminal_commands/models"
)

func GetCommands(c *fiber.Ctx) error {

	var command []models.Command

	database.AppDb.Preload("User").Find(&command)

	if command == nil {
		c.Response().SetStatusCode(fiber.StatusNoContent)
		return c.JSON([]models.Command{})
	}
	return c.JSON(command)

}

func AddCommand(c *fiber.Ctx) error {

	var command models.Command

	err := c.BodyParser(&command)

	if err != nil {
		return c.SendString(fmt.Sprintf("Error %v", err))
	}

	idString := fmt.Sprintf("%v", c.Locals("id"))

	parseUint, err := strconv.ParseUint(idString, 10, 32)

	if err != nil {
		return err
	}

	command.UserID = uint(parseUint)

	database.AppDb.Create(&command)

	return c.JSON(command)

}

func DeleteCommand(c *fiber.Ctx) error {

	id := c.Params("id")

	database.AppDb.Delete(&models.Command{}, id)

	return c.SendString("Deleted")
}

func (app *FiberApp) UseCommandRoutes() {

	commands := app.Group("/commands", middlewares.Authentication)

	commands.Get("/", GetCommands)

	commands.Post("/", AddCommand)

	commands.Delete("/:id", DeleteCommand)
}
