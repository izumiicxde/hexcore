package user

import (
	"errors"
	"hexcore/config"
	"hexcore/types"
	"hexcore/utils"

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

func (s *Store) CreateUser(user *types.User) error {
	if err := config.Validator.Struct(user); err != nil {
		return err
	}

	tx := s.db.Begin()

	if err := tx.Create(user).Error; err != nil {
		tx.Rollback()
		return err
	}

	var schedules []types.SubjectSchedule
	if err := tx.Find(&schedules).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Create a map of subject name -> MaxClasses from utils.Subjects
	subjectMaxClasses := make(map[string]int)
	for _, sub := range utils.Subjects {
		subjectMaxClasses[sub.Name] = sub.MaxClasses
	}

	// Deduplicate and assign max classes
	subjectMap := make(map[string]bool)
	var subjects []types.Subject

	for _, schedule := range schedules {
		if !subjectMap[schedule.Name] {
			subjectMap[schedule.Name] = true
			subjects = append(subjects, types.Subject{
				UserId:     user.ID,
				Name:       schedule.Name,
				MaxClasses: subjectMaxClasses[schedule.Name], // Get calculated MaxClasses
			})
		}
	}

	if err := tx.Create(&subjects).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
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

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	user := new(types.User)
	if err := s.db.Where("email = ?", email).First(user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return user, nil
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
