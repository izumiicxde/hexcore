package attendance

import (
	"hexcore/config"
	"hexcore/types"

	"gorm.io/gorm"
)

type Store struct {
	db *gorm.DB
}

func NewAttendanceStore(db *gorm.DB) *Store {
	return &Store{db}
}

func (s *Store) MarkAttendance(userId int, req *types.AttendanceRequest) error {
	if err := config.Validator.Struct(req); err != nil {
		return err
	}

	// Find the Subject ID for this user
	subject := new(types.Subject)
	if err := s.db.Where("user_id = ? AND name = ?", userId, req.SubjectName).First(subject).Error; err != nil {
		return err
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

func (s *Store) GetAttendaceSummary(userId int) ([]types.Attendance, error) {
	var attendances []types.Attendance
	if err := s.db.Where("user_id = ?", userId).Find(&attendances).Error; err != nil {
		return nil, err
	}
	return attendances, nil
}
