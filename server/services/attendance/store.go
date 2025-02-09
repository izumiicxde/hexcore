package attendance

import (
	"fmt"
	"hexcore/config"
	"hexcore/types"
	"math"

	"gorm.io/gorm"
)

type Store struct {
	db *gorm.DB
}

func NewAttendanceStore(db *gorm.DB) *Store {
	return &Store{db}
}

func (s *Store) GetAttendancePrediction(userId int) ([]types.AttendancePrediction, error) {
	var subjects []types.Subject
	if err := s.db.Where("user_id = ?", userId).Find(&subjects).Error; err != nil {
		return nil, err
	}

	var predictions []types.AttendancePrediction
	for _, subject := range subjects {
		remainingClasses := subject.MaxClasses - subject.TotalTaken
		requiredAttended := int(math.Ceil(0.80 * float64(subject.MaxClasses)))
		canSkip := subject.MaxClasses - requiredAttended - subject.AttendedClasses

		if canSkip < 0 {
			canSkip = 0 // User already below 75%
		}

		predictions = append(predictions, types.AttendancePrediction{
			SubjectName:      subject.Name,
			CanSkip:          canSkip,
			RemainingClasses: remainingClasses,
			ClassesTaken:     subject.TotalTaken,
			TotalClasses:     subject.MaxClasses,
		})
	}
	return predictions, nil
}

func (s *Store) MarkAttendance(userId int, req *types.AttendanceRequest) error {
	if err := config.Validator.Struct(req); err != nil {
		return err
	}

	// Find the Subject ID for this user
	subject := new(types.Subject)
	if err := s.db.Where("user_id = ? AND name = ?", userId, req.SubjectName).First(subject).Error; err != nil {
		return fmt.Errorf("this is the error? %v", err)
	}

	// Insert Attendance
	attendance := types.Attendance{
		UserId:    uint(userId),
		SubjectId: subject.ID,
		Status:    req.Status,
		Date:      req.Date,
	}

	if err := s.db.Create(&attendance).Error; err != nil {
		return err
	}

	// Update Subject Stats
	if req.Status {
		subject.AttendedClasses++
	}
	subject.TotalTaken++
	s.db.Save(&subject)

	return nil
}

func (s *Store) GetAttendanceSummary(userId int) ([]types.Attendance, error) {
	var attendances []types.Attendance
	if err := s.db.
		Preload("User").
		Preload("Subject").
		Where("user_id = ?", userId).
		Find(&attendances).Error; err != nil {
		return nil, err
	}
	return attendances, nil
}
