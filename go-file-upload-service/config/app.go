package config

import (
	"github.com/spf13/viper"
)

type ApplocationBootstrap struct {
	Database  *Database
	Minio     *MinioClient
	Publisher *Publisher
	Server    *Server
}

func NewAppication() *ApplocationBootstrap {

	dsn := viper.GetString("postgres.dsn")
	db := NewDatabase(dsn)

	endpoint := viper.GetString("minio.endpoint")
	accesskey := viper.GetString("minio.accessKey")
	secret := viper.GetString("minio.secretKey")
	minio := NewMinioClient(endpoint, accesskey, secret)

	url := viper.GetString("rabbitmq.url")
	publisher := NewPublisher(url)

	port := viper.GetString("server.port")
	server := NewServer(port)

	return &ApplocationBootstrap{
		Database:  db,
		Minio:     minio,
		Publisher: publisher,
		Server:    server,
	}
}
