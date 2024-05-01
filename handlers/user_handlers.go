package handlers

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
	"os"
	"fmt"
	"time"

	"perpustakaan/models"
	"perpustakaan/repository"
	"perpustakaan/service"
	_ "perpustakaan/docs"
)

// @Summary Signup a new user
// @Description Allows users to create their acccount
// @Tags user 
// @RequestBody Required
// @Accept json
// @Produce json 
func (h *Handler) SignupHandler(c *fiber.Ctx) error {
	var user models.UserInput
	var userRepository = repository.UserRepository{DB: h.DB}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"error": fiber.Map{
					"message": "Unable to parse JSON data",
					"code":    "BODYPARSER_ERROR"}})
	}

	APIError := userRepository.CreateUser(user)
	if APIError != nil {
		return c.Status(APIError.Status).JSON(
			fiber.Map{
				"error": fiber.Map{
					"message": APIError.Error.Message,
					"code":    APIError.Error.Code}})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User created successfully"})
}

func (h *Handler) LoginHandler(c *fiber.Ctx) error {
	var user models.UserInput
	var dbUser models.User
	var userRepository = repository.UserRepository{DB: h.DB}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	dbUser, APIError := userRepository.GetUserByUsername(user.Username)
	if APIError != nil {
		return c.Status(APIError.Status).JSON(
			fiber.Map{
				"error": fiber.Map{
					"message": APIError.Error.Message,
					"code":    APIError.Error.Code}})
	}

	APIError = userRepository.CheckPassword(user, dbUser)
	if APIError != nil {
		return c.Status(APIError.Status).JSON(
			fiber.Map{
				"error": fiber.Map{
					"message": APIError.Error.Message,
					"code":    APIError.Error.Code}})
	}

	token := service.GenerateJWT(dbUser)
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

	// do the logging with file.txt
	data, err := os.ReadFile("logging.txt")
	if err != nil {
		// file doesnt exists
		log := fmt.Sprintln(dbUser.Role, dbUser.Username, "logged in at", time.Now())		
		err := os.WriteFile("logging.txt", []byte(log), 0644)		
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Unable writing to the file"})
		}
	} else {
		log := string(data)
		log += fmt.Sprintln(dbUser.Role, dbUser.Username, "logged in at", time.Now())		
		err := os.WriteFile("logging.txt", []byte(log), 0644)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Unable writing to the file"})
		}
	}

	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"data": fiber.Map{
				"jwt":      tokenString,
				"id":       dbUser.Id,
				"role":     dbUser.Role,
				"username": dbUser.Username}})
}

func (h *Handler) GetUsers(c *fiber.Ctx) error {
	var users []models.User
	var userRepository = repository.UserRepository{DB: h.DB}

	users, APIError := userRepository.GetAllUsers()
	if APIError != nil {
		// reminder that im the goat
		return c.Status(APIError.Status).JSON(
			fiber.Map{
				"error": fiber.Map{
					"message": APIError.Error.Message,
					"code":    APIError.Error.Code}})
	}

	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"message": "Successfully retrieved users data",
			"data":    users})
}

func (h *Handler) GetUser(c *fiber.Ctx) error {
	// this defo can be better
	var id, err = strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"error": fiber.Map{
					"message": "Unable to convert string to integer",
					"code":    "STRCONV_ERROR"}})
	}

	var user models.User
	var userRepository = repository.UserRepository{DB: h.DB}
	user, APIError := userRepository.GetUserById(int(id))
	if APIError != nil {
		return c.Status(APIError.Status).JSON(
			fiber.Map{
				"error": fiber.Map{
					"message": APIError.Error.Message,
					"code":    APIError.Error.Code}})
	}

	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"data": fiber.Map{
				"id":       user.Id,
				"username": user.Username,
				"role":     user.Role}})
}
