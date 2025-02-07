package api

import (
	"hexcore/services/attendance"
	"hexcore/services/user"
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type APIServer struct {
	addr string
	db   *gorm.DB
}

func NewAPIServer(addr string, db *gorm.DB) *APIServer {
	return &APIServer{addr: addr, db: db}
}

func (s *APIServer) Run() error {
	app := fiber.New()
	api := app.Group("/api/v1")

	userstore := user.NewStore(s.db)
	userhandler := user.NewHandler(userstore)
	userhandler.RegisterRoutes(api)

	attendancestore := attendance.NewAttendanceStore(s.db)
	attendancehandler := attendance.NewHandler(attendancestore)
	attendancehandler.RegisterRoutes(api)

	slog.Info("Server running at port " + s.addr)
	return app.Listen(s.addr)
}
