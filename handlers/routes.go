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

	// user
	app.Get("/getBooks", handler.GetBooks)
	app.Get("/getBookById/:id", handler.GetBook)
	app.Get("/getBookByTitle/:title", handler.SearchBook)

	// librarian
	app.Post("/addBook", middleware.OnlyLibrarian, handler.AddBook)
	app.Delete("/deleteBook", middleware.OnlyLibrarian, handler.DeleteBook)

	// users

	// user
	app.Post("/signupHandler", middleware.NotLoggedIn, handler.SignupHandler)
	app.Post("/loginHandler", middleware.NotLoggedIn, handler.LoginHandler)

	// admin
	app.Get("/getUsers", middleware.OnlyAdmin, handler.GetUsers)
	app.Get("/getUserById/:id", middleware.OnlyAdmin, handler.GetUser)

}

