package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"

	"perpustakaan/models"
)

// for future usage
type UserStore interface {
	CreateUser(user models.UserInput, dbUser models.User) error
	GetUserById(id int) (*models.User, error)
	GetAllUser() ([]models.User, error)
}

type UserRepository struct {
	DB *sql.DB
}

// userInput is readonly, user is modifyable
// decided to not pass dbUser as pointer since the value has nothing to do outside of the function
func (s *UserRepository) CreateUser(user models.UserInput, dbUser models.User) error {
	// do the query, user only need read permission, there is chang in dbUser
	err := s.DB.QueryRow("SELECT username FROM ? WHERE username=?", dbUser.TableName(),user.Username).Scan(&dbUser.Username)
	defer s.DB.Close()
	if err != nil { // if there is an error
		// check if error is not errnorows (username available)
		if !errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("Error checking username availability: %s", err.Error())
		}
		// if the above code is passed, meaning the error is no rows (username available)

		//user.Password only need read permission
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("hash kah")
		}

		// insert the credentials
		_, err = s.DB.Exec("INSERT INTO users (username, password, role) VALUES (?, ?, ?)", user.Username, hashedPassword, 1)
		if err != nil {
			return fmt.Errorf("insert kah")
		}

		return nil
	}

	// if no err, the Scan() function scanned a row, meaning the username is already exists
	return fmt.Errorf("Username '%s' already exists", user.Username)
}

func (s *UserRepository) GetUserById(user models.UserInput, dbUser *models.User) (*models.User, error) {
	err := s.DB.QueryRow("SELECT id, username, password, role FROM ? WHERE username=?", dbUser.TableName(), user.Username).Scan(&dbUser.Id, &dbUser.Username, &dbUser.Password, &dbUser.Role)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("Username not registered")
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if err != nil {
		return nil, fmt.Errorf("Incorrect password")
	}

	return dbUser, nil
}

func (s *UserRepository) GetAllUser() ([]models.User, error) {
	rows, err := s.DB.Query("SELECT id, username, password, role FROM users")
	if err != nil {
		return nil, fmt.Errorf("error retrieving rows")
	}
	defer rows.Close()

	var users []models.User
	var user models.User

	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Role); err != nil {
			return nil, fmt.Errorf("error scanning rows")
		}

		users = append(users, user)
	}

	return users, nil
}
