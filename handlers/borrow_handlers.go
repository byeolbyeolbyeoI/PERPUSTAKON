package handlers

import (
	"github.com/gofiber/fiber/v2"

	"perpustakaan/models"
	"perpustakaan/repository"
)

func (h *Handler) BorrowBook(c *fiber.Ctx) error {
	var data models.Borrow
	userRepository := repository.UserRepository{DB: h.DB}
	bookRepository := repository.BookRepository{DB: h.DB}
	borrowRepository := repository.BorrowRepository{DB: h.DB}

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"success": false,
				"message": "Unable converting the params",
				"code":    err.Error(),
			},
		)
	}

	userAvailability, APIError := userRepository.CheckUserAvailability(data.UserId)
	if APIError != nil {
		return c.Status(APIError.Status).JSON(
			fiber.Map{
				"success": APIError.Success,
				"message": APIError.Message,
				"code":    APIError.Code,
			},
		)
	}

	if !userAvailability {
		return c.Status(fiber.StatusForbidden).JSON(
			fiber.Map{
				"success": false,
				"message": "User can only borrow one book at a time",
				"code":    "USER_NOT_AVAILABLE",
			},
		)
	}

	bookAvailability, APIError := bookRepository.CheckBookAvailability(data.BookId)
	if APIError != nil {
		return c.Status(APIError.Status).JSON(
			fiber.Map{
				"success": APIError.Success,
				"message": APIError.Message,
				"code":    APIError.Code,
			},
		)
	}

	if !bookAvailability {
		return c.Status(fiber.StatusForbidden).JSON(
			fiber.Map{
				"success": false,
				"message": "Book is being borrowed",
				"code":    "BOOK_NOT_AVAILABLE",
			},
		)
	}

	APIError = borrowRepository.BorrowBook(data)
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
			"message": "Successfully borrowed the book",
		},
	)
}

func (h *Handler) ReturnBook(c *fiber.Ctx) error {
	type userData struct {
		UserId int `json:"userId"`
	}

	var inputData userData
	var data models.Borrow
	userRepository := repository.UserRepository{DB: h.DB}
	borrowRepository := repository.BorrowRepository{DB: h.DB}

	if err := c.BodyParser(&inputData); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"success": false, "message": "Unable converting the params",
				"code":    err.Error(),
			},
		)
	}

	userAvailability, APIError := userRepository.CheckUserAvailability(inputData.UserId)
	if APIError != nil {
		return c.Status(APIError.Status).JSON(
			fiber.Map{
				"success": APIError.Success,
				"message": APIError.Message,
				"code":    APIError.Code,
			},
		)
	}

	if userAvailability {
		return c.Status(fiber.StatusForbidden).JSON(
			fiber.Map{
				"success": false,
				"message": "User is not borrowing any book",
				"code":    "USER_NOT_AVAILABLE",
			},
		)
	}

	data.UserId = inputData.UserId
	data.BookId, APIError = borrowRepository.GetBookIdByUserId(inputData.UserId)
	if APIError != nil {
		return c.Status(APIError.Status).JSON(
			fiber.Map{
				"success": APIError.Success,
				"message": APIError.Message,
				"code":    APIError.Code,
			},
		)
	}

	APIError = borrowRepository.ReturnBook(data)
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
			"message": "Successfully returned the book",
		},
	)
}
