package storage

import (
	"fmt"
	"hexcore/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresStorage() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.Envs.NEON_URL), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("there was an error connecting to db %v", err)
	}
	return db, nil
}
