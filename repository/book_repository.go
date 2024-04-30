package repository

import (
	"database/sql"

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
