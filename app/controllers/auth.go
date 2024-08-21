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
		return fiber.NewError(fiber.StatusBadRequest)
	}

	if err := ac.UserService.ValidateRegisterModel(&authModel); err != nil {
		return fiber.NewError(fiber.StatusBadRequest)
	}

	userModel := ac.UserService.CreateUserModel(&authModel)	

	if err := ac.UserService.ValidateUserModel(userModel); err != nil {
		return fiber.NewError(fiber.StatusBadRequest)
	}

	if err := ac.UserService.CreateUser(userModel); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"error": false,
			"message": "User has been created",
		})
}


func (ac *AuthController) Login(c *fiber.Ctx) error {
	
	authModel := models.Login{}
	
	if err := c.BodyParser(&authModel); err != nil {
		return fiber.NewError(fiber.StatusBadRequest)
	}

	if err := ac.UserService.ValidateLoginModel(&authModel); err != nil {
		return fiber.NewError(fiber.StatusBadRequest)
	}

	user, err := ac.UserService.GetUserByEmail(authModel.Email)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "No user found with this email/password")
	}

	if err := utils.ComparePassword(authModel.Password, user.Password); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "No user found with this email/password")
	}

	 jwt, err := utils.GenerateJWT(user.Name)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	 // Create cookie
	 cookie := new(fiber.Cookie)
	 cookie.Name = "auth"
	 cookie.Value = jwt
	 cookie.HTTPOnly = true
	 c.Cookie(cookie)

	 return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg": "Login successful",
	})
}