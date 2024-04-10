package handlers

import (
	_ "fmt"

	"github.com/gofiber/fiber/v2"

	"perpustakaan/config"
	"perpustakaan/models"
	"perpustakaan/repository"
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
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User created successfully"})
}
/*
func LoginHandler(c *fiber.Ctx) error {
	db, err := config.Connect()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error connecting to the database"})
	}

	var user models.UserInput 
	var dbUser models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	err = db.QueryRow("SELECT id, username, password, role FROM users WHERE username=?", user.Username).Scan(&dbUser.Id, &dbUser.Username, &dbUser.Password, &dbUser.Role)
	if err == sql.ErrNoRows {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Username is not registered"})
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Incorrect password"})
	}
	// generate token
	token := service.GenerateJWT(dbUser.Id, dbUser.Username, dbUser.Role)
	// sign the jwt
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
*/

func GetUsers(c *fiber.Ctx) error {
	db, err := config.Connect()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error connecting to the database"})
	}

	var users []models.User
	var userRepository = repository.UserRepository{DB: db}

	users, err = userRepository.GetAllUser()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}

	return c.Status(fiber.StatusOK).JSON(users)
}
