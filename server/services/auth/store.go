package auth

import (
	"errors"
	"fmt"
	"hexcore/config"
	"hexcore/types"
	"strings"

	"gorm.io/gorm"
)

type Store struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateUser(user *types.User) error {
	if err := config.Validator.Struct(user); err != nil {
		return err
	}
	// Start transaction
	tx := s.db.Begin()

	// Create user
	if err := tx.Create(user).Error; err != nil {
		tx.Rollback()
		if strings.Contains(err.Error(), "duplicate key value") {
			return fmt.Errorf("user with this email or username already exists")
		}
		return err
	}

	// Create subjects for the user
	for _, subject := range config.Subjects {
		newSubject := types.Subject{
			UserID:     user.ID,
			Name:       subject.Name,
			MaxClasses: subject.MaxClasses,
		}
		if err := tx.Create(&newSubject).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	// Commit transaction if everything is successful
	return tx.Commit().Error
}

func (s *Store) GetUserByIdentifier(identifier string) (*types.User, error) {
	user := new(types.User)
	if err := s.db.Where("email = ? OR username = ?", identifier, identifier).First(user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	return user, nil
}

func (s *Store) GetUserById(id uint) (*types.User, error) {
	user := new(types.User)

	if err := s.db.Where("id = ?", id).First(user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}
	return user, nil
}

func (s *Store) UpdateUser(user *types.User) error {
	tx := s.db.Begin()

	currentUser := new(types.User)
	if err := tx.Where("id = ?", user.ID).First(currentUser).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("user not found %v", err)
	}
	if err := tx.Model(currentUser).Updates(user).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("error updating user %v", err)
	}
	return tx.Commit().Error
}

func (s *Store) DeleteUser(id uint) error {
	user := new(types.User)
	if err := s.db.Unscoped().Where("id = ?", id).First(user).Error; err != nil {
		return fmt.Errorf("user not found: %v", err)
	}

	tx := s.db.Begin()

	if user.DeletedAt.Valid {
		// User already soft deleted — delete permanently
		if err := tx.Unscoped().Where("user_id = ?", user.ID).Delete(&types.Subject{}).Error; err != nil {
			tx.Rollback()
			return err
		}

		if err := tx.Unscoped().Delete(user).Error; err != nil {
			tx.Rollback()
			return err
		}
	} else {
		// Soft delete the user
		if err := tx.Delete(user).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}
