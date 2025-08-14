package handler

type FileResponse struct {
	ID       uint   `json:"id"`
	FileName string `json:"file_name"`
	Size     int64  `json:"size"`
	MimeType string `json:"mime_type"`
	S3Path   string `json:"s3_path"`
	OwnerID  uint   `json:"owner_id"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type UploadSuccessResponse struct {
	Message string       `json:"message"`
	Data    FileResponse `json:"data"`
}

type ListFilesResponse struct {
	Data []FileResponse `json:"data"`
}
