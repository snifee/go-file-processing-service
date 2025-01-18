package main

import (
	"context"
	"fmt"

	"github.com/minio/minio-go"
	"github.com/minio/minio-go/v7/pkg/notification"
)

func main() {
	queueArn := notification.NewArn("aws", "sqs", "us-east-1", "804605494417", "PhotoUpdate")

	queueConfig := notification.NewConfig(queueArn)
	queueConfig.AddEvents(minio.ObjectCreatedAll, minio.ObjectRemovedAll)
	queueConfig.AddFilterPrefix("photos/")
	queueConfig.AddFilterSuffix(".jpg")

	config := notification.Configuration{}
	config.AddQueue(queueConfig)

	err = minioClient.SetBucketNotification(context.Background(), "mybucket", config)
	if err != nil {
		fmt.Println("Unable to set the bucket notification: ", err)
		return
	}
}
