package models

import "gorm.io/gorm"

// File represents the file metadata model in the database
type File struct {
	gorm.Model // Includes fields ID, CreatedAt, UpdatedAt, DeletedAt

	FileName string `gorm:"not null"`
	Size     int64  `gorm:"not null"`
	MimeType string `gorm:"not null"`
	S3Path   string `gorm:"unique;not null"` // Path to the file in the S3 bucket
	OwnerID  uint   `gorm:"not null"`        // The ID of the user who owns the file
}
