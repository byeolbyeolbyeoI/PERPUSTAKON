package repository

import (
	"database/sql"
	"time"

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

func (b *BorrowRepository) BorrowBook(data models.Borrow) *APIError.APIError {
	now := time.Now().Format("2006-01-01")
	bookRepository := BookRepository{DB : b.DB}
	_, err := b.DB.Exec("INSERT INTO borrowed_books (user_id, book_id, borrowed_date) VALUES (?, ?, ?)", data.UserId, data.BookId, now)
	if err != nil {
		return APIError.NewAPIError(fiber.StatusInternalServerError, "Error borrowing book : Unable to insert data", err.Error())
	}

	err = bookRepository.ToggleBookAvailability(data.BookId)
	if err != nil {
		return APIError.NewAPIError(fiber.StatusInternalServerError, "Error borrowing book : Unable to toggle book availability", err.Error())
	}

	return nil
}

func (b *BorrowRepository) GetBookIdByUserId(userId int) (int, *APIError.APIError) {
	var bookId int
	err := b.DB.QueryRow("SELECT book_id FROM borrowed_books WHERE user_id=? AND returned_date IS NULL", userId).Scan(&bookId)
	// tidak mungkin err no rows karna udh dicek
	if err != nil {
		return -1, APIError.NewAPIError(fiber.StatusInternalServerError, "Error getting borrowed book data", err.Error())
	}

	return bookId, nil
}

func (b *BorrowRepository) ReturnBook(data models.Borrow) *APIError.APIError {
	// update returned date
	// toggle book availability
	now := time.Now().Format("2006-01-01")
	bookRepository := BookRepository{DB : b.DB}
	
	_, err := b.DB.Exec("UPDATE borrowed_books SET returned_date=? WHERE user_id=? AND returned_date IS NULL", now, data.UserId)
	if err != nil {
		return APIError.NewAPIError(fiber.StatusInternalServerError, "Error returning book : Error updating book returned date", err.Error())
	}

	err = bookRepository.ToggleBookAvailability(data.BookId)
	if err != nil {
		return APIError.NewAPIError(fiber.StatusInternalServerError, "Error returning book : Unable to update book availability", err.Error())
	}

	return nil
}
