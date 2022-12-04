package routes

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"terminal_commands/database"
	"terminal_commands/models"
	"terminal_commands/utils"
)

func CreateUser(c *fiber.Ctx) error {

	var body = struct {
		Email    string
		Name     string
		Password string
	}{}

	err := c.BodyParser(&body)
	if err != nil {
		return c.SendString(fmt.Sprintf("Failed to user  %v", err))
	}

	user := models.User{
		Email:        body.Email,
		PassWordHash: hashPassword(body.Password),
		Name:         body.Name,
	}

	database.AppDb.Create(&user)

	return c.JSON(struct {
		Id    uint
		Name  string
		Email string
	}{
		user.ID,
		user.Name,
		user.Email,
	})

}

func Login(c *fiber.Ctx) error {

	var body = struct {
		Email    string
		PassWord string
	}{}

	err := c.BodyParser(&body)
	if err != nil {
		return c.SendString("Failed to parse request body")
	}

	var user models.User

	database.AppDb.Find(&user, "email = ?", body.Email)

	if bcrypt.CompareHashAndPassword([]byte(user.PassWordHash), []byte(body.PassWord)) == nil {

		return c.JSON(utils.GenerateJwtToken(user))

	} else {

		return c.SendString("Invalid password")
	}

}

func hashPassword(password string) string {

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(fmt.Sprintf("failed to hash password %v", err))
	}
	return string(bytes)
}

func (app *FiberApp) UseUserRoutes() {

	app.Post("/users", CreateUser)

	app.Post("/login", Login)

}
