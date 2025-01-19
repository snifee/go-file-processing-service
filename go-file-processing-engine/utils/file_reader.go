package utils

import (
	"io"

	"github.com/xuri/excelize/v2"
)

type ExcelFile struct {
	file *excelize.File
}

func NewExcelFile(reader io.Reader, sheetName string) (*ExcelFile, error) {
	file, err := excelize.OpenReader(reader)
	if err != nil {
		return nil, err
	}

	return &ExcelFile{
		file: file,
	}, nil
}

func (f *ExcelFile) GetRows() ([][]string, error) {
	rows, err := f.file.GetRows("Sheet")
	if err != nil {
		return nil, err
	}

	return rows, nil
}
