package repository

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"

	"perpustakaan/models"
	APIError "perpustakaan/error"
)

type BorrowStore interface {
	BorrowBook(data models.Borrow, now string) *APIError.APIError 
}

type BorrowRepository struct {
	DB *sql.DB
}

func (b *BorrowRepository) BorrowBook(data models.Borrow, now string) *APIError.APIError {
	bookRepository := BookRepository{DB : b.DB}
	_, err := b.DB.Exec("INSERT INTO borrowed_books (user_id, book_id, borrowed_date)", data.UserId, data.BookId, now)
	if err != nil {
		return APIError.NewAPIError(fiber.StatusInternalServerError, "Error borrowing books", err.Error())
	}

	err = bookRepository.ToggleBookAvailability(data.BookId)
	if err != nil {
		return APIError.NewAPIError(fiber.StatusInternalServerError, "Error borrowing books", err.Error())
	}

	return nil
}
