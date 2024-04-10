package repository

import (
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"errors"

	"perpustakaan/models"
	"perpustakaan/service"
)

type MyError struct {
  message string
}

func (e *MyError) Error() string {
  return e.message
}

/*
func someFunction() error {
  // Perform some operation
  if err := someOtherFunction(); err != nil {
    return &MyError{message: "Error occurred: " + err.Error()}
  }
  return nil
}
*/

type UserStore interface {
	CreateUser(user models.UserInput, dbUser models.User) error
	GetUserById(id int) (*models.User, error)
	GetAllUser() ([]models.User, error)
}

type UserRepository struct {
	db *sql.DB
}

// userInput is readonly, user is modifyable
// decided to not pass dbUser as pointer since the value has nothing to do outside of the function
func (s *UserRepository) CreateUser(user models.UserInput, dbUser models.User) error {	
	// do the query, user only need read permission, there is chang in dbUser
	err := s.db.QueryRow("SELECT username FROM users WHERE username=?", user.Username).Scan(&dbUser.Username)
	if err != nil { // if there is an error
		// check if error is not errnorows (username available)
		if !errors.Is(err, sql.ErrNoRows){
			return fmt.Errorf("Error checking username availability: ", err.Error())
		}
		// if the above code is passed, meaning the error is no rows (username available)

		//user.Password only need read permission
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		
		// insert the credentials
		_, err = s.db.Exec("INSERT INTO users (username, password, role) VALUES (?, ?, ?)", user.Username, hashedPassword, 1)
		if err != nil {
			return err
		}
	}

	// if no err, the Scan() function scanned a row, meaning the username is already exists
	return fmt.Errorf("Username '%s' already exists", user.Username)
}

func (s *UserRepository) GetUserById(user *models.User) (*models.User, error) {
	
	return user, nil
}

func (s *UserRepository) GetAllUser() ([]models.User, error) {
	var users []models.User
	
	return users, nil 
}

func (r *Repository) Signup(c *fiber.Ctx) error {
	db := r.DB

	var user models.UserInput 
	var dbUser models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	err := db.QueryRow("SELECT username FROM users WHERE username=?", user.Username).Scan(&dbUser.Username)
	if err != sql.ErrNoRows { // if username has already taken
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Username is already taken"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	_, err = db.Exec("INSERT INTO users (username, password, role) VALUES (?, ?, ?)", user.Username, hashedPassword, 1)
	if err != nil {
		fmt.Println("sini kah")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User created successfully"})
}

func (r *Repository) Login(c *fiber.Ctx) error {
	db := r.DB

	type userStruct struct {
		Username string
		Password string
	}

	var user userStruct
	var dbUser models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	err := db.QueryRow("SELECT id, username, password, role FROM users WHERE username=?", user.Username).Scan(&dbUser.Id, &dbUser.Username, &dbUser.Password, &dbUser.Role)
	if err == sql.ErrNoRows {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Username is not registered"})
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Incorrect password"})
	}
	// generate token
	token := service.GenerateJWT(dbUser.Id, dbUser.Username, dbUser.Role)
	// sign the jwt
	tokenString, err := service.SignToken(token)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Unable to sign the token"})
	}

	// set the cookie
	cookie := &fiber.Cookie{
		Name:     "token",
		Value:    tokenString,
		MaxAge:   3600 * 24 * 30,
		HTTPOnly: true,
		SameSite: "lax",
	}
	c.Cookie(cookie)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User logged in succesfully"})
}

func (r *Repository) GetUsers(c *fiber.Ctx) error {
	db := r.DB

	rows, err := db.Query("SELECT id, username, password, role FROM users")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	defer rows.Close()

	var users []models.User
	var user models.User

	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Role); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		users = append(users, user)
	}

	return c.Status(fiber.StatusOK).JSON(users)
}
