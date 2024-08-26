// controller/file_controller.go
package controller

import (
	"fmt"
	"mime/multipart"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/devopscorner/golang-restfulapi-bedrock/src/config"
	"github.com/devopscorner/golang-restfulapi-bedrock/src/model"
	"github.com/devopscorner/golang-restfulapi-bedrock/src/repository"
	"github.com/devopscorner/golang-restfulapi-bedrock/src/utility"
	"github.com/devopscorner/golang-restfulapi-bedrock/src/view"
	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

var fileTracer = otel.Tracer("file-controller")

type FileController struct {
	repo     repository.FileRepository
	s3Client *s3.Client
}

func NewFileController(repo repository.FileRepository, s3Client *s3.Client) *FileController {
	return &FileController{repo: repo, s3Client: s3Client}
}

func (fc *FileController) CreateFile(c *gin.Context) {
	ctxTrace, span := fileTracer.Start(c.Request.Context(), "CreateFile")
	defer span.End()

	file, err := c.FormFile("file")
	if err != nil {
		utility.RecordError(ctxTrace, err)
		utility.LogError(c, "Invalid file upload", err)
		view.ErrorBadRequest(c, err)
		return
	}

	filename := generateUniqueFilename(file.Filename)

	s3URL, err := fc.uploadFileToS3(c, file, filename)
	if err != nil {
		utility.RecordError(ctxTrace, err)
		utility.LogError(c, "Failed to upload file to S3", err)
		view.ErrorInternalServer(c, err)
		return
	}

	fileUpload := model.FileUpload{
		FileName:   filename,
		FileSize:   file.Size,
		FileType:   file.Header.Get("Content-Type"),
		FileURL:    s3URL,
		UploadedBy: c.GetString("user"),
	}

	validate := validator.New()
	if err := validate.Struct(fileUpload); err != nil {
		utility.RecordError(ctxTrace, err)
		utility.LogError(c, "File validation failed", err)
		view.ErrorBadRequest(c, err)
		return
	}

	err = fc.repo.CreateFile(ctxTrace, &fileUpload)
	if err != nil {
		utility.RecordError(ctxTrace, err)
		utility.LogError(c, "Failed to create file record", err)
		view.ErrorInternalServer(c, err)
		return
	}

	utility.RecordFileUpload(fileUpload.FileType, float64(fileUpload.FileSize))

	go fc.analyzeUploadWithBedrock(c, &fileUpload)

	span.SetAttributes(
		attribute.String("filename", fileUpload.FileName),
		attribute.Int64("file_size", fileUpload.FileSize),
		attribute.String("file_type", fileUpload.FileType),
	)

	utility.LogInfo(c, fmt.Sprintf("Created file: %s, Size: %d, Type: %s", fileUpload.FileName, fileUpload.FileSize, fileUpload.FileType))
	view.ViewCreateFile(c, fileUpload)
}

func (fc *FileController) FindAll(c *gin.Context) {
	ctxTrace, span := fileTracer.Start(c.Request.Context(), "FindAll")
	defer span.End()

	files, err := fc.repo.FindAll(ctxTrace)
	if err != nil {
		utility.RecordError(ctxTrace, err)
		utility.LogError(c, "Failed to find files", err)
		view.ErrorInternalServer(c, err)
		return
	}

	span.SetAttributes(attribute.Int("file_count", len(files)))

	utility.LogInfo(c, fmt.Sprintf("Retrieved %d files", len(files)))
	view.ViewFindAllFiles(c, files)
}

func (fc *FileController) FindByID(c *gin.Context) {
	ctxTrace, span := fileTracer.Start(c.Request.Context(), "FindByID")
	defer span.End()

	id := c.Param("id")
	span.SetAttributes(attribute.String("file_id", id))

	file, err := fc.repo.FindByID(ctxTrace, id)
	if err != nil {
		utility.RecordError(ctxTrace, err)
		utility.LogError(c, fmt.Sprintf("Failed to find file: %s", id), err)
		view.ErrorInternalServer(c, err)
		return
	}
	if file == nil {
		utility.LogWarn(c, fmt.Sprintf("File not found: %s", id))
		view.ErrorNotFound(c)
		return
	}

	utility.LogInfo(c, fmt.Sprintf("Retrieved file: %s", id))
	view.ViewFindFileByID(c, file)
}

func (fc *FileController) UpdateFile(c *gin.Context) {
	ctxTrace, span := fileTracer.Start(c.Request.Context(), "UpdateFile")
	defer span.End()

	id := c.Param("id")
	span.SetAttributes(attribute.String("file_id", id))

	var input model.FileUpload
	if err := c.ShouldBindJSON(&input); err != nil {
		utility.RecordError(ctxTrace, err)
		utility.LogError(c, "Invalid input for file update", err)
		view.ErrorBadRequest(c, err)
		return
	}

	file, err := fc.repo.FindByID(c, id)
	if err != nil {
		utility.RecordError(ctxTrace, err)
		utility.LogError(c, fmt.Sprintf("Failed to find file for update: %s", id), err)
		view.ErrorNotFound(c)
		return
	}

	file.FileName = input.FileName
	file.FileSize = input.FileSize
	file.FileType = input.FileType
	file.FileURL = input.FileURL
	file.UploadedBy = input.UploadedBy

	if err := fc.repo.UpdateFile(c, file); err != nil {
		utility.RecordError(ctxTrace, err)
		utility.LogError(c, fmt.Sprintf("Failed to update file: %s", id), err)
		view.ErrorInternalServer(c, err)
		return
	}

	utility.LogInfo(c, fmt.Sprintf("Updated file: %s", id))
	view.ViewUpdateFile(c, file)
}

func (fc *FileController) DeleteFile(c *gin.Context) {
	ctxTrace, span := fileTracer.Start(c.Request.Context(), "DeleteFile")
	defer span.End()

	id := c.Param("id")
	span.SetAttributes(attribute.String("file_id", id))

	if err := fc.repo.DeleteFile(c, id); err != nil {
		utility.RecordError(ctxTrace, err)
		utility.LogError(c, fmt.Sprintf("Failed to delete file: %s", id), err)
		view.ErrorInternalServer(c, err)
		return
	}

	utility.LogInfo(c, fmt.Sprintf("Deleted file: %s", id))
	view.ViewDeleteFile(c)
}

func (fc *FileController) uploadFileToS3(c *gin.Context, file *multipart.FileHeader, filename string) (string, error) {
	ctxTrace, span := fileTracer.Start(c, "uploadFileToS3")
	defer span.End()

	f, err := file.Open()
	if err != nil {
		utility.RecordError(ctxTrace, err)
		return "", err
	}
	defer f.Close()

	return utility.UploadFileToS3(c, fc.s3Client, config.AWSBucketName(), filename, f)
}

func (fc *FileController) analyzeUploadWithBedrock(c *gin.Context, fileInfo *model.FileUpload) {
	ctxTrace, span := fileTracer.Start(c, "analyzeUploadWithBedrock")
	defer span.End()

	analysis, err := utility.AnalyzeWithBedrock(c, fmt.Sprintf("Analyze this file upload: Filename: %s, Size: %d bytes, Type: %s", fileInfo.FileName, fileInfo.FileSize, fileInfo.FileType))
	if err != nil {
		utility.RecordError(ctxTrace, err)
		utility.LogError(c, "Failed to analyze upload with Bedrock", err)
		return
	}

	span.SetAttributes(attribute.String("analysis_result", analysis))
	utility.LogInfo(c, fmt.Sprintf("Bedrock analysis completed for file: %s. Result: %s", fileInfo.FileName, analysis))
}

func generateUniqueFilename(originalFilename string) string {
	return fmt.Sprintf("%d_%s", time.Now().UnixNano(), originalFilename)
}
