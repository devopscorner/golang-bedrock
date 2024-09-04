// controller/file_controller.go
package controller

import (
	"context"
	"fmt"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/devopscorner/golang-bedrock/src/config"
	"github.com/devopscorner/golang-bedrock/src/model"
	"github.com/devopscorner/golang-bedrock/src/repository"
	"github.com/devopscorner/golang-bedrock/src/utility"
	"github.com/devopscorner/golang-bedrock/src/view"
	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

var (
	fileTracer = otel.Tracer("file-controller")
	log        = logrus.New()
)

type FileController struct {
	repo     repository.FileRepository
	s3Client *s3.Client
}

func NewFileController(repo repository.FileRepository, s3Client *s3.Client) *FileController {
	return &FileController{repo: repo, s3Client: s3Client}
}

func (fc *FileController) CreateFile(c *gin.Context) {
	ctx, span := fileTracer.Start(c.Request.Context(), "CreateFile")
	defer span.End()

	file, err := c.FormFile("file")
	if err != nil {
		fc.handleError(ctx, c, "Invalid file upload", err, http.StatusBadRequest)
		return
	}

	fileID := utility.GenerateID()
	filename := generateUniqueFilename(file.Filename, fileID)

	s3URL, err := fc.uploadFileToS3(ctx, file, filename)
	if err != nil {
		fc.handleError(ctx, c, "Failed to upload file to S3", err, http.StatusInternalServerError)
		return
	}

	fileUpload := model.FileUpload{
		ID:         fileID,
		FileName:   filename,
		FileSize:   file.Size,
		FileType:   file.Header.Get("Content-Type"),
		FileURL:    s3URL,
		UploadedBy: "user1@example.com",
		Analysis:   "",
	}

	validate := validator.New()
	if err := validate.Struct(fileUpload); err != nil {
		fc.handleError(ctx, c, "File validation failed", err, http.StatusBadRequest)
		return
	}

	err = fc.repo.CreateFile(ctx, &fileUpload)
	if err != nil {
		fc.handleError(ctx, c, "Failed to create file record", err, http.StatusInternalServerError)
		return
	}

	utility.RecordFileUpload(fileUpload.FileType, float64(fileUpload.FileSize))

	// Use a channel to handle the analysis result and metrics
	resultChan := make(chan struct {
		analysis string
		metrics  utility.Metrics
	})
	go fc.analyzeUploadWithBedrock(context.Background(), &fileUpload, resultChan)

	var analysisMetrics utility.Metrics
	// Wait for the analysis result or timeout
	select {
	case result := <-resultChan:
		fileUpload.Analysis = result.analysis
		analysisMetrics = result.metrics
		fc.logAnalysisMetrics(fileUpload.FileName, result.metrics)
	case <-time.After(30 * time.Second):
		fileUpload.Analysis = "Analysis timed out"
		log.WithField("filename", fileUpload.FileName).Warn("Bedrock analysis timed out")
	}

	// Update the file with the analysis result
	if err := fc.repo.UpdateFile(ctx, &fileUpload); err != nil {
		log.WithFields(logrus.Fields{
			"error":    err,
			"filename": fileUpload.FileName,
		}).Error("Failed to update file with analysis result")
	}

	span.SetAttributes(
		attribute.String("filename", fileUpload.FileName),
		attribute.Int64("file_size", fileUpload.FileSize),
		attribute.String("file_type", fileUpload.FileType),
	)

	log.WithFields(logrus.Fields{
		"filename": fileUpload.FileName,
		"size":     fileUpload.FileSize,
		"type":     fileUpload.FileType,
	}).Info("Created file")

	// Pass the metrics to the view function
	view.ViewCreateFile(c, fileUpload, analysisMetrics)
}

func (fc *FileController) FindAll(c *gin.Context) {
	ctx, span := fileTracer.Start(c.Request.Context(), "FindAll")
	defer span.End()

	files, err := fc.repo.FindAll(ctx)
	if err != nil {
		fc.handleError(ctx, c, "Failed to find files", err, http.StatusInternalServerError)
		return
	}

	span.SetAttributes(attribute.Int("file_count", len(files)))

	log.WithFields(logrus.Fields{
		"count": len(files),
	}).Info("Retrieved files")
	view.ViewFindAllFiles(c, files)
}

func (fc *FileController) FindByID(c *gin.Context) {
	ctx, span := fileTracer.Start(c.Request.Context(), "FindByID")
	defer span.End()

	id := c.Param("id")
	span.SetAttributes(attribute.String("file_id", id))

	file, err := fc.repo.FindByID(ctx, id)
	if err != nil {
		fc.handleError(ctx, c, "Failed to find file", err, http.StatusInternalServerError)
		return
	}
	if file == nil {
		fc.handleError(ctx, c, "File not found", nil, http.StatusNotFound)
		return
	}

	log.WithFields(logrus.Fields{
		"id": id,
	}).Info("Retrieved file")
	view.ViewFindFileByID(c, file)
}

func (fc *FileController) UpdateFile(c *gin.Context) {
	ctx, span := fileTracer.Start(c.Request.Context(), "UpdateFile")
	defer span.End()

	id := c.Param("id")
	span.SetAttributes(attribute.String("file_id", id))

	var input model.FileUpload
	if err := c.ShouldBindJSON(&input); err != nil {
		fc.handleError(ctx, c, "Invalid input for file update", err, http.StatusBadRequest)
		return
	}

	file, err := fc.repo.FindByID(ctx, id)
	if err != nil {
		fc.handleError(ctx, c, "Failed to find file for update", err, http.StatusNotFound)
		return
	}

	file.FileName = input.FileName
	file.FileSize = input.FileSize
	file.FileType = input.FileType
	file.FileURL = input.FileURL
	file.UploadedBy = input.UploadedBy
	file.Analysis = input.Analysis

	if err := fc.repo.UpdateFile(ctx, file); err != nil {
		fc.handleError(ctx, c, "Failed to update file", err, http.StatusInternalServerError)
		return
	}

	log.WithFields(logrus.Fields{
		"id": id,
	}).Info("Updated file")
	view.ViewUpdateFile(c, file)
}

func (fc *FileController) DeleteFile(c *gin.Context) {
	ctx, span := fileTracer.Start(c.Request.Context(), "DeleteFile")
	defer span.End()

	id := c.Param("id")
	span.SetAttributes(attribute.String("file_id", id))

	if err := fc.repo.DeleteFile(ctx, id); err != nil {
		fc.handleError(ctx, c, "Failed to delete file", err, http.StatusInternalServerError)
		return
	}

	log.WithFields(logrus.Fields{
		"id": id,
	}).Info("Deleted file")
	view.ViewDeleteFile(c)
}

func (fc *FileController) uploadFileToS3(ctx context.Context, file *multipart.FileHeader, filename string) (string, error) {
	ctx, span := fileTracer.Start(ctx, "uploadFileToS3")
	defer span.End()

	f, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()

	return utility.UploadFileToS3(ctx, fc.s3Client, config.AWSBucketName(), filename, f)
}

func (fc *FileController) analyzeUploadWithBedrock(ctx context.Context, fileInfo *model.FileUpload, resultChan chan<- struct {
	analysis string
	metrics  utility.Metrics
}) {
	ctx, cancel := context.WithTimeout(ctx, 25*time.Second)
	defer cancel()

	_, span := fileTracer.Start(ctx, "analyzeUploadWithBedrock")
	defer span.End()

	analysis, metrics, err := utility.AnalyzeWithBedrock(ctx, fmt.Sprintf("Analyze this file upload: Filename: %s, Size: %d bytes, Type: %s", fileInfo.FileName, fileInfo.FileSize, fileInfo.FileType))
	if err != nil {
		log.WithFields(logrus.Fields{
			"error":    err,
			"filename": fileInfo.FileName,
		}).Error("Failed to analyze upload with Bedrock")
		analysis = fc.getErrorAnalysis(err)
	}

	span.SetAttributes(attribute.String("analysis_result", analysis))
	log.WithFields(logrus.Fields{
		"filename": fileInfo.FileName,
		"analysis": analysis,
	}).Info("Bedrock analysis completed for file")

	resultChan <- struct {
		analysis string
		metrics  utility.Metrics
	}{analysis, metrics}
}

func (fc *FileController) getErrorAnalysis(err error) string {
	switch {
	case err.Error() == "Bedrock model not found. Please check your model ID and permissions":
		return "Analysis failed: Bedrock model not found. Please check configuration."
	case err.Error() == "Access denied to Bedrock model. Please check your IAM permissions":
		return "Analysis failed: Access denied to Bedrock model. Please check permissions."
	case err.Error() == "Invalid input for Bedrock model":
		return "Analysis failed: Invalid input for Bedrock model."
	case err.Error() == "Bedrock analysis timed out":
		return "Analysis failed: Bedrock request timed out."
	default:
		return "Analysis failed due to an error with Bedrock."
	}
}

func (fc *FileController) logAnalysisMetrics(filename string, metrics utility.Metrics) {
	log.WithFields(logrus.Fields{
		"filename":        filename,
		"totalLatency":    metrics.TotalLatency,
		"uploadLatency":   metrics.UploadLatency,
		"analysisLatency": metrics.AnalysisLatency,
		"inputTokens":     metrics.InputTokens,
		"outputTokens":    metrics.OutputTokens,
	}).Info("Bedrock analysis metrics")
}

func (fc *FileController) handleError(ctx context.Context, c *gin.Context, message string, err error, statusCode int) {
	utility.RecordError(ctx, err)
	log.WithFields(logrus.Fields{
		"error": err,
	}).Error(message)
	view.ErrorResponse(c, statusCode, message)
}

func generateUniqueFilename(originalFilename string, fileId string) string {
	return fmt.Sprintf("%s_%s", fileId, originalFilename)
}
