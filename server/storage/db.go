package storage

import (
	"fmt"
	"hexcore/config"
	"hexcore/types"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresStorage() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.Envs.DB_URL), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("there was an error connecting to db %v", err)
	}
	return db, nil
}

// init the subjectSchedule table
func InitializeSubjectSchedule(db *gorm.DB) error {
	// Check if table already has data
	var count int64
	db.Model(&types.SubjectSchedule{}).Count(&count)
	if count > 0 {
		return nil // Already initialized
	}

	// Insert predefined schedule
	schedules := []types.SubjectSchedule{
		{Name: "ADA", Day: "Wednesday"},
		{Name: "ADA", Day: "Thursday"},
		{Name: "ADA", Day: "Friday"},
		{Name: "ADA", Day: "Saturday"},
		{Name: "IT", Day: "Monday"},
		{Name: "IT", Day: "Wednesday"},
		{Name: "IT", Day: "Friday"},
		{Name: "IT", Day: "Saturday"},
		{Name: "SE", Day: "Monday"},
		{Name: "SE", Day: "Wednesday"},
		{Name: "SE", Day: "Friday"},
		{Name: "SE", Day: "Saturday"},
		{Name: "IC", Day: "Tuesday"},
		{Name: "IC", Day: "Thursday"},
		{Name: "LANG", Day: "Monday"},
		{Name: "LANG", Day: "Wednesday"},
		{Name: "LANG", Day: "Thursday"},
		{Name: "LANG", Day: "Friday"},
		{Name: "ENG", Day: "Monday"},
		{Name: "ENG", Day: "Tuesday"},
		{Name: "ENG", Day: "Wednesday"},
		{Name: "ENG", Day: "Friday"},
		{Name: "OE", Day: "Tuesday"},
		{Name: "OE", Day: "Thursday"},
		{Name: "ADA Lab", Day: "Tuesday"},
		{Name: "IT Lab", Day: "Friday"},
	}

	return db.Create(&schedules).Error
}
