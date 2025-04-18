package repository

import (
	"go-file-processing-engine/internal/model/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FileUploadLogRepository struct {
	Repository[entity.FileUploadLog, uuid.UUID]
}

func NewFileUploadLogRepository(db *gorm.DB) *FileUploadLogRepository {
	return &FileUploadLogRepository{
		Repository: Repository[entity.FileUploadLog, uuid.UUID]{
			db: db,
		},
	}
}
