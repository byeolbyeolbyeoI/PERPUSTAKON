package handlers

import (
	"perpustakaan/models"
	"perpustakaan/repository"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) GetBooks(c *fiber.Ctx) error {
	var books []models.LibraryBook
	var bookRepository = repository.BookRepository{DB: h.DB}

	books, APIError := bookRepository.GetAllBooks()
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
			"message": "Successfully retrieved books data",
			"data":    books})
}
