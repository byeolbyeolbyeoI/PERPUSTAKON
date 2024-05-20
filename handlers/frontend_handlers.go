package handlers

import (
	"github.com/gofiber/fiber/v2"

	"perpustakaan/middleware"
)

func (h *Handler) Login(c *fiber.Ctx) error {
	return c.Render("./frontend/views/login.html", nil)
}

func (h *Handler) LibrarianBookList(c *fiber.Ctx) error {
	if middleware.IsLibrarian(c) == false {
		return c.Render("./frontend/views/error-auth.html", nil)
	}

	return c.Render("./frontend/views/librarian/book-list.html", nil)
}

func (h *Handler) LibrarianUserList(c *fiber.Ctx) error {
	if middleware.IsLibrarian(c) == false {
		return c.Render("./frontend/views/error-auth.html", nil)
	}

	return c.Render("./frontend/views/librarian/user-list.html", nil)
}

func (h *Handler) LibrarianAddBook(c *fiber.Ctx) error {
	if middleware.IsLibrarian(c) == false {
		return c.Render("./frontend/views/error-auth.html", nil)
	}
	
	return c.Render("./frontend/views/librarian/add-book.html", nil)
}

func (h *Handler) LibrarianDeleteBook(c *fiber.Ctx) error {
	if middleware.IsLibrarian(c) == false {
		return c.Render("./frontend/views/error-auth.html", nil)
	}

	return c.Render("./frontend/views/librarian/delete-book.html", nil)
}

func (h *Handler) LibrarianBorrowBook(c *fiber.Ctx) error {
	if middleware.IsLibrarian(c) == false {
		return c.Render("./frontend/views/error-auth.html", nil)
	}

	return c.Render("./frontend/views/librarian/borrow-book.html", nil)
}

func (h *Handler) LibrarianReturnBook(c *fiber.Ctx) error {
	if middleware.IsLibrarian(c) == false {
		return c.Render("./frontend/views/error-auth.html", nil)
	}

	return c.Render("./frontend/views/librarian/return-book.html", nil)
}

func (h *Handler) AdminDashboard(c *fiber.Ctx) error {
	if middleware.IsAdmin(c) == false {
		return c.Render("./frontend/views/error-auth.html", nil)
	}

	return c.Render("./frontend/views/admin/dashboard.html", nil)
}

func (h *Handler) AdminUserList(c *fiber.Ctx) error {
	if middleware.IsAdmin(c) == false {
		return c.Render("./frontend/views/error-auth.html", nil)
	}

	return c.Render("./frontend/views/admin/user-list.html", nil)
}

func (h *Handler) AdminAddUser(c *fiber.Ctx) error {
	if middleware.IsAdmin(c) == false {
		return c.Render("./frontend/views/error-auth.html", nil)
	}

	return c.Render("./frontend/views/admin/add-user.html", nil)
}

func (h *Handler) AdminDeleteUser(c *fiber.Ctx) error {
	if middleware.IsAdmin(c) == false {
		return c.Render("./frontend/views/error-auth.html", nil)
	}

	return c.Render("./frontend/views/admin/delete-user.html", nil)
}

func (h *Handler) UserDashboard(c *fiber.Ctx) error {
	if middleware.IsUser(c) == false {
		return c.Render("./frontend/views/error-auth.html", nil)
	}

	return c.Render("./frontend/views/user/dashboard.html", nil)
}

func (h *Handler) UserBookList(c *fiber.Ctx) error {
	if middleware.IsUser(c) == false {
		return c.Render("./frontend/views/error-auth.html", nil)
	}

	return c.Render("./frontend/views/user/book-list.html", nil)
}
