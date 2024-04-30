package handlers

import (
	"perpustakaan/models"
	"perpustakaan/repository"

	"strconv"

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

func (h *Handler) GetBook(c *fiber.Ctx) error {
	// this defo can be better
	var id, err = strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"error": fiber.Map{
					"message": "Unable to convert string to integer",
					"code":    "STRCONV_ERROR"}})
	}

	var book models.LibraryBook
	var bookRepository = repository.BookRepository{DB: h.DB} 
	book, APIError := bookRepository.GetBookById(int(id))
	if APIError != nil {
		return c.Status(APIError.Status).JSON(
			fiber.Map{
				"error": fiber.Map{
					"message": APIError.Error.Message,
					"code":    APIError.Error.Code}})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Succes retrieving book data",
		"data": fiber.Map{
			"id": book.Book.Id,
			"title": book.Book.Title,
			"author": book.Book.Author,
			"genres": book.Book.Genres,
			"synopsis": book.Book.Synopsis,
			"releaseYear": book.Book.ReleaseYear,
			"available": book.Available,
		},
	})
}
