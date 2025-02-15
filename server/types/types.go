package types

import (
	"time"

	"gorm.io/gorm"
)

type UserStore interface {
	CreateUser(*User) error
	// GetUser(uint) (*User, error)
	// GetUserByUsername(string) (*User, error)
	// GetUserByEmail(string) (*User, error)
	// UpdateUser(*User) error
	// DeleteUser(uint) error
}

// table structure
type User struct {
	gorm.Model
	Username string `gorm:"unique" json:"username" validate:"required,min=4,max=24"`
	Email    string `gorm:"unique" json:"email" validate:"required,email"`
	Fullname string `json:"fullname" validate:"required,min=4,max=24"`
	Password string `json:"password" validate:"required"` // Hashed password
	Role     string `json:"role"`                         // student/teacher/admin
	Subjects []Subject
}

type Subject struct {
	gorm.Model
	UserID          uint   `json:"user_id" gorm:"index;constraint:OnDelete:CASCADE;"`
	Name            string `json:"name" gorm:"uniqueIndex:user_subject"`
	MaxClasses      int    `json:"max_classes"`
	TotalTaken      int    `json:"total_taken"`
	AttendedClasses int    `json:"attended_classes"`
	Schedules       []Schedule
}

type Schedule struct {
	gorm.Model
	SubjectID uint         `json:"subject_id" gorm:"index;constraint:OnDelete:CASCADE;"`
	Day       time.Weekday `json:"day"`        // 0 = Sunday, 1 = Monday, ...
	StartTime string       `json:"start_time"` // "10:00 AM"
	EndTime   string       `json:"end_time"`   // "11:00 AM"
}

type Attendance struct {
	gorm.Model
	UserID    uint      `json:"user_id" gorm:"index;constraint:OnDelete:CASCADE;"`
	SubjectID uint      `json:"subject_id" gorm:"index;constraint:OnDelete:CASCADE;"`
	Date      time.Time `json:"date" gorm:"index"`
	Status    bool      `json:"status"` // true = attended, false = missed
}
