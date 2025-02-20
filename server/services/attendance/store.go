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

func (s *Store) GetClassesTillToday(userID uint) ([]types.ClassSchedule, error) {
	var classes []types.ClassSchedule

	err := s.db.
		Table("schedules").
		Select("subjects.*, schedules.start_time, schedules.end_time, COALESCE(attendances.status, false) AS status").
		Joins("JOIN subjects ON schedules.subject_name = subjects.name").
		Joins("LEFT JOIN attendances ON attendances.subject_id = subjects.id AND DATE(attendances.date) <= CURRENT_DATE").
		Where("subjects.user_id = ?", userID).
		Order("schedules.start_time ASC").
		Scan(&classes).Error

	if err != nil {
		return nil, fmt.Errorf("failed to fetch classes: %v", err)
	}

	return classes, nil
}

func (s *Store) GetTodaysClasses(userID uint) ([]types.ClassSchedule, error) {
	today := time.Now().Weekday().String()

	classes, err := s.GetClassesByDay(today, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch classes for %s: %v", today, err)
	}
	return classes, nil
}

// GetClassesByDay fetches the subjects for a given day provided in query params.
func (s *Store) GetClassesByDay(day string, userID uint) ([]types.ClassSchedule, error) {
	var classes []types.ClassSchedule

	// Validate the day input
	validDays := map[string]bool{
		"Monday": true, "Tuesday": true, "Wednesday": true, "Thursday": true,
		"Friday": true, "Saturday": true, "Sunday": true,
	}
	if !validDays[day] {
		return nil, fmt.Errorf("invalid day provided: %s", day)
	}

	// Query schedules, subjects, and attendance status for today
	err := s.db.
		Table("schedules").
		Select("subjects.*, schedules.start_time, schedules.end_time, COALESCE(attendances.status, false) AS status").
		Joins("JOIN subjects ON schedules.subject_name = subjects.name").
		Joins("LEFT JOIN attendances ON attendances.subject_id = subjects.id AND DATE(attendances.date) = CURRENT_DATE").
		Where("schedules.day = ? AND subjects.user_id = ?", day, userID).
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

	// Update both total_taken and attended_classes
	updateQuery := s.db.Model(&types.Subject{}).Where("id = ?", subjectID).
		Update("total_taken", gorm.Expr("total_taken + 1"))

	if status {
		updateQuery = updateQuery.Update("attended_classes", gorm.Expr("attended_classes + 1"))
	}

	if err := updateQuery.Error; err != nil {
		return fmt.Errorf("failed to update subject attendance: %v", err)
	}

	return nil
}

func (s *Store) GetAttendanceSummary(userID uint) (map[string]interface{}, error) {
	var subjects []types.Subject
	err := s.db.Where("user_id = ?", userID).Find(&subjects).Error
	if err != nil {
		return nil, fmt.Errorf("failed to fetch subjects: %v", err)
	}

	totalClasses := 0
	totalAttended := 0
	summary := make(map[string]interface{})

	subjectDetails := make(map[string]map[string]float64)

	for _, subject := range subjects {
		if subject.MaxClasses == 0 {
			continue
		}

		attended := float64(subject.AttendedClasses)
		maxClasses := float64(subject.MaxClasses)
		percentage := (attended / maxClasses) * 100

		subjectDetails[subject.Name] = map[string]float64{
			"max_classes":      maxClasses,
			"attended_classes": attended,
			"remaining":        maxClasses - attended,
			"percentage":       percentage,
		}

		totalClasses += subject.MaxClasses
		totalAttended += subject.AttendedClasses
	}

	overallPercentage := 0.0
	if totalClasses > 0 {
		overallPercentage = (float64(totalAttended) / float64(totalClasses)) * 100
	}

	summary["subjects"] = subjectDetails
	summary["total_classes"] = totalClasses
	summary["total_attended"] = totalAttended
	summary["total_missed"] = totalClasses - totalAttended
	summary["overall_percentage"] = overallPercentage

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
