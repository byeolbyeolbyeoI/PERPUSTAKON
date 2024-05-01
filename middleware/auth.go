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
				"error": fiber.Map{
					"message": "You are not logged in",
					"code":    "AUTHORIZE_ERROR"}})
	}

	tokenString := GetTokenString(c)
	if tokenString == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"error": fiber.Map{
					"message": "Error retrieving JWT Token",
					"code":    "TOKEN_ERROR"}})
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
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "youre not an admin"})
		}
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
	}

	return c.Next()
}

func OnlyLibrarian(c *fiber.Ctx) error {
	ok := IsLoggedIn(c)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(
			fiber.Map{
				"error": fiber.Map{
					"message": "You are not logged in",
					"code":    "AUTHORIZE_ERROR"}})
	}

	tokenString := GetTokenString(c)
	if tokenString == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"error": fiber.Map{
					"message": "Error retrieving JWT Token",
					"code":    "TOKEN_ERROR"}})
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
		if role != "librarian" {
			return c.Status(fiber.StatusUnauthorized).JSON(
				fiber.Map{
					"error": fiber.Map{
						"message": "You are not a librarian",
						"code":    "AUTHORIZE_ERROR"}})
		}
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				fiber.Map{
					"error": fiber.Map{
						"message": err.Error(),
						"code":    err.Error()}})
		}
	}

	return c.Next()
}

func NotLoggedIn(c *fiber.Ctx) error {
	ok := IsLoggedIn(c)
	if ok {
		return c.Status(fiber.StatusUnauthorized).JSON(
			fiber.Map{
				"error": fiber.Map{
					"message": "You are already logged in",
					"code":    "AUTHORIZE_ERROR"}})
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
