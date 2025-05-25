package main

import (
	"go-file-upload-service/config"
	"go-file-upload-service/handler"
	"log"
)

func main() {
	app := config.NewAppication()

	handler.NewFileUploadHandler(app)

	app.Server.Engine.Run(":8080")
	log.Println("Application running")
}
