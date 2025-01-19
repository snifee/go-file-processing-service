package handler

import (
	"go-file-upload-service/config"
	"go-file-upload-service/internal/model/dto"
	"go-file-upload-service/internal/repository"
	"go-file-upload-service/internal/service"

	"log"

	"github.com/gin-gonic/gin"
)

type FileUploadHandler struct {
	fileUploadService *service.FileUploadService
}

func NewFileUploadHandler(app *config.ApplicationBootstrap) *FileUploadHandler {

	repo := repository.NewFileUploadLogRepository(app.Database)
	service := service.NewFileUploadService(repo, app)

	handler := &FileUploadHandler{
		fileUploadService: service,
	}

	rg := app.Server.Engine.Group("file")
	rg.POST("/upload", handler.uploadFile)

	return handler
}

func (h *FileUploadHandler) uploadFile(c *gin.Context) {

	data, err := c.MultipartForm()

	request := dto.FileUpload{File: data.File["file"][0], Uploader: data.Value["uploader"][0]}

	if err != nil {
		log.Println(err.Error())

		c.JSON(500, dto.Response{
			Message:    "Internal server error",
			StatusCode: 500,
			Data:       nil,
		})
		return
	}

	if request.File.Header.Get("Content-Type") != "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet" {
		c.JSON(400, dto.Response{
			Message:    "File type must be pdf",
			StatusCode: 400,
			Data:       nil,
		})
		return
	}

	err = h.fileUploadService.UploadFile(request)

	if err != nil {
		log.Println(err.Error())
		c.JSON(500, dto.Response{
			Message:    "Error when uploading file to storage",
			StatusCode: 500,
			Data:       nil,
		})
		return
	}

	c.JSON(200, dto.Response{
		Message:    "Success",
		StatusCode: 200,
		Data:       nil,
	})
	return
}
