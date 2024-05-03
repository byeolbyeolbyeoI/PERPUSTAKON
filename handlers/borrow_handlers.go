package handlers

import (
	"github.com/gofiber/fiber/v2"
	"time"
	"fmt"

	"perpustakaan/repository"
	"perpustakaan/models"
)

func (h *Handler) BorrowBook(c *fiber.Ctx) error {
	var borrow models.Borrow 
	userRepository := repository.UserRepository{DB: h.DB}
	bookRepository := repository.BookRepository{DB: h.DB}
	borrowRepository := repository.BorrowRepository{DB: h.DB}

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

	if !bookAvailability {
		return c.Status(fiber.StatusForbidden).JSON(
			fiber.Map{
				"success": false,
				"message": "Boos is being borrowed",
				"code": "BOOK_NOT_AVAILABLE",
			},
		)
	}

	now := time.Now().Format("2006-12-21")
	fmt.Println("sekarang : ", now)
	APIError = borrowRepository.BorrowBook(borrow, now)
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

	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"success": true,
			"message": "Successfully borrowed the book",
		})
}

func (h *Handler) ReturnBook(c *fiber.Ctx) error {
	return nil
}
