package main

import (
	"go-file-processing-engine/config"
	"go-file-processing-engine/handler"
	"log"
)

func main() {
	app := config.NewAppication()

	handler.NewFileUploadHandler(app)

	app.Server.Engine.Run(":8080")
	log.Println("Application running")
}
