package utils

import (
	"io"

	"github.com/xuri/excelize/v2"
)

type ExcelFile struct {
	file      *excelize.File
	header    *excelize.HeaderFooterOptions
	SheetList []string
}

func NewExcelFile(reader io.Reader, sheetName string) (*ExcelFile, error) {
	file, err := excelize.OpenReader(reader, excelize.Options{})
	if err != nil {
		return nil, err
	}

	return &ExcelFile{
		file:      file,
		SheetList: file.GetSheetList(),
	}, nil
}
func (f *ExcelFile) GetRows(sheetList string) ([][]string, error) {

	rows, err := f.file.GetRows("Sheet")
	if err != nil {
		return nil, err
	}

	return rows, nil
}
