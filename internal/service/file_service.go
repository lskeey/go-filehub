package service

import (
	"errors"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lskeey/go-filehub/internal/models"
	"github.com/lskeey/go-filehub/internal/repository"
)

type FileService struct {
	fileRepo *repository.FileRepository
}

func NewFileService(repo *repository.FileRepository) *FileService {
	return &FileService{fileRepo: repo}
}

// UploadFile handles the business logic of uploading a file.
func (s *FileService) UploadFile(c *gin.Context, fileHeader *multipart.FileHeader, userID uint) (*models.File, error) {
	// Generate a unique file name to prevent collisions
	// Format: <userID>-<timestamp>-<original_filename>
	uniqueFileName := fmt.Sprintf("%d-%d-%s", userID, time.Now().Unix(), fileHeader.Filename)

	// Define the path to save the file
	savePath := filepath.Join("uploads", uniqueFileName)

	// Save the file to the local server
	if err := c.SaveUploadedFile(fileHeader, savePath); err != nil {
		return nil, err
	}

	// Create a record for the database
	fileMetadata := &models.File{
		FileName: fileHeader.Filename,
		Size:     fileHeader.Size,
		MimeType: fileHeader.Header.Get("Content-Type"),
		S3Path:   savePath, // For now, this is the local path
		OwnerID:  userID,
	}

	// Save metadata to the database
	if err := s.fileRepo.CreateFile(fileMetadata); err != nil {
		// Here you might want to add logic to delete the saved file if DB insert fails
		return nil, err
	}

	return fileMetadata, nil
}

// ListUserFiles retrieves all files for a given user.
func (s *FileService) ListUserFiles(userID uint) ([]models.File, error) {
	return s.fileRepo.FindFilesByOwnerID(userID)
}

// GetFileByID retrieves a single file record.
func (s *FileService) GetFileByID(fileID uint) (*models.File, error) {
	return s.fileRepo.FindFileByID(fileID)
}

// DeleteFile handles the logic for deleting a file.
func (s *FileService) DeleteFile(fileID, userID uint) error {
	// 1. Get file metadata to verify ownership and get the path
	file, err := s.fileRepo.FindFileByID(fileID)
	if err != nil {
		return errors.New("file not found")
	}

	// 2. IMPORTANT: Check if the user owns the file
	if file.OwnerID != userID {
		return errors.New("unauthorized: you do not own this file")
	}

	// 3. Delete the physical file from storage
	if err := os.Remove(file.S3Path); err != nil {
		// Log the error but you might still want to proceed to delete the DB record
		// depending on desired behavior for orphaned records.
		fmt.Printf("Failed to delete physical file %s: %v\n", file.S3Path, err)
	}

	// 4. Delete the metadata from the database
	return s.fileRepo.DeleteFileByID(fileID)
}
