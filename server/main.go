package main

import (
	"log/slog"
	"os"

	"hexcore/cmd/api"
	"hexcore/config"
	"hexcore/storage"
	"hexcore/types"
)

func main() {

	db, err := storage.NewPostgresStorage()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	if err = db.AutoMigrate(
		&types.User{},
		&types.Subject{},
	); err != nil {
		slog.Error(err.Error())
	}

	api := api.NewAPIServer(config.Envs.PORT, db) // provide proper addr string and db

	if err := api.Run(); err != nil {
		slog.Error("error starting the server", "err", err)
	}
}
