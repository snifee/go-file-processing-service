package config

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type ApplicationBootstrap struct {
	Database      *gorm.DB
	Minio         *MinioClient
	Publisher     *Publisher
	Server        *Server
	Configuration *viper.Viper
}

func NewAppication() *ApplicationBootstrap {

	configuration := NewViperConfig()

	dsn := configuration.GetString("postgres.dsn")

	db := NewDatabase(dsn)

	endpoint := configuration.GetString("minio.endpoint")
	accesskey := configuration.GetString("minio.accessKey")
	secret := configuration.GetString("minio.secretKey")
	minio := NewMinioClient(endpoint, accesskey, secret)

	url := configuration.GetString("rabbitmq.url")
	publisher := NewPublisher(url)

	port := configuration.GetString("server.port")
	server := NewServer(port)

	return &ApplicationBootstrap{
		Database:      db,
		Minio:         minio,
		Publisher:     publisher,
		Server:        server,
		Configuration: configuration,
	}
}
