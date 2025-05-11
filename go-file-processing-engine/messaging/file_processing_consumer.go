package messaging

import (
	"encoding/json"
	"go-file-processing-engine/config"
	"go-file-processing-engine/internal/model/dto"
	"go-file-processing-engine/internal/repository"
	"go-file-processing-engine/internal/service"
	"log"
)

type FileProcessingConsumer struct {
	fileUploadService *service.FileUploadService
	consumer          *config.Consumer
}

func NewFileProcessingConsumer(app *config.ApplicationBootstrap) *FileProcessingConsumer {

	repo := repository.NewProductRepository(app.Database)
	service := service.NewFileUploadService(repo, app)

	consumer := &FileProcessingConsumer{
		fileUploadService: service,
		consumer:          app.Consumer,
	}

	return consumer
}

func (c *FileProcessingConsumer) StartProcessingFile() {
	msgs, err := c.consumer.StartReceiveMessage()

	if err != nil {
		log.Printf("error accour when try to start receive message : %v", err.Error())
	}

	var forever chan struct{}

	go func() {
		for m := range msgs {
			log.Printf("Received a message: %s", m.Body)

			// Create an empty struct
			var data dto.FileProcessingMessage
			err := json.Unmarshal(m.Body, &data)
			if err != nil {
				continue
			}
			// Read data from the buffer into the struct, handling endianness

			c.fileUploadService.ProcessFile(data.Filename)
			log.Printf("Done")
			m.Ack(false)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

}
