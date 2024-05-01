package handlers 

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"

	"perpustakaan/middleware"
	"perpustakaan/config"
	_ "perpustakaan/docs"
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

	// reminder that im the goat
	app.Use(func (c *fiber.Ctx) error {
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				fiber.Map{
					"error": fiber.Map{
						"message": "Unable to connect to the database",
						"code": "DATABASE_ERROR"}})	
		}

		return c.Next()
	})

	//swagger 
	app.Get("/swagger/*", swagger.HandlerDefault)

	// books
	app.Get("/getBooks", handler.GetBooks)
	app.Get("/getBook/:id", handler.GetBook)
	app.Get("/searchBook/:title", handler.SearchBook)
	app.Post("/addBook", handler.AddBook)
	app.Delete("/deleteBook", handler.DeleteBook)

	// user
	app.Get("/getUsers", middleware.OnlyAdmin, handler.GetUsers)
	app.Get("/getUsers/:id", middleware.OnlyAdmin, handler.GetUser)

	//not now
	app.Post("/signupHandler", middleware.NotLoggedIn, handler.SignupHandler)
	app.Post("/loginHandler", middleware.NotLoggedIn, handler.LoginHandler)
}

