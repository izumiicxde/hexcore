package types

import "gorm.io/gorm"

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

// type RegisterUser struct {
// 	Email    string `json:"email" validate:"required,email"`
// 	FullName string `json:"full_name" validate:"required"`
// 	Username string `json:"username" validate:"required"`
// 	Password string `json:"password" validate:"required"`
// 	Role     string `json:"role,omitempty" gorm:"default:user"`
// }
