package messaging

import (
	"go-file-processing-engine/config"
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
	msgs, err := c.consumer.ReceiveMessage()

	if err != nil {
		log.Printf("error accour when try to start receive message : %v", err.Error())
	}

	go func() {
		for m := range msgs {
			log.Printf("Received a message: %s", m.Body)
			// dotCount := bytes.Count(m.Body, []byte("."))
			c.fileUploadService.ProcessFile(string(m.Body))
			log.Printf("Done")
			m.Ack(false)
		}
	}()

}
