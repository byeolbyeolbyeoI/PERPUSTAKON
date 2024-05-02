package handlers

import (
	"github.com/gofiber/fiber/v2"

	"perpustakaan/repository"
)

func (h *Handler) BorrowBook(c *fiber.Ctx) error {
	type borrowStruct struct {
		UserId int `json:"userId"`
		BookId int `json:"bookId"`
	}

	userRepository := repository.UserRepository{DB: h.DB}
	bookRepository := repository.BookRepository{DB: h.DB}

	var borrow borrowStruct

	if err := c.BodyParser(&borrow); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"success": false,
				"message": "Unable converting the params",
				"code": err.Error(),
			},
		)
	}

	userAvailability, APIError := userRepository.CheckUserAvailability(borrow.UserId)
	if APIError != nil {
		return c.Status(APIError.Status).JSON(
			fiber.Map{
				"success": APIError.Success,
				"message": APIError.Message,
				"code": APIError.Code,
			},
		)
	}

	if !userAvailability {
		return c.Status(fiber.StatusForbidden).JSON(
			fiber.Map{
				"success": false,
				"message": "User can only borrow one book at a time",
				"code": "USER_NOT_AVAILABLE",
			},
		)
	}

	// user is available
	bookAvailability, APIError := bookRepository.CheckBookAvailability(borrow.BookId)
	if APIError != nil {
		return c.Status(APIError.Status).JSON(
			fiber.Map{
				"success": APIError.Success,
				"message": APIError.Message,
				"code": APIError.Code,
			},
		)
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
