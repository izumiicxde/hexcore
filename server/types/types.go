package types

import (
	"time"

	"gorm.io/gorm"
)

type AttendanceStore interface {
	GetTodaysClasses(userID uint) ([]Subject, error)
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

// User represents a user in the system.
//
// @Description User model for authentication and attendance tracking.
// @property {string} username - Unique username (min: 4, max: 24).
// @property {string} email - Unique email address.
// @property {string} fullname - Full name of the user (min: 4, max: 24).
// @property {string} password - Hashed password.
// @property {string} role - Role of the user (student/teacher/admin).
// @property {boolean} isVerified - Whether the user has verified their email.
// @property {string} verificationToken - Token for email verification.
// @property {string} tokenExpiry - Expiration time for the verification token.
// @property {array} subjects - List of subjects assigned to the user.
type User struct {
	// swaggerignore
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

// Subject represents a subject with attendance tracking.
//
// @Description Subject model that contains user assignments and attendance tracking.
// @property {integer} userId - The ID of the user who owns the subject.
// @property {string} name - The name of the subject (unique per user).
// @property {integer} maxClasses - Maximum number of classes.
// @property {integer} totalTaken - Total number of classes conducted.
// @property {integer} attendedClasses - Number of attended classes.
type Subject struct {
	// swaggerignore
	gorm.Model

	UserID          uint   `json:"userId" gorm:"index;constraint:OnDelete:CASCADE;"`
	Name            string `json:"name" gorm:"uniqueIndex:user_subject"`
	MaxClasses      int    `json:"maxClasses"`
	TotalTaken      int    `json:"totalTaken"`
	AttendedClasses int    `json:"attendedClasses"`
}

// Schedule represents a class schedule for a subject.
//
// @Description Schedule model defining class timings for a subject.
// @property {string} subjectName - The name of the subject.
// @property {string} day - The day of the week (e.g., "Monday").
// @property {string} startTime - Class start time (HH:MM AM/PM).
// @property {string} endTime - Class end time (HH:MM AM/PM).
type Schedule struct {
	// swaggerignore
	gorm.Model

	SubjectName string `json:"subjectName" gorm:"index"`
	Day         string `json:"day"`       // "Sunday", "Monday", etc.
	StartTime   string `json:"startTime"` // "10:00 AM"
	EndTime     string `json:"endTime"`   // "11:00 AM"
}

// Attendance represents an attendance record.
//
// @Description Attendance tracking for each subject.
// @property {integer} userId - The ID of the user who attended the class.
// @property {integer} subjectId - The ID of the subject.
// @property {string} date - The date of attendance.
// @property {boolean} status - True if attended, false if missed.
type Attendance struct {
	// swaggerignore
	gorm.Model

	UserID    uint      `json:"userId" gorm:"index;constraint:OnDelete:CASCADE;"`
	SubjectID uint      `json:"subjectId" gorm:"index;constraint:OnDelete:CASCADE;"`
	Date      time.Time `json:"date" gorm:"index" swaggertype:"string" format:"date"`
	Status    bool      `json:"status"`
}
