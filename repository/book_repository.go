package repository

import (
	"database/sql"
	"fmt"

	"strconv"
	"github.com/gofiber/fiber/v2"

	"perpustakaan/models"
	APIError "perpustakaan/error"
)

type BookStore interface {
	GetAllBooks() ([]models.LibraryBook, *APIError.APIError)
}

type BookRepository struct {
	DB *sql.DB
}

func (b *BookRepository) GetAllBooks() ([]models.LibraryBook, *APIError.APIError) {
	rows, err := b.DB.Query("SELECT id, title, author, genres, synopsis, releaseYear, available FROM books")
	if err != nil {
		return nil, APIError.NewAPIError(fiber.StatusInternalServerError, "Error retrieving rows", err.Error())
	}
	defer rows.Close()

	var libraryBooks []models.LibraryBook
	var libraryBook models.LibraryBook
	var genresString string

	for rows.Next() {
		err := rows.Scan(
			&libraryBook.Book.Id, 
			&libraryBook.Book.Title, 
			&libraryBook.Book.Author, 
			&genresString, 
			&libraryBook.Book.Synopsis, 
			&libraryBook.Book.ReleaseYear, 
			&libraryBook.Available)
		if err != nil {
			return nil, APIError.NewAPIError(fiber.StatusInternalServerError, "Error scanning rows", err.Error())
		}

		libraryBook.Book.Genres = libraryBook.Split(genresString)

		libraryBooks = append(libraryBooks, libraryBook)
	}

	return libraryBooks, nil
}

func (b *BookRepository) GetBookById(id int) (models.LibraryBook, *APIError.APIError) {
	var dbBook models.LibraryBook
	var genresString string
	err := b.DB.QueryRow("SELECT id, title, author, genres, synopsis, releaseYear, available FROM books WHERE id=?", id).Scan(
			&dbBook.Book.Id, 
			&dbBook.Book.Title, 
			&dbBook.Book.Author, 
			&genresString, 
			&dbBook.Book.Synopsis, 
			&dbBook.Book.ReleaseYear, 
			&dbBook.Available)
	if err == sql.ErrNoRows {
		return dbBook, APIError.NewAPIError(fiber.StatusInternalServerError, "Id is not registered", err.Error())
	}

	dbBook.Book.Genres = dbBook.Split(genresString)

	return dbBook, nil
}

func (b *BookRepository) AddBook(book models.LibraryBook) *APIError.APIError {
	var genresString string
	var bookAvailability string
	bookAvailability = strconv.FormatBool(book.Available)
	genresString = book.Join(book.Book.Genres)
	_, err := b.DB.Exec("INSERT INTO books (title, author, genres, synopsis, releaseYear, available) VALUES (?, ?, ?, ?, ?, ?)", 
			&book.Book.Title, 
			&book.Book.Author, 
			genresString, 
			&book.Book.Synopsis, 
			&book.Book.ReleaseYear, 
			bookAvailability,
	)
	if err != nil {
		return APIError.NewAPIError(fiber.StatusInternalServerError, "Error adding book data", err.Error())
	}

	return nil
}

func (b *BookRepository) CheckBookAvailability(id int) (bool, *APIError.APIError) {
	var availability string
	err := b.DB.QueryRow("SELECT available FROM books WHERE id=?", id).Scan(&availability)
	if err == sql.ErrNoRows {
		return false, APIError.NewAPIError(fiber.StatusInternalServerError, "Book is not registered", err.Error())
	}
	if err != nil {
		return false, APIError.NewAPIError(fiber.StatusInternalServerError, "Error scanning rows", err.Error())
	}

	if availability == "false" {
		return false, nil
	}

	return true, nil
}

func (b *BookRepository) ToggleBookAvailability(id int) error {
	var availability bool
	err := b.DB.QueryRow("SELECT available FROM books WHERE id=?", id).Scan(&availability)
	if err == sql.ErrNoRows {
		return err
	}
	if err != nil {
		return err
	}

	availability = !availability

	_, err = b.DB.Exec("UPDATE books SET available=? WHERE id=?", strconv.FormatBool(availability), id)
	if err != nil {
		return err
	}
	fmt.Println(availability)

	return nil
}
