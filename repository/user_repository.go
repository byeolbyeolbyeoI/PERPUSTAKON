package repository

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"

	"perpustakaan/models"
)

type APIError struct {
	Status int `json:"status"`
	Error APIErrorDetails `json:"error"`
}

type APIErrorDetails struct {
	Message string `json:"message"`
	Code string `json:"code"`
}

func (e *APIError) ToJSON() string {
	b, err := json.Marshal(e)
	if err != nil {
		return fmt.Sprintf("Error converting to JSON: %s", err.Error())
	}

	return string(b)
}

func newAPIError(status int, message string, code string) *APIError {
	return &APIError{
		Status: status, Error: APIErrorDetails{
			Message: message,
			Code: code,
		},
	}
}

// for future usage
type UserStore interface {
	CreateUser(user models.UserInput) error
	GetUserById(user models.UserInput) (*models.User, error)
	GetAllUser() ([]models.User, error)
}

type UserRepository struct {
	DB *sql.DB
}

// userInput is readonly, user is modifyable
// decided to not pass dbUser as pointer since the value has nothing to do outside of the function
func (s *UserRepository) CreateUser(user models.UserInput) *APIError {
	var dbUser models.User
	// do the query, user only need read permission, there is chang in dbUser
	err := s.DB.QueryRow("SELECT username FROM users WHERE username=?", user.Username).Scan(&dbUser.Username)
	defer s.DB.Close()
	if err != nil { // if there is an error
		// check if error is not errnorows (username available)
		if !errors.Is(err, sql.ErrNoRows) {
			return newAPIError(fiber.StatusInternalServerError, "Error checking username availability", err.Error())
		}
		// if the above code is passed, meaning the error is no rows (username available)

		//user.Password only need read permission
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return newAPIError(fiber.StatusInternalServerError, "Error hashing the password", err.Error())
		}

		// insert the credentials
		_, err = s.DB.Exec("INSERT INTO users (username, password, role) VALUES (?, ?, ?)", user.Username, hashedPassword, 1)
		if err != nil {
			return newAPIError(fiber.StatusInternalServerError, "Error creating user", err.Error())
		}

		return nil
	}

	// if no err, the Scan() function scanned a row, meaning the username is already exists
	return newAPIError(fiber.StatusConflict, "Username already exists", "USERNAME_TAKEN")
}

func (s *UserRepository) GetUserByUsername(user models.UserInput) (models.User, *APIError) {
	var dbUser models.User
	err := s.DB.QueryRow("SELECT id, username, password, role FROM users WHERE username=?", user.Username).Scan(&dbUser.Id, &dbUser.Username, &dbUser.Password, &dbUser.Role)
	if err == sql.ErrNoRows {
		return dbUser, newAPIError(fiber.StatusInternalServerError, "Username is not registered", err.Error())
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if err != nil {
		return dbUser, newAPIError(fiber.StatusInternalServerError, "Incorrect Password", err.Error())
	}

	return dbUser, nil
}

func (s *UserRepository) GetAllUser() ([]models.User, *APIError) {
	rows, err := s.DB.Query("SELECT id, username, password, role FROM users")
	if err != nil {
		return nil, newAPIError(fiber.StatusInternalServerError, "Error retrieving rows", err.Error())
	}
	defer rows.Close()

	var users []models.User
	var user models.User

	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Role); err != nil {
			return nil, newAPIError(fiber.StatusInternalServerError, "Error scanning rows", err.Error())
		}

		users = append(users, user)
	}

	return users, nil
}
