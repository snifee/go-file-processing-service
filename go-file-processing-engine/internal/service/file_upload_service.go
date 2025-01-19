package service

import (
	"go-file-processing-engine/config"
	"go-file-processing-engine/internal/model/entity"
	"go-file-processing-engine/internal/repository"
	"go-file-processing-engine/utils"
	"log"

	"github.com/spf13/viper"
)

type FileUploadService struct {
	minioClient     *config.MinioClient
	orderRepository *repository.OrderDetailRepository
	consumer        *config.Consumer
	rabbitMqChannel *config.ConsumerChannel
	configuration   *viper.Viper
}

func NewFileUploadService(repository *repository.OrderDetailRepository, app *config.ApplicationBootstrap) *FileUploadService {

	return &FileUploadService{
		minioClient: app.Minio,
		// fileRepository: fileRepository,
		orderRepository: repository,
		// rabbitMqChannel: app.Consumer.CreateChannel(),
		configuration: app.Configuration,
	}
}

func (s *FileUploadService) ProcessFile(fileName string) error {

	reader, err := s.minioClient.GetObjectFileReader(fileName, s.configuration.GetString("minio.dir.bucketName"))
	if err != nil {
		log.Println("error when get object")
		return err
	}

	file, err := utils.NewExcelFile(reader, "sheet1")
	if err != nil {
		log.Println("error when creating excel file")
		return err
	}

	rows, err := file.GetRows()

	for _, row := range rows {
		order := entity.Order{}

		err := s.orderRepository.Create(&order)
		if err != nil {
			log.Println("error when creating excel file")
			return err
		}

	}

	return nil
}
