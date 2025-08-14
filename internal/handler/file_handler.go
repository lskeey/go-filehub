package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lskeey/go-filehub/internal/service"
)

type FileHandler struct {
	fileService *service.FileService
}

func NewFileHandler(s *service.FileService) *FileHandler {
	return &FileHandler{fileService: s}
}

// UploadFile handles the file upload request.
//
// @Summary Upload a file
// @Description Uploads a file for the authenticated user. Max file size is 10MB.
// @Tags files
// @Accept  multipart/form-data
// @Produce  json
// @Param   file  formData  file  true  "File to upload"
// @Success 200   {object}  UploadSuccessResponse
// @Failure 400   {object}  ErrorResponse
// @Failure 401   {object}  ErrorResponse
// @Failure 500   {object}  ErrorResponse
// @Security BearerAuth
// @Router /files/upload [post]
func (h *FileHandler) UploadFile(c *gin.Context) {
	// Retrieve the file from the form data
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	}

	// Retrieve userID from the context (set by the AuthMiddleware)
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID not found in context"})
		return
	}

	// Here you can add validation for file size and type
	// For example: max size 10MB
	const maxFileSize = 10 * 1024 * 1024
	if fileHeader.Size > maxFileSize {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File size exceeds the limit of 10MB"})
		return
	}

	// Call the service to handle the file upload
	fileMetadata, err := h.fileService.UploadFile(c, fileHeader, userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "File uploaded successfully",
		"data":    fileMetadata,
	})
}

// ListFiles handles listing all files for the authenticated user.
//
// @Summary List user's files
// @Description Retrieves a list of all files uploaded by the authenticated user.
// @Tags files
// @Produce  json
// @Success 200   {object}  ListFilesResponse
// @Failure 401   {object}  ErrorResponse
// @Failure 500   {object}  ErrorResponse
// @Security BearerAuth
// @Router /files [get]
func (h *FileHandler) ListFiles(c *gin.Context) {
	userID, _ := c.Get("userID")

	files, err := h.fileService.ListUserFiles(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve files"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": files})
}

// DownloadFile handles serving a specific file for download.
//
// @Summary Download a file
// @Description Downloads a specific file by its ID. The user must own the file.
// @Tags files
// @Produce  application/octet-stream
// @Param   id    path      int  true  "File ID"
// @Success 200   {file}    file
// @Failure 400   {object}  ErrorResponse
// @Failure 403   {object}  ErrorResponse
// @Failure 404   {object}  ErrorResponse
// @Security BearerAuth
// @Router /files/{id}/download [get]
func (h *FileHandler) DownloadFile(c *gin.Context) {
	userID, _ := c.Get("userID")

	fileID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file ID"})
		return
	}

	file, err := h.fileService.GetFileByID(uint(fileID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	// Authorization check
	if file.OwnerID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to download this file"})
		return
	}

	// Serve the file for download
	c.FileAttachment(file.S3Path, file.FileName)
}

// DeleteFile handles the deletion of a specific file.
//
// @Summary Delete a file
// @Description Deletes a specific file by its ID. The user must own the file.
// @Tags files
// @Produce  json
// @Param   id    path      int  true  "File ID"
// @Success 200   {object}  SuccessResponse
// @Failure 400   {object}  ErrorResponse
// @Failure 403   {object}  ErrorResponse
// @Failure 404   {object}  ErrorResponse
// @Security BearerAuth
// @Router /files/{id} [delete]
func (h *FileHandler) DeleteFile(c *gin.Context) {
	userID, _ := c.Get("userID")

	fileID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file ID"})
		return
	}

	err = h.fileService.DeleteFile(uint(fileID), userID.(uint))
	if err != nil {
		if err.Error() == "unauthorized: you do not own this file" {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File deleted successfully"})
}
