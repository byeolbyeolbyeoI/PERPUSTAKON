package main

import (
	"github.com/gofiber/fiber/v2"
	"perpustakaan/bootstrap"
)

// @title PERPUSTAKON
// @version 0.2 
// @host localhost:9000 
func main(){
	app := fiber.New()
	bootstrap.InitializeApp(app)
}
