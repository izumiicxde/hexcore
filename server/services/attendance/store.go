package attendance

import (
	"hexcore/types"

	"gorm.io/gorm"
)

type AttendanceStore struct {
	db *gorm.DB
}

func NewAttendanceStore(db *gorm.DB) *AttendanceStore {
	return &AttendanceStore{db}
}

// Get a subject by userID & subject name
func (s *AttendanceStore) GetSubject(userID uint, subjectName string) (*types.Subject, error) {
	subject := new(types.Subject)
	if err := s.db.Where("user_id = ? AND name = ?", userID, subjectName).First(subject).Error; err != nil {
		return nil, err
	}
	return subject, nil
}

// Update attendance
func (s *AttendanceStore) UpdateAttendance(userID uint, subjectName string, attended bool) (*types.Subject, error) {
	subject, err := s.GetSubject(userID, subjectName)
	if err != nil {
		return nil, err
	}

	subject.TotalTaken++
	if attended {
		subject.AttendedClasses++
	}

	if err := s.db.Save(subject).Error; err != nil {
		return nil, err
	}

	return subject, nil
}

// Get all subjects for a user
func (s *AttendanceStore) GetUserSubjects(userID uint) ([]types.Subject, error) {
	var subjects []types.Subject
	if err := s.db.Where("user_id = ?", userID).Find(&subjects).Error; err != nil {
		return nil, err
	}
	return subjects, nil
}
