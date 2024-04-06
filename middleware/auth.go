package middleware

import (
	"github.com/gofiber/fiber/v2"
	_ "perpustakaan/models"
)

func OnlyLibrarian(fn fiber.Handler) fiber.Handler { 
	return func(c *fiber.Ctx) error {
		// pass user

		return fn(c)
	}
}

func OnlyAdmin(fn fiber.Handler) fiber.Handler { // doesnt matter if it is passed by value cz we're not going to change anything
	return func(c *fiber.Ctx) error {
		// pass user

		return fn(c)
	}
}
