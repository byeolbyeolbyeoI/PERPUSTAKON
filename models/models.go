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

func (lb *LibraryBook) CheckAvailability(db *sql.DB, id int) (bool, error) {
	// logic
	return lb.Available, nil
}

type Peminjaman struct {
	Id int `json:"id"`
	IdBuku int `json:"idBuku"`
	IdUser int `json:"idUser"`
}


