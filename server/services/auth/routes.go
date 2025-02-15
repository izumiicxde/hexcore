package auth

import (
	"hexcore/types"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(s types.UserStore) *Handler {
	return &Handler{store: s}
}

func (h *Handler) RegisterRoutes(router fiber.Router) {
	router.Post("/signup", h.Signup)
}

func (h *Handler) Signup(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{"message": "signup"})
}
