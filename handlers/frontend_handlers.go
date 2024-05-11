package handlers

import (
	"github.com/gofiber/fiber/v2"

	"perpustakaan/middleware"
)

func (h *Handler) Login(c *fiber.Ctx) error {
	return c.Render("./frontend/html/login.html", nil)
}

func (h *Handler) LibrarianBookList(c *fiber.Ctx) error {
	if middleware.IsLibrarian(c) == false {
		return c.Render("./frontend/html/error-auth.html", nil)
	}

	return c.Render("./frontend/html/librarian/book-list.html", nil)
}

func (h *Handler) LibrarianUserList(c *fiber.Ctx) error {
	if middleware.IsLibrarian(c) == false {
		return c.Render("./frontend/html/error-auth.html", nil)
	}

	return c.Render("./frontend/html/librarian/user-list.html", nil)
}

func (h *Handler) LibrarianAddBook(c *fiber.Ctx) error {
	if middleware.IsLibrarian(c) == false {
		return c.Render("./frontend/html/error-auth.html", nil)
	}
	
	return c.Render("./frontend/html/librarian/add-book.html", nil)
}

func (h *Handler) LibrarianDeleteBook(c *fiber.Ctx) error {
	if middleware.IsLibrarian(c) == false {
		return c.Render("./frontend/html/error-auth.html", nil)
	}

	return c.Render("./frontend/html/librarian/delete-book.html", nil)
}

func (h *Handler) LibrarianBorrow(c *fiber.Ctx) error {
	if middleware.IsLibrarian(c) == false {
		return c.Render("./frontend/html/error-auth.html", nil)
	}

	return c.Render("./frontend/html/librarian/borrow.html", nil)
}

func (h *Handler) LibrarianReturn(c *fiber.Ctx) error {
	if middleware.IsLibrarian(c) == false {
		return c.Render("./frontend/html/error-auth.html", nil)
	}

	return c.Render("./frontend/html/librarian/return.html", nil)
}
