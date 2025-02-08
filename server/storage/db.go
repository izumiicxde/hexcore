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
		{Name: "ADA", Day: "Wednesday", StartTime: "8:30 AM", EndTime: "9:30 AM"},
		{Name: "ADA", Day: "Thursday", StartTime: "9:30 AM", EndTime: "10:30 AM"},
		{Name: "ADA", Day: "Friday", StartTime: "9:30 AM", EndTime: "10:30 AM"},
		{Name: "ADA", Day: "Saturday", StartTime: "8:30 AM", EndTime: "9:30 AM"},

		{Name: "IT", Day: "Monday", StartTime: "9:30 AM", EndTime: "10:30 AM"},
		{Name: "IT", Day: "Wednesday", StartTime: "11:45 AM", EndTime: "12:45 PM"},
		{Name: "IT", Day: "Friday", StartTime: "10:45 AM", EndTime: "11:45 AM"},
		{Name: "IT", Day: "Saturday", StartTime: "9:30 AM", EndTime: "10:30 AM"},

		{Name: "SE", Day: "Monday", StartTime: "1:30 PM", EndTime: "2:30 PM"},
		{Name: "SE", Day: "Tuesday", StartTime: "8:30 AM", EndTime: "9:30 AM"},
		{Name: "SE", Day: "Wednesday", StartTime: "1:30 PM", EndTime: "2:30 PM"},
		{Name: "SE", Day: "Saturday", StartTime: "10:45 AM", EndTime: "11:45 AM"},

		{Name: "IC", Day: "Wednesday", StartTime: "9:30 AM", EndTime: "10:30 AM"},
		{Name: "IC", Day: "Thursday", StartTime: "10:45 AM", EndTime: "11:45 AM"},

		{Name: "LANG", Day: "Monday", StartTime: "10:45 AM", EndTime: "11:45 AM"},
		{Name: "LANG", Day: "Wednesday", StartTime: "10:45 AM", EndTime: "11:45 AM"},
		{Name: "LANG", Day: "Thursday", StartTime: "11:45 AM", EndTime: "12:45 PM"},
		{Name: "LANG", Day: "Friday", StartTime: "11:45 AM", EndTime: "12:45 PM"},

		{Name: "ENG", Day: "Monday", StartTime: "8:30 AM", EndTime: "9:30 AM"},
		{Name: "ENG", Day: "Tuesday", StartTime: "2:30 PM", EndTime: "3:30 PM"},
		{Name: "ENG", Day: "Wednesday", StartTime: "2:30 PM", EndTime: "3:30 PM"},
		{Name: "ENG", Day: "Friday", StartTime: "8:30 AM", EndTime: "9:30 AM"},

		{Name: "OE", Day: "Tuesday", StartTime: "1:30 PM", EndTime: "2:30 PM"},
		{Name: "OE", Day: "Thursday", StartTime: "1:30 PM", EndTime: "2:30 PM"},

		{Name: "ADA Lab", Day: "Tuesday", StartTime: "9:30 AM", EndTime: "11:45 AM"},
		{Name: "IT Lab", Day: "Friday", StartTime: "1:30 PM", EndTime: "4:30 PM"},
	}

	return db.Create(&schedules).Error
}
