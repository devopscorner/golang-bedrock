// view/file_view.go
package view

import (
	"net/http"

	"github.com/devopscorner/golang-bedrock/src/model"
	"github.com/devopscorner/golang-bedrock/src/utility"
	"github.com/gin-gonic/gin"
)

// ----- View Response -----

// GET /files
// Find all files
func ViewFindAllFiles(ctx *gin.Context, viewFiles []model.FileUpload) {
	ctx.JSON(http.StatusOK, gin.H{"data": viewFiles})
}

// GET /files/:id
// Find a file
func ViewFindFileByID(ctx *gin.Context, viewFile *model.FileUpload) {
	ctx.JSON(http.StatusOK, gin.H{"data": viewFile})
}

// POST /files
// ViewCreateFile includes metrics in the response
func ViewCreateFile(c *gin.Context, file model.FileUpload, metrics utility.Metrics) {
	c.JSON(http.StatusCreated, gin.H{
		"data": gin.H{
			"id":         file.ID,
			"fileName":   file.FileName,
			"fileSize":   file.FileSize,
			"fileType":   file.FileType,
			"fileURL":    file.FileURL,
			"uploadedBy": file.UploadedBy,
			"analysis":   file.Analysis,
			"createdAt":  file.CreatedAt,
			"updatedAt":  file.UpdatedAt,
			"metrics": gin.H{
				"totalLatency":    metrics.TotalLatency.String(),
				"uploadLatency":   metrics.UploadLatency.String(),
				"analysisLatency": metrics.AnalysisLatency.String(),
				"inputTokens":     metrics.InputTokens,
				"outputTokens":    metrics.OutputTokens,
			},
		},
	})
}

// PUT /files/:id
// Update a file
func ViewUpdateFile(ctx *gin.Context, viewFile *model.FileUpload) {
	ctx.JSON(http.StatusOK, gin.H{"data": viewFile})
}

// DELETE /files/:id
// Delete a file
func ViewDeleteFile(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"data": true})
}
