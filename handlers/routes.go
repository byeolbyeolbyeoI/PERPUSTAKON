package handlers 

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"

	"perpustakaan/middleware"
	"perpustakaan/config"
)

type Handler struct {
	DB *sql.DB
}

func NewHandler() (*Handler, error) {
	db, err := config.Connect()
	if err != nil {
		return nil, err
	}

	return &Handler{
		DB: db,
	}, nil
}


func SetupRoutes(app *fiber.App) {
	handler, err := NewHandler()

	app.Use(func (c *fiber.Ctx) error {
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				fiber.Map{
					"error": fiber.Map{
						"message": "Unable to parse JSON data",
						"code": "BODYPARSER_ERROR"}})	
		}

		return c.Next()
	})

	// admin
	app.Get("/users", middleware.OnlyAdmin, handler.GetUsers)

	// librarian

	// user
	app.Post("/signup", handler.SignupHandler)
	app.Post("/login", handler.LoginHandler)
}

