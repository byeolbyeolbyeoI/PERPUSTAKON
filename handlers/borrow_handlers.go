package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) BorrowBook(c *fiber.Ctx) error {
	type borrowStruct struct {
		UserId int `json:"userId"`
		BookId int `json:"bookId"`
	}

	var borrow borrowStruct

	if err := c.BodyParser(&borrow); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"error": fiber.Map{
					"message": "Unable to parse JSON data",
					"code":    "BODYPARSER_ERROR"}})
	}

	// set time format same as mysql
	// check if user is borrowing (check if they have returned_date null)
	// check if book available
	// insert into borrow_book table
	// toggle book availability

	/*
	APIError := userRepository.CreateUser(user)
	if APIError != nil {
		return c.Status(APIError.Status).JSON(
			fiber.Map{
				"error": fiber.Map{
					"message": APIError.Error.Message,
					"code":    APIError.Error.Code}})
	}
	*/

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User created successfully"})
}

func (h *Handler) ReturnBook(c *fiber.Ctx) error {
	return nil
}
