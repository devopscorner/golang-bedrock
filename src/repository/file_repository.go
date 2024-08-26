// repository/file_repository.go
package repository

import (
	"context"

	"github.com/devopscorner/golang-restfulapi-bedrock/src/model"
	"github.com/devopscorner/golang-restfulapi-bedrock/src/utility"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"gorm.io/gorm"
)

var tracer = otel.Tracer("file-repository")

type FileRepository interface {
	FindAll(ctx context.Context) ([]model.FileUpload, error)
	FindByID(ctx context.Context, id string) (*model.FileUpload, error)
	CreateFile(ctx context.Context, file *model.FileUpload) error
	UpdateFile(ctx context.Context, file *model.FileUpload) error
	DeleteFile(ctx context.Context, id string) error
}

type fileRepository struct {
	db *gorm.DB
}

func NewFileRepository(db *gorm.DB) FileRepository {
	return &fileRepository{db: db}
}

func (r *fileRepository) FindAll(ctx context.Context) ([]model.FileUpload, error) {
	ctx, span := tracer.Start(ctx, "FindAll")
	defer span.End()

	var files []model.FileUpload
	result := r.db.WithContext(ctx).Find(&files)
	if result.Error != nil {
		utility.RecordError(ctx, result.Error)
		return nil, result.Error
	}

	span.SetAttributes(attribute.Int("file_count", len(files)))
	return files, nil
}

func (r *fileRepository) FindByID(ctx context.Context, id string) (*model.FileUpload, error) {
	ctx, span := tracer.Start(ctx, "FindByID")
	defer span.End()

	span.SetAttributes(attribute.String("file_id", id))

	var file model.FileUpload
	result := r.db.WithContext(ctx).First(&file, "id = ?", id)
	if result.Error != nil {
		utility.RecordError(ctx, result.Error)
		return nil, result.Error
	}
	return &file, nil
}

func (r *fileRepository) CreateFile(ctx context.Context, file *model.FileUpload) error {
	ctx, span := tracer.Start(ctx, "CreateFile")
	defer span.End()

	span.SetAttributes(
		attribute.String("filename", file.FileName),
		attribute.Int64("file_size", file.FileSize),
		attribute.String("file_type", file.FileType),
	)

	result := r.db.WithContext(ctx).Create(file)
	if result.Error != nil {
		utility.RecordError(ctx, result.Error)
	}
	return result.Error
}

func (r *fileRepository) UpdateFile(ctx context.Context, file *model.FileUpload) error {
	ctx, span := tracer.Start(ctx, "UpdateFile")
	defer span.End()

	span.SetAttributes(
		attribute.String("file_id", file.ID),
		attribute.String("filename", file.FileName),
		attribute.Int64("file_size", file.FileSize),
		attribute.String("file_type", file.FileType),
	)

	result := r.db.WithContext(ctx).Save(file)
	if result.Error != nil {
		utility.RecordError(ctx, result.Error)
	}
	return result.Error
}

func (r *fileRepository) DeleteFile(ctx context.Context, id string) error {
	ctx, span := tracer.Start(ctx, "DeleteFile")
	defer span.End()

	span.SetAttributes(attribute.String("file_id", id))

	result := r.db.WithContext(ctx).Delete(&model.FileUpload{}, "id = ?", id)
	if result.Error != nil {
		utility.RecordError(ctx, result.Error)
	}
	return result.Error
}
