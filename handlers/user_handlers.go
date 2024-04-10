package handlers

import (
	"database/sql"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"

	"perpustakaan/repository"
	"perpustakaan/models"
)

func SignupHandler(c *fiber.Ctx) error {
	var user models.UserInput 
	var dbUser models.User

	err := repository.UserStore.CreateUser(user, dbUser)

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User created successfully"})
}

func LoginHandler(c *fiber.Ctx) error {
	type userStruct struct {
		Username string
		Password string
	}

	var user userStruct
	var dbUser models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	err := db.QueryRow("SELECT id, username, password, role FROM users WHERE username=?", user.Username).Scan(&dbUser.Id, &dbUser.Username, &dbUser.Password, &dbUser.Role)
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

func GetUsers(c *fiber.Ctx) error {
	rows, err := db.Query("SELECT id, username, password, role FROM users")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	defer rows.Close()

	var users []models.User
	var user models.User

	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Role); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		users = append(users, user)
	}

	return c.Status(fiber.StatusOK).JSON(users)
}

