package controllers

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/pamateus-henrique/infinitepay-firewatchers-api/app/models"
	"github.com/pamateus-henrique/infinitepay-firewatchers-api/pkg/utils"
	"github.com/pamateus-henrique/infinitepay-firewatchers-api/plataform/database"
)


func Test(c *fiber.Ctx) error {

	db, err := database.OpenDBConnection()

	if err != nil {
		fmt.Println(err)
	}

	user, err := db.GetUserByEmail("mateus@gmail.com")

	fmt.Println(user)

	return c.JSON(fiber.Map{
		"error": false,
		"msg": nil,
	})
}

func Register(c *fiber.Ctx) error {

	authModel := models.Register{}
	userModel := models.User{}
	validate := validator.New()

	if err := c.BodyParser(&authModel); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg": nil,
		})
	}

	if err := validate.Struct(authModel); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg": err.Error(),
		})
	}

	userModel.Name = authModel.Name
	userModel.Email = authModel.Email
	userModel.Password = utils.GeneratePassword(authModel.Password)
	userModel.Avatar_url = "abc"
	userModel.Role = "Viewer"
	userModel.Team = "None"


	if err := validate.Struct(userModel); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg": err.Error(),
		})
	}

	db, err := database.OpenDBConnection()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"message": err.Error(),
		})
	}

	if err := db.CreateUser(&userModel); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"error": false,
			"message": "User has been created",
		})
}