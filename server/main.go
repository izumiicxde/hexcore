package main

import (
	"log/slog"
	"os"

	"hexcore/cmd/api"
	"hexcore/config"
	"hexcore/storage"
	"hexcore/types"

	"gorm.io/gorm"
)

func main() {

	db, err := storage.NewPostgresStorage()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	migrate(db)

	if err = storage.InitializeSubjectSchedule(db); err != nil {
		slog.Error("error initializing subject schedule", "err", err)
	}

	api := api.NewAPIServer(config.Envs.PORT, db) // provide proper addr string and db

	if err := api.Run(); err != nil {
		slog.Error("error starting the server", "err", err)
	}
}

func migrate(db *gorm.DB) {
	db.Exec("CREATE UNIQUE INDEX user_subject_date ON attendances (user_id, subject_id, date)")
	if err := db.AutoMigrate(
		&types.User{},
		&types.Subject{},
		&types.Attendance{},
		&types.SubjectSchedule{},
	); err != nil {
		slog.Error("error migrating the db", "err", err.Error())
		os.Exit(1)
	}
}
