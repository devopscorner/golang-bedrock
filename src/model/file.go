package model

import "time"

type FileUpload struct {
	ID         string `gorm:"primaryKey"`
	FileName   string `json:"fileName"`
	FileSize   int64  `json:"fileSize"`
	FileType   string `json:"fileType"`
	FileURL    string `json:"fileURL"`
	UploadedBy string `json:"uploadedBy"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
