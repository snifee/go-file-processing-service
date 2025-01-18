package entity

import (
	"time"

	"github.com/google/uuid"
)

type FileUploadLog struct {
	ID             uuid.UUID `gorm:"primary_key; unique;type:uuid; column:id;default:uuid_generate_v4()"`
	Filename       string    `gorm:"column:filename"`
	Extention      string    `gorm:"column:extention"`
	UploadDatetime time.Time `gorm:"column:upload_datetime"`
	UploadedBy     string    `gorm:"column:uploaded_by"`
	IsProcessed    bool      `gorm:"column:is_processed"`
}

func (f *FileUploadLog) TableName() string {
	return "file_upload_log"
}
