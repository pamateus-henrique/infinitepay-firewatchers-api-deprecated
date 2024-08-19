package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pamateus-henrique/infinitepay-firewatchers-api/app/models"
	"github.com/pamateus-henrique/infinitepay-firewatchers-api/app/services"
	"github.com/pamateus-henrique/infinitepay-firewatchers-api/pkg/utils"
)


type AuthController struct {
	UserService *services.UserService
}


func NewAuthController (UserService *services.UserService) *AuthController {
	return &AuthController{
		UserService: UserService,
	}
}

func (ac *AuthController) Register(c *fiber.Ctx) error {

	authModel := models.Register{}

	if err := c.BodyParser(&authModel); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg": nil,
		})
	}

	if err := ac.UserService.ValidateRegisterModel(&authModel); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg": err.Error(),
		})
	}

	userModel := ac.UserService.CreateUserModel(&authModel)	

	if err := ac.UserService.ValidateUserModel(userModel); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg": err.Error(),
		})
	}

	if err := ac.UserService.CreateUser(userModel); err != nil {
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


func (ac *AuthController) Login(c *fiber.Ctx) error {
	
	authModel := models.Login{}
	
	if err := c.BodyParser(&authModel); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg": nil,
		})
	}

	if err := ac.UserService.ValidateLoginModel(&authModel); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg": nil,
		})
	}

	user, err := ac.UserService.GetUserByEmail(authModel.Email)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg": "No user with this email/password",
		})
	}

	if err := utils.ComparePassword(authModel.Password, user.Password); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg": "No user with this email/password",
		})
	}

	





	return nil
}