// cmd/migrate_file_upload.go
package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/devopscorner/golang-bedrock/src/config"
	"github.com/devopscorner/golang-bedrock/src/driver"
	"github.com/devopscorner/golang-bedrock/src/model"
	"gorm.io/gorm"
)

func main() {
	_, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Connect to database
	driver.ConnectDatabase()

	MigrateUploader(driver.DB)
	fmt.Println("Migration Done...")
}

func MigrateUploader(db *gorm.DB) error {
	if err := db.AutoMigrate(&model.FileUpload{}); err != nil {
		return err
	}

	files := []model.FileUpload{
		{FileName: "document1.pdf", FileSize: 1024 * 1024, FileType: "application/pdf", UploadedBy: "user1@example.com"},
		{FileName: "image1.jpg", FileSize: 512 * 1024, FileType: "image/jpeg", UploadedBy: "user2@example.com"},
		{FileName: "spreadsheet1.xlsx", FileSize: 2048 * 1024, FileType: "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", UploadedBy: "user3@example.com"},
		{FileName: "presentation1.pptx", FileSize: 3072 * 1024, FileType: "application/vnd.openxmlformats-officedocument.presentationml.presentation", UploadedBy: "user1@example.com"},
		{FileName: "document2.docx", FileSize: 1536 * 1024, FileType: "application/vnd.openxmlformats-officedocument.wordprocessingml.document", UploadedBy: "user2@example.com"},
	}

	for _, file := range files {
		t := time.Now().UnixNano()
		id_time := strconv.FormatInt(t, 10)
		file.ID = id_time
		file.CreatedAt = time.Now()
		file.UpdatedAt = time.Now()
		file.FileURL = fmt.Sprintf("https://example-bucket.s3.amazonaws.com/%s", file.FileName) // Example S3 URL
		if err := db.Create(&file).Error; err != nil {
			fmt.Printf("Failed to insert data for file: %+v\n", file)
			fmt.Printf("Error: %v\n", err)
			return err
		} else {
			fmt.Printf("Sample file: %+v, created successfully! \n", file.FileName)
		}
	}

	return nil
}
