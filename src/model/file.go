package model

import "time"

type FileUpload struct {
	ID         string    `gorm:"primaryKey" json:"id"`
	FileName   string    `json:"fileName"`
	FileSize   int64     `json:"fileSize"`
	FileType   string    `json:"fileType"`
	FileURL    string    `json:"fileURL"`
	UploadedBy string    `json:"uploadedBy"`
	Analysis   string    `json:"analysis"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
