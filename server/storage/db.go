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

func CreateSchedules(db *gorm.DB) error {
	// Check if schedules already exist to avoid duplicate entries
	var count int64
	if err := db.Model(&types.Schedule{}).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return nil // Table already has data, no need to insert again
	}

	// Insert schedules into the database
	return db.Create(&config.Schedules).Error
}
