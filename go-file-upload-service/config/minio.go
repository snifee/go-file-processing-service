package config

import (
	"context"
	"log"
	"mime/multipart"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/viper"
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

func (m *MinioClient) PutObject(file multipart.File, fileName string, fileSize int64) (minio.UploadInfo, error) {

	ctx := context.Background()

	contentType := "application/octet-stream"
	bucketName := viper.GetString("minio.dir.bucketName")

	// Upload the test file with FPutObject
	info, err := m.client.PutObject(ctx, bucketName, fileName, file, fileSize, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Println("Fail upload file to Minio Server\n" + err.Error())
		return info, err
	}

	return info, nil
}
