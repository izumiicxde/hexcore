package api

import (
	"hexcore/services/attendance"
	"hexcore/services/user"
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000", // Frontend URL
		AllowMethods:     "GET, POST, PUT, DELETE",
		AllowHeaders:     "Content-Type, Authorization",
		AllowCredentials: true,
	}))

	userstore := user.NewStore(s.db)
	userhandler := user.NewHandler(userstore)
	userhandler.RegisterRoutes(api)

	attendancestore := attendance.NewAttendanceStore(s.db)
	attendancehandler := attendance.NewHandler(attendancestore)
	attendancehandler.RegisterRoutes(api)

	slog.Info("Server running at port " + s.addr)
	return app.Listen(s.addr)
}
