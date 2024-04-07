package middleware

import (
	"github.com/gofiber/fiber/v2"
	"fmt"
	"os"
	"log"
	_ "perpustakaan/models"
	"github.com/golang-jwt/jwt/v5"
)

func OnlyAdmin(c *fiber.Ctx) error {
	tokenString := c.Cookies("token")
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "you're not even logged in dude"})
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// check if signing methd is valid
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

