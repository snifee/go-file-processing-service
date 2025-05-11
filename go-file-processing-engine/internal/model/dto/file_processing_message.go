package dto

type FileProcessingMessage struct {
	Uploader        string `json:"uploader"`
	Filename        string `json:"filename"`
	UploadTimestamp string `json:"upload_timestamp"`
}
