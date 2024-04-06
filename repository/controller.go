package repository

import (
	"database/sql"
	"log"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"

	_ "perpustakaan/middleware"
	"perpustakaan/models"
)

func (r *Repository) Signup(c *fiber.Ctx) error {
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

	err := db.QueryRow("SELECT username FROM users WHERE username=?", user.Username).Scan(&dbUser.Username)
	if err != sql.ErrNoRows { // if username has already taken
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Username is already taken"})
	}

	/*
		if err != nil {
			fmt.Println("atas kah")
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
	*/

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

	err := db.QueryRow("SELECT id, username, password FROM users WHERE username=?", user.Username).Scan(&dbUser.Id, &dbUser.Username, &dbUser.Password)
	if err == sql.ErrNoRows {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Username is not registered"})
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Incorrect password"})
	}

	// *jwt.Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": dbUser.Id,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// complete, signed jwt
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}


	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    tokenString,
		MaxAge:   3600 * 24 * 30,
		HTTPOnly: true,
		SameSite: "lax",
	})

	// parse the token
	token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// no idea might check on it later
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		log.Fatal(err)
	}

	// taking the claims from token using jwt.MapClaims type??
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// if exp, id not valid
		var subject models.User	
		err := db.QueryRow("SELECT id, username, password, role FROM users WHERE id = ?", claims["sub"]).Scan(&subject.Id, &subject.Username, &subject.Password, &subject.Role)
		if err == sql.ErrNoRows {
			// no rows, no way??
			log.Fatal(err)	
		}

		c.Set("subject", subject)

	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"token": tokenString})
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
