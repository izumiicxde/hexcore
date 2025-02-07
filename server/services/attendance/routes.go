package attendance

import (
	"hexcore/types"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	store types.AttendanceStore
}

func NewHandler(store types.AttendanceStore) *Handler {
	return &Handler{store}
}

func (h *Handler) RegisterRoutes(router fiber.Router) {
	router.Post("/attendance", h.MarkAttendance)
	router.Get("/attendance/:id", h.GetAttendanceSummary)
}

func (h *Handler) GetAttendanceSummary(c *fiber.Ctx) error {
	userID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid user id"})
	}
	subjects, err := h.store.GetUserSubjects(uint(userID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not fetch attendance"})
	}

	response := []map[string]interface{}{}

	for _, sub := range subjects {
		attendancePercent := (float64(sub.AttendedClasses) / float64(sub.TotalTaken)) * 100
		requiredMin := int(0.75 * float64(sub.TotalTaken))
		skippable := sub.AttendedClasses - requiredMin
		if skippable < 0 {
			skippable = 0
		}

		response = append(response, map[string]interface{}{
			"subject":             sub.Name,
			"attendance_percent":  attendancePercent,
			"remaining_skippable": skippable,
		})
	}

	return c.JSON(response)
}

func (h *Handler) MarkAttendance(c *fiber.Ctx) error {
	var req struct {
		UserId  uint   `json:"user_id"`
		Subject string `json:"subject"`
		Status  string `json:"status"` // "present" or "absent"
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid input"})
	}

	// Find the subject record
	subject, err := h.store.UpdateAttendance(req.UserId, req.Subject, req.Status == "present")
	if err != nil {
		return c.JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{
		"message": "attendance updated",
		"subject": subject,
	})
}
