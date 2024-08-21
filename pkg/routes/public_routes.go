package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/pamateus-henrique/infinitepay-firewatchers-api/app/controllers"
	"github.com/pamateus-henrique/infinitepay-firewatchers-api/app/queries"
	"github.com/pamateus-henrique/infinitepay-firewatchers-api/app/services"
)


func PublicRoutes(a *fiber.App, db *sqlx.DB){
	route := a.Group("/api/v1")


	//auth route
	userQueries := queries.CreateUserQueries(db);
	authService := services.CreateUserService(userQueries);
	authController := controllers.NewAuthController(authService)

	route.Post("/register", authController.Register)
	route.Post("/login", authController.Login)
}