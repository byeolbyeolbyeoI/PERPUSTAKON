package models

import (
	"database/sql"
	"strings"
)

// should i do getter aswell?
type User struct {
	Id int `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role string `json:"role"`
}

type UserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Book struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Genres []string `json:"genre"`
	Synopsis string `json:"synopsis"`
	ReleaseYear int `json:"releaseYear"`
}

type LibraryBook struct {
	Book Book `json:"book"`
	Available bool `json:"available"`
}

func (lb *LibraryBook) Join(genres []string) string {
	return strings.Join(genres, ", ")
}

func (lb *LibraryBook) Split(genres string) []string {
	return strings.Split(genres, ", ")
}

func (lb *LibraryBook) CheckLibraryBookAvailability(db *sql.DB) (bool, error) {
	err := db.QueryRow("SELECT available FROM books WHERE id=?", lb.Book.Id).Scan(&lb.Available)
	if err != nil {
		return false, err
		// error handling
	}

	return lb.Available, nil
}
