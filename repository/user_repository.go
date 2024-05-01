package repository

import (
	"database/sql"
	"errors"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"

	"perpustakaan/models"
	APIError "perpustakaan/error"
)

// for future usage
type UserStore interface {
	CreateUser(user models.UserInput) *APIError.APIError
	GetUserById(user models.UserInput) (*models.User, *APIError.APIError)
	GetAllUser() ([]models.User, *APIError.APIError)
}

type UserRepository struct {
	DB *sql.DB
}

// userInput is readonly, user is modifyable
// decided to not pass dbUser as pointer since the value has nothing to do outside of the function
func (s *UserRepository) CreateUser(user models.UserInput) *APIError.APIError {
	var dbUser models.User
	// do the query, user only need read permission, there is chang in dbUser
	err := s.DB.QueryRow("SELECT username FROM users WHERE username=?", user.Username).Scan(&dbUser.Username)
	if err != nil { // if there is an error
		// check if error is not errnorows (username available)
		if !errors.Is(err, sql.ErrNoRows) {
			return APIError.NewAPIError(fiber.StatusInternalServerError, "Error checking username availability", err.Error())
		}
		// if the above code is passed, meaning the error is no rows (username available)

		//user.Password only need read permission
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return APIError.NewAPIError(fiber.StatusInternalServerError, "Error hashing the password", err.Error())
		}

		// insert the credentials
		_, err = s.DB.Exec("INSERT INTO users (username, password, role) VALUES (?, ?, ?)", user.Username, hashedPassword, "user")
		if err != nil {
			return APIError.NewAPIError(fiber.StatusInternalServerError, "Error creating user", err.Error())
		}

		return nil
	}

	// if no err, the Scan() function scanned a row, meaning the username is already exists
	return APIError.NewAPIError(fiber.StatusConflict, "Username already exists", "USERNAME_TAKEN")
}
 
func (s *UserRepository) CheckPassword(user models.UserInput, dbUser models.User) *APIError.APIError {
	err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if err != nil {
		return APIError.NewAPIError(fiber.StatusInternalServerError, "Incorrect Password", err.Error())
	}

	return nil
}

func (s *UserRepository) GetUserByUsername(username string) (models.User, *APIError.APIError) {
	var dbUser models.User
	err := s.DB.QueryRow("SELECT id, username, password, role FROM users WHERE username=?", username).Scan(&dbUser.Id, &dbUser.Username, &dbUser.Password, &dbUser.Role)
	if err == sql.ErrNoRows {
		return dbUser, APIError.NewAPIError(fiber.StatusInternalServerError, "Username is not registered", err.Error())
	}

	return dbUser, nil
}

func (s *UserRepository) GetUserById(id int) (models.User, *APIError.APIError) {
	var dbUser models.User
	err := s.DB.QueryRow("SELECT id, username, password, role FROM users WHERE id=?", id).Scan(&dbUser.Id, &dbUser.Username, &dbUser.Password, &dbUser.Role)
	if err == sql.ErrNoRows {
		return dbUser, APIError.NewAPIError(fiber.StatusInternalServerError, "Id is not registered", err.Error())
	}
	return dbUser, nil
}

func (s *UserRepository) GetAllUsers() ([]models.User, *APIError.APIError) {
	rows, err := s.DB.Query("SELECT id, username, password, role FROM users")
	if err != nil {
		return nil, APIError.NewAPIError(fiber.StatusInternalServerError, "Error retrieving rows", err.Error())
	}
	defer rows.Close()

	var users []models.User
	var user models.User

	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Role); err != nil {
			return nil, APIError.NewAPIError(fiber.StatusInternalServerError, "Error scanning rows", err.Error())
		}

		users = append(users, user)
	}

	return users, nil
}
