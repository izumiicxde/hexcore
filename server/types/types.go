package types

import (
	"time"

	"gorm.io/gorm"
)

type AttendanceStore interface {
	GetTodaysClasses(userID uint) ([]ClassSchedule, error)
	GetClassesByDay(day string) ([]ClassSchedule, error)
	MarkAttendance(userID uint, subjectID uint, status bool) error
	GetAttendanceSummary(userID uint) (map[string]float64, error)
	CalculateSkippableClasses(userID uint) (map[string]int, error)
	IsAttendanceMarked(userID uint, subjectID uint) (bool, error)
	ResetAttendance(userID uint) error
}

type UserStore interface {
	CreateUser(*User) error
	GetUserByIdentifier(string) (*User, error)
	GetUserById(uint) (*User, error)
	UpdateUser(*User) error
	DeleteUser(uint) error
}

type ClassSchedule struct {
	Subject
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

type User struct {
	gorm.Model
	Username          string    `gorm:"unique" json:"username" validate:"required,min=4,max=24"`
	Email             string    `gorm:"unique" json:"email" validate:"required,email"`
	Fullname          string    `json:"fullname" validate:"required,min=4,max=24"`
	Password          string    `json:"password" validate:"required"` // Hashed password
	Role              string    `json:"role"`                         // student/teacher/admin
	IsVerified        bool      `json:"isVerified" gorm:"default:false"`
	VerificationToken string    `json:"verificationToken" gorm:"default:''"`
	TokenExpiry       time.Time `json:"tokenExpiry" swaggertype:"string" format:"date-time"`
	Subjects          []Subject `json:"subjects"`
}

type Subject struct {
	gorm.Model
	UserID          uint   `json:"userId" gorm:"index;constraint:OnDelete:CASCADE;"`
	Name            string `json:"name" gorm:"uniqueIndex:user_subject"`
	MaxClasses      int    `json:"maxClasses"`
	TotalTaken      int    `json:"totalTaken"`
	AttendedClasses int    `json:"attendedClasses"`
}

type Schedule struct {
	gorm.Model
	SubjectName string `json:"subjectName" gorm:"index"`
	Day         string `json:"day"`       // "Sunday", "Monday", etc.
	StartTime   string `json:"startTime"` // "10:00 AM"
	EndTime     string `json:"endTime"`   // "11:00 AM"
}

type Attendance struct {
	gorm.Model
	UserID    uint      `json:"userId" gorm:"index;constraint:OnDelete:CASCADE;"`
	SubjectID uint      `json:"subjectId" gorm:"index;constraint:OnDelete:CASCADE;"`
	Date      time.Time `json:"date" gorm:"index" swaggertype:"string" format:"date"`
	Status    bool      `json:"status"`
}
