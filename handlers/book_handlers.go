package handlers

import (
	"perpustakaan/models"
	"perpustakaan/repository"
	"database/sql"

	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) GetBooks(c *fiber.Ctx) error {
	var books []models.LibraryBook
	var bookRepository = repository.BookRepository{DB: h.DB}

	books, APIError := bookRepository.GetAllBooks()
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
			"message": "Books data retrieved successfully",
			"data":    books})
}

func (h *Handler) GetBook(c *fiber.Ctx) error {
	// this defo can be better
	var id, err = strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"success": false,
				"message": "Params invalid",
				"code":    err.Error(),
			},
		)
	}

	var book models.LibraryBook
	var bookRepository = repository.BookRepository{DB: h.DB}
	book, APIError := bookRepository.GetBookById(int(id))
	if APIError != nil {
		return c.Status(APIError.Status).JSON(
			fiber.Map{
				"success": APIError.Success,
				"message": APIError.Message,
				"code":    APIError.Code,
			},
		)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Book data retrieved successfully",
		"data": fiber.Map{
			"id":          book.Book.Id,
			"title":       book.Book.Title,
			"author":      book.Book.Author,
			"genres":      book.Book.Genres,
			"synopsis":    book.Book.Synopsis,
			"releaseYear": book.Book.ReleaseYear,
			"available":   book.Available,
		},
	})
}

func (h *Handler) DeleteBook(c *fiber.Ctx) error {
	type bookStruct struct {
		Id int `json:"bookId"`
	}

	var book bookStruct

	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"success": false,
				"message": "Input not valid",
				"code":    err.Error(),
			},
		)
	}

	var bookId int
	err := h.DB.QueryRow("SELECT id FROM books WHERE id=?", book.Id).Scan(&bookId)
	if err == sql.ErrNoRows {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"success": false,
				"message": "Book id is not registered",
				"code":    err.Error(),
			},
		)
	}

	_, err = h.DB.Exec("DELETE FROM books WHERE id=?", book.Id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"success": false,
				"message": "Error deleting book data",
				"code":    err.Error(),
			},
		)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Book deleted successfully",
	})
}

func (h *Handler) AddBook(c *fiber.Ctx) error {
	var book models.LibraryBook
	var bookRepository = repository.BookRepository{DB: h.DB}

	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"success": false,
				"message": "Input not valid",
				"code":    err.Error(),
			},
		)
	}

	APIError := bookRepository.AddBook(book)
	if APIError != nil {
		return c.Status(APIError.Status).JSON(
			fiber.Map{
				"success": APIError.Success,
				"message": APIError.Message,
				"code":    APIError.Code,
			},
		)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Book added successfully",
	})
}

func (h *Handler) SearchBook(c *fiber.Ctx) error {
	var title = c.Params("title")

	var books []models.LibraryBook
	var bookRepository = repository.BookRepository{DB: h.DB}

	books, APIError := bookRepository.GetAllBooks()
	if APIError != nil {
		return c.Status(APIError.Status).JSON(
			fiber.Map{
				"success": APIError.Success,
				"message": APIError.Message,
				"code":    APIError.Code,
			},
		)
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
				"success": false,
				"message": "Book not found",
				"code":    "Title not found",
			},
		)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Succes retrieving book data",
		"data": fiber.Map{
			"id":          books[index].Book.Id,
			"title":       books[index].Book.Title,
			"author":      books[index].Book.Author,
			"genres":      books[index].Book.Genres,
			"synopsis":    books[index].Book.Synopsis,
			"releaseYear": books[index].Book.ReleaseYear,
			"available":   books[index].Available,
		},
	})
}
