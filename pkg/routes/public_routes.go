package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pamateus-henrique/infinitepay-firewatchers-api/app/controllers"
)


func PublicRoutes(a *fiber.App){
	route := a.Group("/api/v1")

	route.Post("/register", controllers.Register)
}