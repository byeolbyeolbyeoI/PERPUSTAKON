package bootstrap

import (
	"log"
	"os"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"perpustakaan/config"
	"perpustakaan/repository"
)

func InitializeApp(app *fiber.App) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	DBConfig := &config.Config{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Protocol: os.Getenv("DB_PROTOCOL"),
		Path:     os.Getenv("DB_PATH"),
		DBName:   os.Getenv("DB_DBNAME"),
	}

	db, err := config.Connect(DBConfig)
	if err != nil {
		log.Fatal("could not load the database")
	}

	repo := repository.Repository{
		DB: db,
	}

	repo.SetupRoutes(app)
	log.Fatal(app.Listen(":9000"))
}
