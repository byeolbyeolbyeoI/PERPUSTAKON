package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"perpustakaan/handler"
)

func main(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
	app := fiber.New()

	handler.AddRoutes(app)	

	log.Fatal(app.Listen(":9000"))
}
