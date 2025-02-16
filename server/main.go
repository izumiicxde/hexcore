package main

import (
	"hexcore/cmd/api"
	"hexcore/config"
	"hexcore/storage"
)

func main() {
	db := storage.NewPostgresStorage()
	storage.AutoMigrate(db)

	api := api.NewAPIServer(config.Envs.PORT, db)

	api.Run()
}