// view/file_view.go
package view

import (
	"net/http"

	"github.com/devopscorner/golang-restfulapi-bedrock/src/model"
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
// Create new file
func ViewCreateFile(ctx *gin.Context, viewFile model.FileUpload) {
	ctx.JSON(http.StatusCreated, gin.H{"data": viewFile})
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
