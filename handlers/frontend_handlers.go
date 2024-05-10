package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) Login(c *fiber.Ctx) error {
	return c.Render("./frontend/html/login.html", nil)
}

func (h *Handler) LibrarianDashboard(c *fiber.Ctx) error {
	return c.Render("./frontend/html/librarian-dashboard.html", nil)
}

func (h *Handler) LibrarianAddBook(c *fiber.Ctx) error {
	return c.Render("./frontend/html/librarian-add-book.html", nil)
}

func (h *Handler) LibrarianDeleteBook(c *fiber.Ctx) error {
	return c.Render("./frontend/html/librarian-delete-book.html", nil)
}

func (h *Handler) LibrarianBorrow(c *fiber.Ctx) error {
	return c.Render("./frontend/html/librarian-borrow.html", nil)
}

func (h *Handler) LibrarianReturn(c *fiber.Ctx) error {
	return c.Render("./frontend/html/librarian-return.html", nil)
}