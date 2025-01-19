package dto

import (
	"mime/multipart"
)

type FileUpload struct {
	Uploader string
	File     *multipart.FileHeader
}
