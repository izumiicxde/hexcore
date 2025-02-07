package user

import (
	"hexcore/services/auth"
	"hexcore/types"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router fiber.Router) {
	router.Post("/users/register", h.register)
	router.Post("/users/login", h.login)

	router.Get("/users", h.getAll)
	router.Get("/users/:id", h.getById)
	router.Put("/users/:id", h.update)
	router.Delete("/users/:id", h.delete)
}

func (h *Handler) login(c *fiber.Ctx) error {
	// Check if a valid token is already present in the cookies
	cookie := c.Cookies("token")
	if cookie != "" {
		claims, err := auth.ValidateToken(cookie)
		if err == nil {
			// Token is valid, return user data
			userID := int(claims["user_id"].(float64))
			user, err := h.store.GetUserById(userID)
			if err == nil {
				return c.JSON(fiber.Map{
					"message": "already logged in",
					"user":    user,
				})
			}
		}
	}

	var req struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid input"})
	}

	// Fetch user by username
	user, err := h.store.GetUserByUsername(req.Username)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid credentials"})
	}

	// Compare password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid credentials"})
	}

	// Generate JWT token
	token, err := auth.GenerateToken(user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not generate token"})
	}

	// Set token in an HTTP-only cookie (prevents JavaScript access)
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour), // Expires in 24 hours
		HTTPOnly: true,
		Secure:   true, // Use true if deploying on HTTPS
		SameSite: "Strict",
	})

	return c.JSON(fiber.Map{
		"message": "success",
		"user":    user,
	})
}

// register handles user registration
func (h *Handler) register(c *fiber.Ctx) error {
	user := new(types.User)
	// Parse request body
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}
	p, err := auth.HashPassword(user.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	user.Password = p
	user.Role = "user"
	// Create user in the database
	if err := h.store.CreateUser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

// getAll fetches all users
func (h *Handler) getAll(c *fiber.Ctx) error {
	users, err := h.store.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve users"})
	}

	return c.JSON(users)
}

// getById fetches a single user by ID
func (h *Handler) getById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	user, err := h.store.GetUserById(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	return c.JSON(user)
}

// update modifies user details
func (h *Handler) update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	user := new(types.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	user.ID = uint(id)
	if err := h.store.UpdateUser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "User updated successfully"})
}

// delete removes a user
func (h *Handler) delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	if err := h.store.DeleteUser(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "User deleted successfully"})
}
