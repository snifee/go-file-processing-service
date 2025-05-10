package config

import (
	"bytes"
	"context"
	"io"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioClient struct {
	client *minio.Client
}

func NewMinioClient(endpoint, accesskey, secret string) *MinioClient {
	// endpoint := viper.GetString("minio.endpoint")
	// accessKeyID := viper.GetString("minio.accessKey")
	// secretAccessKey := viper.GetString("minio.secretKey")
	// useSSL := viper.GetBool("minio.useSSL")

	// Initialize minio client object.
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accesskey, secret, ""),
		Secure: true,
	})
	if err != nil {
		log.Println(endpoint)
		log.Fatalln("Failed when make connection to Minio Server \n \t" + err.Error())

	}

	log.Printf("%#v\n", client) // minioClient is now set up

	return &MinioClient{
		client: client,
	}
}

func (m *MinioClient) GetObjectFileReader(fileName, bucketName string) (io.Reader, error) {
	var result io.Reader
	ctx := context.Background()

	object, err := m.client.GetObject(ctx, bucketName, fileName, minio.GetObjectOptions{})
	if err != nil {
		log.Printf("Fail to get file [%s] from Minio Server | %s\n", fileName, err.Error())
		return nil, err
	}

	var objectBytes []byte
	_, err = object.Read(objectBytes)
	if err != nil {
		log.Printf("Fail to read file [%s] from Minio Server | %s\n", fileName, err.Error())
		return nil, err
	}

	result = bytes.NewReader(objectBytes)

	return result, nil
}
