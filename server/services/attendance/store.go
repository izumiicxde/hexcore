package attendance

import (
	"fmt"
	"hexcore/config"
	"hexcore/types"
	"time"

	"gorm.io/gorm"
)

type Store struct {
	db *gorm.DB
}

func NewAttendanceStore(db *gorm.DB) *Store {
	return &Store{db}
}

func (s *Store) MarkAttendance(a *types.Attendance) error {
	if err := config.Validator.Struct(a); err != nil {
		return err
	}

	a.Date = time.Now()
	result := s.db.Create(a)

	if result.RowsAffected == 0 {
		return fmt.Errorf("error creating attendance")
	}
	return result.Error
}
