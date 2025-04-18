package config

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type ApplicationBootstrap struct {
	Database      *gorm.DB
	Minio         *MinioClient
	Consumer      *Consumer
	Server        *Server
	Configuration *viper.Viper
}

func NewAppication() *ApplicationBootstrap {

	configuration := NewViperConfig()
	db := NewDatabase(configuration.GetString("postgres.dsn"))
	minio := NewMinioClient(
		configuration.GetString("minio.endpoint"),
		configuration.GetString("minio.accessKey"),
		configuration.GetString("minio.secretKey"),
	)
	consumer := NewConsumer(configuration.GetString("rabbitmq.url"), "file_process_1")
	server := NewServer(configuration.GetString("server.port"))

	return &ApplicationBootstrap{
		Database:      db,
		Minio:         minio,
		Consumer:      consumer,
		Server:        server,
		Configuration: configuration,
	}
}
