package main

import (
	"github.com/gofiber/fiber/v2"
	"perpustakaan/bootstrap"
)

func main(){
	app := fiber.New()
	bootstrap.InitializeApp(app)
}
