package attendance

import (
	"hexcore/services/auth"
	"hexcore/types"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	store types.AttendanceStore
}

func NewHandler(store types.AttendanceStore) *Handler {
	return &Handler{store}
}

func (h *Handler) AuthMiddleware(c *fiber.Ctx) error {
	token := c.Cookies("token")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}
	res, err := auth.ValidateToken(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized", "message": err.Error()})
	}

	userId := int(res["user_id"].(float64))
	c.Locals("userId", userId)
	return c.Next()
}

func (h *Handler) RegisterRoutes(router fiber.Router) {
	router.Use(h.AuthMiddleware)

	router.Post("/attendance", h.MarkAttendance)
}

func (h *Handler) MarkAttendance(c *fiber.Ctx) error {
	userId := c.Locals("userId").(int)

	req := new(types.AttendanceRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid input: " + err.Error()})
	}
	h.store.MarkAttendance(userId, req)
}
