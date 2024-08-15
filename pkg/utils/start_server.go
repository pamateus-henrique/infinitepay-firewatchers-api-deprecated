package utils

import (
	"log"

	"github.com/gofiber/fiber/v2"
)



func StartServer(a *fiber.App){

	if err := a.Listen(":3000"); err != nil {
		log.Fatalf("Oops... Server is not running! Reason: %v", err)
	}

}	