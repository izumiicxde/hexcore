package main

import (
	"hexcore/cmd/api"
	"hexcore/config"
	"hexcore/storage"
	"log"
)

func main() {
	db := storage.NewPostgresStorage()
	storage.AutoMigrate(db)
	if err := storage.CreateSchedules(db); err != nil {
		log.Fatal(err)
	}

	api := api.NewAPIServer(config.Envs.PORT, db)

	api.Run()
}
