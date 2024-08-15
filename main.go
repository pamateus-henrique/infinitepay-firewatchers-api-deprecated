package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pamateus-henrique/infinitepay-firewatchers-api/pkg/routes"
	"github.com/pamateus-henrique/infinitepay-firewatchers-api/pkg/utils"
)


func main(){
	app := fiber.New()

	
	routes.PublicRoutes(app)

	utils.StartServer(app)
}