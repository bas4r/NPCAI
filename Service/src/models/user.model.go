package models

import (
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name      string    `gorm:"type:varchar(255);not null"`
	Email     string    `gorm:"uniqueIndex;not null"`
	Password  string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

type NewUserInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required,min=8"`
}

type LogInInput struct {
	Email    string `json:"email"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}

// Validate validates the new user data
func (u *NewUserInput) Validate() error {
	// Validate the username
	if u.Name == "" {
		return fmt.Errorf("name is required")
	}

	// Validate the password
	if u.Password == "" {
		return fmt.Errorf("password is required")
	}

	// Validate the email
	if u.Email == "" {
		return fmt.Errorf("email is required")
	}

	return nil
}

// Save saves the new user to the database
func (u *NewUserInput) Save(db *gorm.DB) error {
	// Connect to the database

	// Create the user
	user := &User{
		ID:        uuid.New(),
		Name:      u.Name,
		Email:     u.Email,
		Password:  u.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	result := db.Create(&user)

	if result.Error != nil {
		log.Fatal(result.Error)
		return result.Error
	}

	return nil
}
