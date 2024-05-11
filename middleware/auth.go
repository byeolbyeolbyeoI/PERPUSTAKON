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
	ok := IsLoggedIn(c)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(
			fiber.Map{
				"success": false,
				"message": "You are not logged in",
				"code":    "User not logged in",
			},
		)
	}

	tokenString := GetTokenString(c)
	if tokenString == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"success": false,
				"message": "Error retrieving JWT token",
				"code":    "JWT Token error",
			},
		)
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
		if role != "admin" {
			return c.Status(fiber.StatusUnauthorized).JSON(
				fiber.Map{
					"success": false,
					"message": "You are not an admin",
					"code":    "NOT_ADMIN",
				},
			)
		}
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(
				fiber.Map{
					"success": false,
					"message": "Error claiming role",
					"code":    err.Error(),
				},
			)
		}
	}

	return c.Next()
}

func OnlyLibrarian(c *fiber.Ctx) error {
	ok := IsLoggedIn(c)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(
			fiber.Map{
				"success": false,
				"message": "You are not logged in",
				"code":    "User not logged in",
			},
		)
	}

	tokenString := GetTokenString(c)
	if tokenString == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"success": false,
				"message": "Error retrieving JWT token",
				"code":    "JWT Token error",
			},
		)
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
		if role != "librarian" {
			return c.Status(fiber.StatusUnauthorized).JSON(
				fiber.Map{
					"success": false,
					"message": "You are not a librarian",
					"code":    "NOT_LIBRARIAN",
				},
			)
		}
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				fiber.Map{
					"success": false,
					"message": "Error claiming role",
					"code":    err.Error(),
				},
			)
		}
	}

	return c.Next()
}

func NotLoggedIn(c *fiber.Ctx) error {
	ok := IsLoggedIn(c)
	if ok {
		return c.Status(fiber.StatusUnauthorized).JSON(
			fiber.Map{
				"success": false,
				"message": "You are already logged in",
				"code":    "User already logged in",
			},
		)
	}

	return c.Next()
}

func IsLoggedIn(c *fiber.Ctx) bool {
	tokenString := c.Cookies("token")
	if tokenString == "" {
		return false
	}

	return true
}

func GetTokenString(c *fiber.Ctx) string {
	tokenString := c.Cookies("token")
	if tokenString == "" {
		return ""
	}

	return tokenString
}

func IsLibrarian(c *fiber.Ctx) bool {
	ok := IsLoggedIn(c)
	if !ok {
		return false 
	}

	tokenString := GetTokenString(c)
	if tokenString == "" {
		return false
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
		if role != "librarian" {
			return false
	}
		if err != nil {
			log.Fatal(err)
		}
	}

	return true
}

func IsAdmin(c *fiber.Ctx) bool {
	ok := IsLoggedIn(c)
	if !ok {
		return false 
	}

	tokenString := GetTokenString(c)
	if tokenString == "" {
		return false
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
		if role != "admin" {
			return false
	}
		if err != nil {
			log.Fatal(err)
		}
	}

	return true
}

func IsUser(c *fiber.Ctx) bool {
	ok := IsLoggedIn(c)
	if !ok {
		return false 
	}

	tokenString := GetTokenString(c)
	if tokenString == "" {
		return false
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
		if role != "user" {
			return false
	}
		if err != nil {
			log.Fatal(err)
		}
	}

	return true
}