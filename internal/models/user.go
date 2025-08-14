package models

import "gorm.io/gorm"

// User represents the user model in the database
type User struct {
	gorm.Model // Includes fields ID, CreatedAt, UpdatedAt, DeletedAt

	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Files    []File `gorm:"foreignKey:OwnerID"` // A user can have many files
}
