package main

import (
	"go-file-processing-engine/config"
	"go-file-processing-engine/messaging"

	"log"
)

func main() {
	app := config.NewAppication()

	messaging := messaging.NewFileProcessingConsumer(app)
	messaging.StartProcessingFile()

	// app.Server.Engine.Run(":8080")
	log.Println("Application running")
}
