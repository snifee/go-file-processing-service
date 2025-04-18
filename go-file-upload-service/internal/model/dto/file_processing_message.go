package dto

import "time"

type FileProcessingMessage struct {
	Uploader        string    `json:"uploader"`
	Filename        string    `json:"filename"`
	UploadTimestamp time.Time `json:"upload_timestamp"`
}
