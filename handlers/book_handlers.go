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

func (h *Handler) DeleteBook(c *fiber.Ctx) error {
	type bookStruct struct {
		Id int `json:"id"`
	}

	var book bookStruct

	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"error": fiber.Map{
					"message": "Error parsing the input",
					"code":    "BODYPARSER_ERROR"}})
	}

	_, err := h.DB.Exec("DELETE FROM books WHERE id=?", book.Id)	
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"error": fiber.Map{
					"message": "Error deleting book data",
					"code":    "DATABASE_ERROR"}})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Successfully delete book data"})
}

func (h *Handler) AddBook(c *fiber.Ctx) error {
	var book models.LibraryBook
	var bookRepository = repository.BookRepository{DB: h.DB}

	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"error": fiber.Map{
					"message": "Unable to parse JSON data",
					"code":    "BODYPARSER_ERROR"}})
	}

	APIError := bookRepository.AddBook(book)
	if APIError != nil {
		return c.Status(APIError.Status).JSON(
			fiber.Map{
				"error": fiber.Map{
					"message": APIError.Error.Message,
					"code":    APIError.Error.Code}})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Book Added successfully"})
}


func (h *Handler) SearchBook(c *fiber.Ctx) error {
	var title = c.Params("title")

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

	var found = false
	var index int
	
	for key, book := range books {
		if book.Book.Title == title {
			index = key
			found = true
			break // break the for if found
		}
	}

	if found == false {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"error": fiber.Map{
					"message": "Title not found",
					"code":    "BOOK_NOT_FOUND"}})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Succes retrieving book data",
		"data": fiber.Map{
			"id": books[index].Book.Id,
			"title": books[index].Book.Title,
			"author": books[index].Book.Author,
			"genres": books[index].Book.Genres,
			"synopsis": books[index].Book.Synopsis,
			"releaseYear": books[index].Book.ReleaseYear,
			"available": books[index].Available,
		},
	})
}
