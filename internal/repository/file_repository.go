package repository

import (
	"github.com/lskeey/go-filehub/internal/models"
	"gorm.io/gorm"
)

type FileRepository struct {
	DB *gorm.DB
}

func NewFileRepository(db *gorm.DB) *FileRepository {
	return &FileRepository{DB: db}
}

// CreateFile saves file metadata to the database.
func (r *FileRepository) CreateFile(file *models.File) error {
	return r.DB.Create(file).Error
}

// FindFilesByOwnerID retrieves all files owned by a specific user.
func (r *FileRepository) FindFilesByOwnerID(userID uint) ([]models.File, error) {
	var files []models.File
	err := r.DB.Where("owner_id = ?", userID).Find(&files).Error
	return files, err
}

// FindFileByID retrieves a single file by its ID.
func (r *FileRepository) FindFileByID(fileID uint) (*models.File, error) {
	var file models.File
	err := r.DB.First(&file, fileID).Error
	return &file, err
}

// DeleteFileByID removes a file record from the database.
// Note: GORM uses soft deletes by default if the model has gorm.DeletedAt.
// To permanently delete, use Unscoped().Delete()
func (r *FileRepository) DeleteFileByID(fileID uint) error {
	return r.DB.Delete(&models.File{}, fileID).Error
}
