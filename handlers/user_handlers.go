package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"

	"perpustakaan/config"
	"perpustakaan/models"
	"perpustakaan/repository"
	"perpustakaan/service"
)

func SignupHandler(c *fiber.Ctx) error {
	db, err := config.Connect()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error connecting to the database"})
	}

	var user models.UserInput
	var dbUser models.User
	var userRepository = repository.UserRepository{DB: db}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	err = userRepository.CreateUser(user, dbUser)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User created successfully"})
}

func LoginHandler(c *fiber.Ctx) error {
	db, err := config.Connect()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error connecting to the database"})
	}

	var user models.UserInput
	var dbUser *models.User
	var userRepository = repository.UserRepository{DB: db}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	dbUser, err = userRepository.GetUserById(user, dbUser)	
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":err.Error()})
	}

	fmt.Println(dbUser.Username)
	token := service.GenerateJWT(dbUser.Id, dbUser.Username, dbUser.Role)
	tokenString, err := service.SignToken(token)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Unable to sign the token"})
	}

	// set the cookie
	cookie := &fiber.Cookie{
		Name:     "token",
		Value:    tokenString,
		MaxAge:   3600 * 24 * 30,
		HTTPOnly: true,
		SameSite: "lax",
	}
	c.Cookie(cookie)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User logged in succesfully"})
}

func GetUsers(c *fiber.Ctx) error {
	db, err := config.Connect()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error connecting to the database"})
	}

	var users []models.User
	var userRepository = repository.UserRepository{DB: db}

	users, err = userRepository.GetAllUser()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(users)
}
