package types

import (
	"gorm.io/gorm"
)

type Subject struct {
	gorm.Model
	UserId          uint   `json:"user_id,omitempty" gorm:"index"`
	Name            string `json:"name,omitempty" gorm:"uniqueIndex:student_subject"`
	MaxClasses      int    `json:"max_classes,omitempty"`      // Total planned classes (e.g., 80)
	TotalTaken      int    `json:"total_taken,omitempty"`      // How many classes have been conducted so far
	AttendedClasses int    `json:"attended_classes,omitempty"` // How many the student attended
}

type AttendanceStore interface {
	GetSubject(userID uint, subjectName string) (*Subject, error)
	UpdateAttendance(userID uint, subjectName string, attended bool) (*Subject, error)
	GetUserSubjects(userID uint) ([]Subject, error)
}

type UserStore interface {
	CreateUser(user *User) error
	GetUserByUsername(username string) (*User, error)
	GetAllUsers() ([]User, error)
	GetUserById(id int) (*User, error)
	UpdateUser(user *User) error
	DeleteUser(id int) error

	InitializeTable(userId uint) error
}

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"unique" validate:"required,email"`
	FullName string `json:"full_name" validate:"required"`
	Password string `json:"password" validate:"required"`
	Role     string `json:"role" validate:"required"`
	Username string `json:"username" gorm:"unique" validate:"required"`
}
