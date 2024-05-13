package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os"
	"strconv"
	"time"

	"perpustakaan/models"
	"perpustakaan/repository"
	"perpustakaan/service"
)

// @Summary Signup a new user
// @Description Allows users to create their acccount
// @Tags user
// @RequestBody Required
// @Accept json
// @Produce json

func (h *Handler) AddUser(c *fiber.Ctx) error {
	var user models.User
	var userRepository = repository.UserRepository{DB: h.DB}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"success": false,
				"message": "Error parsing the body",
				"code":    err.Error(),
			},
		)
	}

	APIError := userRepository.AddUser(user)
	if APIError != nil {
		return c.Status(APIError.Status).JSON(
			fiber.Map{
				"success": APIError.Success,
				"message": APIError.Message,
				"code":    APIError.Code,
			},
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"success": true,
			"message": "Successfully add user",
		})
}

func (h *Handler) DeleteUser(c *fiber.Ctx) error {
	type userStruct struct {
		Id int `json:"userId"`
	}

	var user userStruct

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"success": false,
				"message": "Input not valid",
				"code":    err.Error(),
			},
		)
	}

	_, err := h.DB.Exec("DELETE FROM users WHERE id=?", user.Id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"success": false,
				"message": "Error deleting user",
				"code":    err.Error(),
			},
		)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "User deleted successfully",
	})
}
func (h *Handler) SignupHandler(c *fiber.Ctx) error {
	var user models.UserInput
	var userRepository = repository.UserRepository{DB: h.DB}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"success": false,
				"message": "Error parsing the body",
				"code":    err.Error(),
			},
		)
	}

	APIError := userRepository.CreateUser(user)
	if APIError != nil {
		return c.Status(APIError.Status).JSON(
			fiber.Map{
				"success": APIError.Success,
				"message": APIError.Message,
				"code":    APIError.Code,
			},
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"success": true,
			"message": "Successfully signed up",
		})
}

func (h *Handler) LoginHandler(c *fiber.Ctx) error {
	var user models.UserInput
	var dbUser models.User
	var userRepository = repository.UserRepository{DB: h.DB}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"success": false,
				"message": "Error parsing the body",
				"code":    err.Error(),
			},
		)
	}

	dbUser, APIError := userRepository.GetUserByUsername(user.Username)
	if APIError != nil {
		return c.Status(APIError.Status).JSON(
			fiber.Map{
				"success": APIError.Success,
				"message": APIError.Message,
				"code":    APIError.Code,
			},
		)
	}

	APIError = userRepository.CheckPassword(user, dbUser)
	if APIError != nil {
		return c.Status(APIError.Status).JSON(
			fiber.Map{
				"success": APIError.Success,
				"message": APIError.Message,
				"code":    APIError.Code,
			},
		)
	}

	token := service.GenerateJWT(dbUser)
	tokenString, err := service.SignToken(token)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"success": false,
				"message": "Unable to sign the token",
				"code":    err.Error(),
			},
		)
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
			fmt.Println("error:", err.Error())
		}
	} else {
		log := string(data)
		log += fmt.Sprintln(dbUser.Role, dbUser.Username, "logged in at", time.Now())
		err := os.WriteFile("logging.txt", []byte(log), 0644)
		if err != nil {
			fmt.Println("error", err.Error())
		}
	}

	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"success": true,
			"message": "Successfully logged in",
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
		return c.Status(APIError.Status).JSON(
			fiber.Map{
				"success": APIError.Success,
				"message": APIError.Message,
				"code":    APIError.Code,
			},
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"success": true,
			"message": "Successfully retrieved users data",
			"data":    users})
}

func (h *Handler) GetUser(c *fiber.Ctx) error {
	// this defo can be better
	var id, err = strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"success": false,
				"message": "Unable converting the params",
				"code":    err.Error(),
			},
		)
	}

	var user models.User
	var userRepository = repository.UserRepository{DB: h.DB}
	user, APIError := userRepository.GetUserById(int(id))
	if APIError != nil {
		return c.Status(APIError.Status).JSON(
			fiber.Map{
				"success": APIError.Success,
				"message": APIError.Message,
				"code":    APIError.Code,
			},
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"success": true,
			"message": "Successfully retrieved user data",
			"data": fiber.Map{
				"id":       user.Id,
				"username": user.Username,
				"role":     user.Role,
			},
		})
}
