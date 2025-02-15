package storage

import (
	"hexcore/config"
	"hexcore/types"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresStorage() *gorm.DB {
	db, err := gorm.Open(postgres.Open(config.Envs.DB_URL))
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func AutoMigrate(db *gorm.DB) {
	if err := db.AutoMigrate(&types.User{}, &types.Schedule{}, &types.Subject{}, &types.Attendance{}); err != nil {
		log.Fatal(err)
	}
}
