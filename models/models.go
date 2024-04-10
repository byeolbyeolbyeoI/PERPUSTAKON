package models

import (
	"database/sql"
)

type User struct {
	Id int `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role int `json:"role"`
}

type UserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Book struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Genre []string `json:"genre"`
	// create login on array of genres
	Synopsis string `json:"string"`
	ReleaseYear int `json:"releaseYear"`
}

func (b *Book) GetId() int {
	return b.Id
}

type LibraryBook struct {
	Book Book `json:"book"`
	Available bool `json:"available"`
}

func (lb *LibraryBook) CheckLibraryBookAvailability(db *sql.DB) (bool, error) {
	err := db.QueryRow("SELECT available FROM books WHERE id=?", lb.Book.Id).Scan(&lb.Available)
	if err != nil {
		return false, err
		// error handling
	}

	return lb.Available, nil
}

type Peminjaman struct {
	Id int `json:"id"`
	IdBuku int `json:"idBuku"`
	IdUser int `json:"idUser"`
}


