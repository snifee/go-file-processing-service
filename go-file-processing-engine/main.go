package main

import (
	"go-file-processing-engine/config"
	"go-file-processing-engine/messaging"

	"log"
)

func main() {
	log.Println("Application running")

	app := config.NewAppication()

	messaging := messaging.NewFileProcessingConsumer(app)
	// for {
	messaging.StartProcessingFile()
	// }

	// app.Server.Engine.Run(":8080")
}
