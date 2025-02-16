package api

import (
	"hexcore/services/attendance"
	"hexcore/services/auth"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type APIServer struct {
	addr string
	db   *gorm.DB
}

func NewAPIServer(addr string, db *gorm.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}
func (s *APIServer) Run() {
	app := fiber.New()
	subrouter := app.Group("/api/v1")

	userStore := auth.NewStore(s.db)
	authHandler := auth.NewHandler(userStore)
	authHandler.RegisterRoutes(subrouter)

	attendanceStore := attendance.NewStore(s.db)
	attendanceHandler := attendance.NewHandler(attendanceStore)
	attendanceHandler.RegisterRoutes(subrouter)

	log.Fatal(app.Listen(":" + s.addr))
}
