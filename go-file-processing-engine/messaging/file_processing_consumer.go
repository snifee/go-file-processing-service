package handler

import (
	"go-file-processing-engine/config"
	"go-file-processing-engine/internal/repository"
	"go-file-processing-engine/internal/service"
)

type FileProcessingConsumer struct {
	fileUploadService *service.FileUploadService
	consumer          *config.Consumer
}

func NewFileProcessingConsumer(app *config.ApplicationBootstrap) *FileProcessingConsumer {

	repo := repository.NewFileUploadLogRepository(app.Database)
	service := service.NewFileUploadService(repo, app)

	handler := &FileProcessingConsumer{
		fileUploadService: service,
		consumer:          app.Consumer,
	}

	return handler
}

func (h *FileProcessingConsumer) processFile() {

}
