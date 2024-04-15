package handlers

import (
	"github.com/gofiber/fiber/v2"

	"perpustakaan/middleware"
	"perpustakaan/models"
	"perpustakaan/repository"
	"perpustakaan/service"
)

func (h *Handler) SignupHandler(c *fiber.Ctx) error {
	var user models.UserInput
	var userRepository = repository.UserRepository{DB: h.DB}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"error": fiber.Map{
					"message": "Unable to parse JSON data",
					"code": "BODYPARSER_ERROR"}})
	}

	APIError := userRepository.CreateUser(user)
	if APIError != nil {
		return c.Status(APIError.Status).JSON(
			fiber.Map{
				"error": fiber.Map{
					"message": APIError.Error.Message,
					"code": APIError.Error.Code}})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User created successfully"})
}

func (h *Handler) LoginHandler(c *fiber.Ctx) error {
	if  _, loggedIn := middleware.IsLoggedIn(c); loggedIn {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "You are already logged in"})
	}

	var user models.UserInput
	var dbUser models.User
	var userRepository = repository.UserRepository{DB: h.DB}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	dbUser, APIError := userRepository.GetUserByUsername(user)	
	if APIError != nil {
		return c.Status(APIError.Status).JSON(
			fiber.Map{
				"error": fiber.Map{
					"message": APIError.Error.Message,
					"code": APIError.Error.Code}})
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

	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"data": fiber.Map{
				"jwt": tokenString,
				"id": dbUser.Id,
				"role": dbUser.Role,
				"username": dbUser.Username}})
}

func (h *Handler) GetUsers(c *fiber.Ctx) error {
	var users []models.User
	var userRepository = repository.UserRepository{DB: h.DB}

	users, APIError := userRepository.GetAllUser()
	if APIError != nil {
		// return error that suits the error inside of GetAllUser() function since it vary in errors (im the goat)
		return c.Status(APIError.Status).JSON(
			fiber.Map{
				"error": fiber.Map{
					"message": APIError.Error.Message,
					"code": APIError.Error.Code}})
	}

	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"message": "Successfully retrieved users data",
			"data": users})
}

/*
this code is ass
func UpdateUser(c *fiber.Ctx) error {
	var user models.User

	tokenString, ok := middleware.IsLoggedIn(c)
	fmt.Println("token string :", tokenString)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(
			fiber.Map{
				"error": fiber.Map{
					"message": "User not logged in",
					"code": "AUTHENTICATION_ERROR"}})
	}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"error": fiber.Map{
					"message": "Unable to process the input",
					"code": "BODYPARSER_ERROR"}})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Update user data successfully"})
}
*/
