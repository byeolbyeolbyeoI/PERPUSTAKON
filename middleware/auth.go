package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"os"
	_ "perpustakaan/models"
)

func OnlyAdmin(c *fiber.Ctx) error {
	tokenString, ok := IsLoggedIn(c)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "you're not even logged in dude"})
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		log.Fatal(err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// if exp, id not valid
		role := fmt.Sprint(claims["role"])
		if role != "3" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "youre not an admin"})
		}
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
	}

	return c.Next()
}

func OnlyLibrarian(c *fiber.Ctx) error {
	tokenString, ok := IsLoggedIn(c)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "you're not even logged in dude"})
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		fmt.Println("token : ", token)
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		log.Fatal(err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// if exp, id not valid
		role := fmt.Sprint(claims["role"])
		if role != "2" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "youre not a librarian/admin"})
		}
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
	}

	return c.Next()
}

func IsLoggedIn(c *fiber.Ctx) (string, bool) {
	tokenString := c.Cookies("token")
	if tokenString == "" {
		return "", false
	}

	return tokenString, true
}
