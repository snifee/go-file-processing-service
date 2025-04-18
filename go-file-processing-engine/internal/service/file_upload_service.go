package service

import (
	"go-file-processing-engine/config"
	"go-file-processing-engine/internal/model/entity"
	"go-file-processing-engine/internal/repository"
	"go-file-processing-engine/utils"
	"log"
	"strconv"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type FileUploadService struct {
	minioClient       *config.MinioClient
	productRepository *repository.ProductRepository
	consumer          *config.Consumer
	configuration     *viper.Viper
	db                *gorm.DB
}

func NewFileUploadService(repository *repository.ProductRepository, app *config.ApplicationBootstrap) *FileUploadService {

	return &FileUploadService{
		minioClient: app.Minio,
		// fileRepository: fileRepository,
		productRepository: repository,
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

	rows, err := file.GetRows(file.SheetList[0])

	for index, row := range rows {
		if index == 0 {
			continue
		}
		product, err := s.toProduct(row)

		err = s.productRepository.Insert(product)
		if err != nil {
			log.Println("error when creating excel file")
			continue
		}

	}

	return nil
}

func (s *FileUploadService) toProduct(values []string) (entity.Product, error) {
	var result entity.Product
	var err error

	result.ProductID, err = strconv.Atoi(values[0])
	if err != nil {
		return result, err
	}
	result.ProductName = values[1]
	result.SupplierID, err = strconv.Atoi(values[2])
	if err != nil {
		return result, err
	}
	result.CategoryID, err = strconv.Atoi(values[3])
	if err != nil {
		return result, err
	}
	result.QuantityPerUnit = values[4]
	result.UnitPrice, err = strconv.ParseFloat(values[5], 64)
	if err != nil {
		return result, err
	}
	result.UnitsInStock, err = strconv.Atoi(values[6])
	if err != nil {
		return result, err
	}
	result.UnitsOnOrder, err = strconv.Atoi(values[7])
	if err != nil {
		return result, err
	}
	result.ReorderLevel, err = strconv.Atoi(values[8])
	if err != nil {
		return result, err
	}
	result.Discontinued, err = strconv.ParseBool(values[8])
	if err != nil {
		return result, err
	}

	return result, err
}
