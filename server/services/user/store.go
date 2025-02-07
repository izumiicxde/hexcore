package user

import (
	"errors"
	"hexcore/config"
	"hexcore/types"

	"gorm.io/gorm"
)

// Store struct implements UserStore interface
type Store struct {
	db *gorm.DB
}

// NewStore initializes a new user store
func NewStore(db *gorm.DB) *Store {
	return &Store{db: db}
}

// CreateUser inserts a new user into the database with validation
func (s *Store) CreateUser(user *types.User) error {
	// Validate user struct before inserting
	if err := config.Validator.Struct(user); err != nil {
		return err
	}

	return s.db.Create(user).Error
}

// GetUserByUsername retrieves a user by username
func (s *Store) GetUserByUsername(username string) (*types.User, error) {
	var user types.User
	if err := s.db.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (s *Store) GetAllUsers() ([]types.User, error) {
	var users []types.User

	if err := s.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// GetUserById retrieves a user by ID
func (s *Store) GetUserById(id int) (*types.User, error) {
	var user types.User
	if err := s.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

// UpdateUser updates a user's details with validation
func (s *Store) UpdateUser(user *types.User) error {
	// Validate updated user data
	if err := config.Validator.Struct(user); err != nil {
		return err
	}

	// Ensure the user exists before updating
	existingUser, err := s.GetUserById(int(user.ID))
	if err != nil {
		return err
	}

	return s.db.Model(existingUser).Updates(user).Error
}

func (s *Store) DeleteUser(id int) error {
	var user types.User

	// Fetch user including soft deleted ones
	if err := s.db.Unscoped().First(&user, id).Error; err != nil {
		return errors.New("user not found")
	}

	// If user is already soft deleted, permanently delete
	if !user.DeletedAt.Time.IsZero() { // Check if deleted_at is set
		return s.db.Unscoped().Delete(&user).Error
	}

	// Otherwise, soft delete the user
	return s.db.Delete(&user).Error
}
