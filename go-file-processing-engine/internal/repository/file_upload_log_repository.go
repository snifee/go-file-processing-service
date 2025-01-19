package repository

import (
	"go-file-processing-engine/internal/model/entity"

	"gorm.io/gorm"
)

type FileUploadLogRepository struct {
	Repository[entity.FileUploadLog]
}

func NewFileUploadLogRepository(db *gorm.DB) *FileUploadLogRepository {
	return &FileUploadLogRepository{
		Repository: Repository[entity.FileUploadLog]{
			DB: db,
		},
	}
}
