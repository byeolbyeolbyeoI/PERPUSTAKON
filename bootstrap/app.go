package bootstrap

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"perpustakaan/handlers"
)

func InitializeApp(app *fiber.App) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	handlers.SetupRoutes(app)
	log.Fatal(app.Listen(":9000"))
}
