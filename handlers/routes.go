package handlers 

import (
	"github.com/gofiber/fiber/v2"

	"perpustakaan/middleware"
)

func SetupRoutes(app *fiber.App) {
	// admin
	app.Get("/users", middleware.OnlyAdmin, GetUsers)

	// librarian

	// user
	app.Post("/signup", SignupHandler)
	app.Post("/login", LoginHandler)
}

