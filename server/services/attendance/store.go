package attendance

import (
	"fmt"
	"time"

	"hexcore/types"

	"gorm.io/gorm"
)

type Store struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetTodaysClasses(userID uint) ([]types.Subject, error) {
	var subjects []types.Subject
	today := time.Now().Weekday().String()

	err := s.db.
		Joins("JOIN schedules ON schedules.subject_name = subjects.name").
		Where("schedules.day = ?", today).
		Find(&subjects).Error

	if err != nil {
		return nil, fmt.Errorf("failed to fetch today's classes: %v", err)
	}

	return subjects, nil
}

// GetClassesByDay fetches the subjects for a given day provided in query params.
func (s *Store) GetClassesByDay(day string) ([]types.ClassSchedule, error) {
	var classes []types.ClassSchedule

	// Validate the day input
	validDays := map[string]bool{
		"Monday": true, "Tuesday": true, "Wednesday": true, "Thursday": true,
		"Friday": true, "Saturday": true, "Sunday": true,
	}
	if !validDays[day] {
		return nil, fmt.Errorf("invalid day provided: %s", day)
	}

	// Query schedules and subjects for the given day
	err := s.db.
		Table("schedules").
		Select("subjects.*, schedules.start_time, schedules.end_time").
		Joins("JOIN subjects ON schedules.subject_name = subjects.name").
		Where("schedules.day = ?", day).
		Scan(&classes).Error

	if err != nil {
		return nil, fmt.Errorf("failed to fetch classes for %s: %v", day, err)
	}

	return classes, nil
}

func (s *Store) MarkAttendance(userID uint, subjectID uint, status bool) error {
	today := time.Now().Truncate(24 * time.Hour)

	// Check if attendance is already marked
	var existing types.Attendance
	if err := s.db.Where("user_id = ? AND subject_id = ? AND date = ?", userID, subjectID, today).First(&existing).Error; err == nil {
		return fmt.Errorf("attendance already marked for subject %d", subjectID)
	}

	attendance := types.Attendance{
		UserID:    userID,
		SubjectID: subjectID,
		Date:      today,
		Status:    status,
	}

	if err := s.db.Create(&attendance).Error; err != nil {
		return fmt.Errorf("failed to mark attendance: %v", err)
	}

	// Update attended class count in the subject
	if status {
		if err := s.db.Model(&types.Subject{}).
			Where("id = ?", subjectID).
			Update("attended_classes", gorm.Expr("attended_classes + 1")).Error; err != nil {
			return fmt.Errorf("failed to update attended class count: %v", err)
		}
	}

	return nil
}

func (s *Store) GetAttendanceSummary(userID uint) (map[string]float64, error) {
	var subjects []types.Subject
	err := s.db.Where("user_id = ?", userID).Find(&subjects).Error
	if err != nil {
		return nil, fmt.Errorf("failed to fetch subjects: %v", err)
	}

	summary := make(map[string]float64)
	for _, subject := range subjects {
		if subject.MaxClasses == 0 {
			continue
		}
		percentage := (float64(subject.AttendedClasses) / float64(subject.MaxClasses)) * 100
		summary[subject.Name] = percentage
	}

	return summary, nil
}

func (s *Store) CalculateSkippableClasses(userID uint) (map[string]int, error) {
	var subjects []types.Subject
	err := s.db.Where("user_id = ?", userID).Find(&subjects).Error
	if err != nil {
		return nil, fmt.Errorf("failed to fetch subjects: %v", err)
	}

	skippable := make(map[string]int)
	for _, subject := range subjects {
		requiredAttended := int(0.8 * float64(subject.MaxClasses))
		if subject.AttendedClasses >= requiredAttended {
			skippable[subject.Name] = subject.MaxClasses - requiredAttended
		} else {
			skippable[subject.Name] = 0
		}
	}

	return skippable, nil
}

// 5️⃣ Check if Attendance is Marked
func (s *Store) IsAttendanceMarked(userID uint, subjectID uint) (bool, error) {
	today := time.Now().Truncate(24 * time.Hour)

	var attendance types.Attendance
	err := s.db.Where("user_id = ? AND subject_id = ? AND date = ?", userID, subjectID, today).First(&attendance).Error

	if err == nil {
		return true, nil
	} else if err == gorm.ErrRecordNotFound {
		return false, nil
	}

	return false, fmt.Errorf("failed to check attendance: %v", err)
}

func (s *Store) ResetAttendance(userID uint) error {
	if err := s.db.Where("user_id = ?", userID).Delete(&types.Attendance{}).Error; err != nil {
		return fmt.Errorf("failed to reset attendance: %v", err)
	}
	return nil
}
