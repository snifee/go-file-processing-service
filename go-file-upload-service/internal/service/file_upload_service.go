package service

import (
	"fmt"

	"go-file-upload-service/config"
	"go-file-upload-service/internal/model/dto"
	"go-file-upload-service/internal/model/entity"
	"go-file-upload-service/internal/repository"
	"go-file-upload-service/utils"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/spf13/viper"
)

/*
FileUploadService is a service that will handle
logic of the endpoint
*/
type FileUploadService struct {
	minioClient     *config.MinioClient
	fileRepository  *repository.FileUploadLogRepository
	publisher       *config.Publisher
	rabbitMqChannel *config.PublisherChannel
	configuration   *viper.Viper
}

/*
NewFileUploadService is a constructor like func to
initialize a FileUploadService
*/
func NewFileUploadService(repository *repository.FileUploadLogRepository, app *config.ApplicationBootstrap) *FileUploadService {

	return &FileUploadService{
		minioClient:     app.Minio,
		fileRepository:  repository,
		publisher:       app.Publisher,
		rabbitMqChannel: app.Publisher.CreateChannel("file_process_1"),
		configuration:   app.Configuration,
	}
}

/*
UploadFile is a function that will become
func handler of REST API
*/
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

	fileSize := request.File.Size
	fileType := "xlsx"
	fileName := fmt.Sprintf("%s_%s.%s", fileID.String(), request.Uploader, fileType)

	info, err := s.minioClient.PutObject(file, fileName, fileSize, s.configuration.GetString("minio.dir.bucketName"))
	if err != nil {
		log.Printf("Error when upload file to minio: %s", err.Error())
		return err
	}

	logFile := entity.FileUploadLog{ID: fileID, Filename: fileName, Extention: fileType, UploadedBy: request.Uploader, UploadDatetime: time.Now(), IsProcessed: false}

	err = s.fileRepository.Insert(&logFile)

	if err != nil {
		log.Printf("Error when inserted log to database: %s", err.Error())
		return err
	}

	log.Printf("Successfully uploaded %s of size %d\n", fileName, info.Size)

	msg := dto.FileProcessingMessage{
		Uploader:        request.Uploader,
		Filename:        fileName,
		UploadTimestamp: time.Now(),
	}

	byteMessage, err := utils.JSONSerializer(msg)
	if err != nil {
		log.Printf("Error when serialize message: %s", err.Error())
		return err
	}

	s.rabbitMqChannel.SendMessage(byteMessage)
	return nil
}
