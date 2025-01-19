package service

import (
	"fmt"

	"go-file-upload-service/config"
	"go-file-upload-service/internal/model/dto"
	"go-file-upload-service/internal/model/entity"
	"go-file-upload-service/internal/repository"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/spf13/viper"
)

type FileUploadService struct {
	minioClient     *config.MinioClient
	fileRepository  *repository.FileUploadLogRepository
	publisher       *config.Publisher
	rabbitMqChannel *config.PublisherChannel
	configuration   *viper.Viper
}

func NewFileUploadService(repository *repository.FileUploadLogRepository, app *config.ApplicationBootstrap) *FileUploadService {

	return &FileUploadService{
		minioClient:     app.Minio,
		fileRepository:  repository,
		publisher:       app.Publisher,
		rabbitMqChannel: app.Publisher.CreateChannel(),
		configuration:   app.Configuration,
	}
}

func (s *FileUploadService) UploadFile(request dto.FileUpload) error {

	file, err := request.File.Open()

	if err != nil {
		return err
	}

	fileID, err := uuid.NewV7()

	if err != nil {
		log.Printf("Error when creating id when insert to file_uplod_log: %s", err.Error())
		return err
	}

	fileName := fmt.Sprintf("%s_%s", fileID.String(), request.Uploader)
	fileSize := request.File.Size
	fileType := "xlsx"

	info, err := s.minioClient.PutObject(file, fileName, fileSize, s.configuration.GetString("minio.dir.bucketName"))
	if err != nil {
		log.Printf("Error when upload file to minio: %s", err.Error())
		return err
	}

	logFile := entity.FileUploadLog{ID: fileID, Filename: fileName, Extention: fileType, UploadedBy: request.Uploader, UploadDatetime: time.Now(), IsProcessed: false}

	err = s.fileRepository.Create(&logFile)

	if err != nil {
		log.Printf("Error when inserted log to database: %s", err.Error())
		return err
	}

	log.Printf("Successfully uploaded %s of size %d\n", fileName, info.Size)

	s.rabbitMqChannel.SendMessage([]byte(fileName))
	return nil
}
