package attendance

import (
	"fmt"
	"hexcore/middleware"
	"hexcore/types"
	"hexcore/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	store types.AttendanceStore
}

func NewHandler(s types.AttendanceStore) *Handler {
	return &Handler{store: s}
}

// RegisterRoutes sets up the attendance-related routes
func (h *Handler) RegisterRoutes(router fiber.Router) {
	attendance := router.Group("/attendance", middleware.AuthMiddleware())

	attendance.Get("/today", h.GetTodaysClasses)
	attendance.Post("/mark", h.MarkAttendance)
	attendance.Get("/summary", h.GetAttendanceSummary)
	attendance.Get("/skippable", h.CalculateSkippableClasses)
	attendance.Get("/is-marked/:subjectId", h.IsAttendanceMarked)
	attendance.Get("/day", h.GetClassesByDay)
	attendance.Get("/progress", h.HandleGetClassesTillToday)
}

func (h *Handler) HandleGetClassesTillToday(c *fiber.Ctx) error {
	// Retrieve userID from context (assuming middleware sets it)
	userID, ok := c.Locals("userID").(uint)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthorized"})
	}

	// Fetch all classes till today
	classes, err := h.store.GetClassesTillToday(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	// Send response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Fetched classes",
		"classes": classes,
	})
}

func (h *Handler) GetClassesByDay(c *fiber.Ctx) error {
	day := c.Query("day")

	if day == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "day parameter is required"})
	}

	day = strings.ToUpper(day[:1]) + strings.ToLower(day[1:]) // Normalize to title case

	classes, err := h.store.GetClassesByDay(day, 0)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"classes": classes})
}

func (h *Handler) GetTodaysClasses(c *fiber.Ctx) error {
	userId := c.Locals("userId").(uint)
	classes, err := h.store.GetTodaysClasses(userId)
	if err != nil {
		return utils.WriteError(c, http.StatusInternalServerError, err)
	}

	return c.Status(200).JSON(fiber.Map{"message": "success", "classes": classes})
}

func (h *Handler) MarkAttendance(c *fiber.Ctx) error {
	var req struct {
		SubjectID uint `json:"subjectId"`
		Status    bool `json:"status"`
	}

	if err := c.BodyParser(&req); err != nil {
		return utils.WriteError(c, http.StatusBadRequest, err)
	}

	userId := c.Locals("userId").(uint)

	err := h.store.MarkAttendance(userId, req.SubjectID, req.Status)
	if err != nil {
		return utils.WriteError(c, http.StatusInternalServerError, err)
	}

	return utils.WriteJSON(c, http.StatusOK, map[string]any{"message": "Attendance marked successfully"})
}

func (h *Handler) GetAttendanceSummary(c *fiber.Ctx) error {

	userId := c.Locals("userId").(uint)

	summary, err := h.store.GetAttendanceSummary(userId)
	if err != nil {
		return utils.WriteError(c, http.StatusInternalServerError, err)
	}

	return utils.WriteJSON(c, http.StatusOK, fiber.Map{"message": "fetched user details", "success": true, "summary": summary})
}

func (h *Handler) CalculateSkippableClasses(c *fiber.Ctx) error {
	userId := c.Locals("userId").(uint)

	skippable, err := h.store.CalculateSkippableClasses(userId)
	if err != nil {
		return utils.WriteError(c, http.StatusInternalServerError, err)
	}

	return utils.WriteJSON(c, http.StatusOK, skippable)
}

func (h *Handler) IsAttendanceMarked(c *fiber.Ctx) error {
	subjectID, err := strconv.Atoi(c.Params("subjectId"))
	if err != nil {
		return utils.WriteError(c, http.StatusBadRequest, fmt.Errorf("invalid subject ID"))
	}

	userId := c.Locals("userId").(uint)
	isMarked, err := h.store.IsAttendanceMarked(userId, uint(subjectID))
	if err != nil {
		return utils.WriteError(c, http.StatusInternalServerError, err)
	}

	return utils.WriteJSON(c, http.StatusOK, map[string]any{"is_marked": isMarked})
}

// func (h *Handler) ResetAttendance(c *fiber.Ctx) error {
// 	err := h.store.ResetAttendance()
// 	if err != nil {
// 		return utils.WriteError(c, http.StatusInternalServerError, err)
// 	}

// 	return utils.WriteJSON(c, http.StatusOK, map[string]any{"message": "Attendance records reset"})
// }
