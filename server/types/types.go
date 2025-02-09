package types

import (
	"time"

	"gorm.io/gorm"
)

type AttendanceStore interface {
	MarkAttendance(int, *AttendanceRequest) error
	GetAttendaceSummary(userId int) ([]Attendance, error)
}
type AttendanceRequest struct {
	SubjectName string    `json:"subject_name" validate:"required"`
	Status      bool      `json:"status" validate:"required"`
	Date        time.Time `json:"date" validate:"required"`
}

type SubjectSchedule struct {
	gorm.Model
	Name      string `json:"name" gorm:"index"`                  // Matches the Subject.Name
	Day       string `json:"day" gorm:"index"`                   // "Monday", "Tuesday", etc.
	StartTime string `json:"start_time" gorm:"type:varchar(10)"` // Example: "10:00 AM"
	EndTime   string `json:"end_time" gorm:"type:varchar(10)"`   // Example: "11:00 AM"
}

type Attendance struct {
	gorm.Model
	UserId    uint      `gorm:"index;constraint:OnDelete:CASCADE;" validate:"required" json:"user_id"`
	SubjectId uint      `gorm:"index;constraint:OnDelete:CASCADE;" validate:"required" json:"subject_id"`
	Status    bool      `json:"status" validate:"required"`
	Date      time.Time `json:"date"  gorm:"index;uniqueIndex:user_subject_date" validate:"required"`

	// Relations
	User    User    `gorm:"foreignKey:UserId"`
	Subject Subject `gorm:"foreignKey:SubjectId"`
}

type Subject struct {
	gorm.Model
	UserId          uint   `json:"user_id" gorm:"index;constraint:OnDelete:CASCADE;"`
	Name            string `json:"name" gorm:"uniqueIndex:user_subject"` // Unique per user
	MaxClasses      int    `json:"max_classes"`
	TotalTaken      int    `json:"total_taken"`
	AttendedClasses int    `json:"attended_classes"`
}

type UserStore interface {
	CreateUser(user *User) error
	GetUserByUsername(username string) (*User, error)
	GetAllUsers() ([]User, error)
	GetUserById(id int) (*User, error)
	UpdateUser(user *User) error
	DeleteUser(id int) error
}

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"unique" validate:"required,email"`
	FullName string `json:"full_name" validate:"required"`
	Password string `json:"password" validate:"required"`
	Role     string `json:"role" validate:"required"`
	Username string `json:"username" gorm:"unique" validate:"required"`
}
