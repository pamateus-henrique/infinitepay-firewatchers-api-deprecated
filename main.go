package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/pamateus-henrique/infinitepay-firewatchers-api/pkg/middleware"
	"github.com/pamateus-henrique/infinitepay-firewatchers-api/pkg/routes"
	"github.com/pamateus-henrique/infinitepay-firewatchers-api/pkg/utils"
	"github.com/pamateus-henrique/infinitepay-firewatchers-api/plataform/database"
)


func main(){
	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorHandler,
	})
	db, err := database.OpenDBConnection();

	if err != nil {
		log.Fatal("Error connecting to the database, shutting down server.")
	}

	routes.PublicRoutes(app, db)

	utils.StartServer(app)
}