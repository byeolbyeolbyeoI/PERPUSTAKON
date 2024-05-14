package repository

import (
	"database/sql"
	"strconv"
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"

	APIError "perpustakaan/error"
	"perpustakaan/models"
)

type BorrowStore interface {
	BorrowBook(data models.Borrow, now string) *APIError.APIError 
}

type BorrowRepository struct {
	DB *sql.DB
}

func (b *BorrowRepository) BorrowBook(data models.Borrow) *APIError.APIError {
	now := time.Now().Format("2006-01-02")
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

func (b *BorrowRepository) ReturnBook(data models.Borrow) (int, *APIError.APIError) {
	// update returned date
	// toggle book availability
	format := "2006-01-02"
	now := time.Now().Format(format)
	bookRepository := BookRepository{DB : b.DB}

	var borrowedDate string
	err := b.DB.QueryRow("SELECT borrowed_date FROM borrowed_books WHERE user_id=? AND returned_date IS NULL", data.UserId).Scan(&borrowedDate)
	if err != nil {
		return 0, APIError.NewAPIError(fiber.StatusInternalServerError, "Error returning book : Unable to retrieve borrowed date data", err.Error())
	}
	_, err = b.DB.Exec("UPDATE borrowed_books SET returned_date=? WHERE user_id=? AND returned_date IS NULL", now, data.UserId)
	if err != nil {
		return 0, APIError.NewAPIError(fiber.StatusInternalServerError, "Error returning book : Error updating book returned date", err.Error())
	}

	err = bookRepository.ToggleBookAvailability(data.BookId)
	if err != nil {
		return 0, APIError.NewAPIError(fiber.StatusInternalServerError, "Error returning book : Unable to update book availability", err.Error())
	}

	parsedBorrowedDate, err := time.Parse(format, borrowedDate)
	if err != nil {
		return 0, APIError.NewAPIError(fiber.StatusInternalServerError, "Error returning book :  Error parsing borrowed date", err.Error())
	}

	// tanggal pinjam + seminggu, telat
	lateReturnDate := parsedBorrowedDate.AddDate(0, 0, 7)
	formattedLateReturnDate := lateReturnDate.Format(format)

	lateReturnDateYear, err := strconv.Atoi(strings.TrimSpace(formattedLateReturnDate[0:4]))
	if err != nil {
		return 0, APIError.NewAPIError(fiber.StatusInternalServerError, "Error returning book :  Error converting data", err.Error())
	}
	lateReturnDateMonth,  err := strconv.Atoi(strings.TrimSpace(formattedLateReturnDate[5:7]))
	if err != nil {
		return 0, APIError.NewAPIError(fiber.StatusInternalServerError, "Error returning book :  Error converting data", err.Error())
	}
	lateReturnDateDay, err := strconv.Atoi(strings.TrimSpace(formattedLateReturnDate[8:10]))
	if err != nil {
		return 0, APIError.NewAPIError(fiber.StatusInternalServerError, "Error returning book :  Error converting data", err.Error())
	}

	returnedDateYear, err := strconv.Atoi(strings.TrimSpace(now[0:4]))
	if err != nil {
		return 0, APIError.NewAPIError(fiber.StatusInternalServerError, "Error returning book :  Error converting data", err.Error())
	}
	returnedDateMonth,  err := strconv.Atoi(strings.TrimSpace(now[5:7]))
	if err != nil {
		return 0, APIError.NewAPIError(fiber.StatusInternalServerError, "Error returning book :  Error converting data", err.Error())
	}
	returnedDateDay, err := strconv.Atoi(strings.TrimSpace(now[8:10]))
	if err != nil {
		return 0, APIError.NewAPIError(fiber.StatusInternalServerError, "Error returning book :  Error converting data", err.Error())
	}

	lateReturnDateTotalDays := lateReturnDateYear * 365 + lateReturnDateMonth * 30 + lateReturnDateDay
	returnedDateTotalDays := returnedDateYear * 365 +  returnedDateMonth * 30 + returnedDateDay

	if returnedDateTotalDays > lateReturnDateTotalDays {
		lateTotal := returnedDateTotalDays - lateReturnDateTotalDays
		fmt.Println(lateTotal)
		return lateTotal * 10000, nil
	}

	return 0, nil 
}
