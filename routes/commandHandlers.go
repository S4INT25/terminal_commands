package routes

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"terminal_commands/database"
	"terminal_commands/middlewares"
	"terminal_commands/models"
	"terminal_commands/utils"
)

func GetCommands(c *fiber.Ctx) error {

	var commands []models.Command

	database.AppDb.Preload("User").Find(&commands)

	if commands == nil {
		c.Response().SetStatusCode(fiber.StatusNoContent)
		return c.JSON([]models.Command{})
	}

	response := utils.Map(commands, func(command models.Command) models.CommandResponse { return command.ToResponse() })

	return c.JSON(response)

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

	return c.JSON(models.CommandResponse{
		Id:          command.ID,
		Name:        command.Name,
		Description: command.Description,
		Platform:    command.Platform.Name(),
		User: models.UserResponse{
			Id:    command.User.ID,
			Name:  command.User.Name,
			Email: command.User.Email,
		},
	})

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
