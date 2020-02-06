package fileutil

import (
	"errors"
	"fmt"
	"tng/common/utils/cfgutil"
	"mime/multipart"
	"time"
)

// ExportType define export type
type ExportType string

// define const
const (
	XlxsTypeHeader              = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
	XlsTypeHeader               = "application/vnd.ms-excel"
	CsvTypeHeader               = "text/csv"
	MaxFileSize                 = 8 * 1024 * 1024 // 8MB
	ExportCSVFile    ExportType = "csv"
	ExportXLSFile    ExportType = "xlsx"
	TimeFormatInFile            = "02/01/2006 15:04:05"
	DateFormatInFile            = "20060201"
)

// FileBuilder define interface FileBuilder
type FileBuilder interface {
	SetData(file multipart.File, fh *multipart.FileHeader) FileBuilder
	Export() FileBuilder
	BuildImport() FileImport
	BuildExport(exportType ExportType) FileExport
}

type fileBuilder struct {
	file multipart.File
	fh   *multipart.FileHeader
}

// FileImport define interface file import
type FileImport interface {
	SetMaxSize(maxSize int64) FileImport
	Execute(input interface{}, dataType interface{}) error
}

// FileExport define interface export builder
type FileExport interface {
	Execute(input interface{}) (file string, err error)
	ExecuteByFileName(input interface{}, fileName string) (file string, err error)
}

type fileExport struct {
	fileType   ExportType
	exportPath string
}

type fileImport struct {
	file     multipart.File
	fh       *multipart.FileHeader
	maxSize  int64
	filePath string
}

// Execute run builder import file
func (f *fileImport) Execute(input, dataType interface{}) (err error) {
	if f.file == nil || f.fh == nil {
		return errors.New("file upload is invalid")
	}
	fileType := f.fh.Header.Get("Content-Type")
	fileTypeUpload := GetHeaderTypeSupport(fileType)
	if fileTypeUpload == 0 {
		return fmt.Errorf("invalid file %v format", fileType)
	}
	if f.fh.Size > f.maxSize {
		return fmt.Errorf("file upload limited is %v", f.maxSize)
	}
	cols := GetFieldName(dataType)
	localFileName := fmt.Sprintf("%v_%v", time.Now().Unix(), f.fh.Filename)
	filePath := fmt.Sprintf("%v/%v", f.filePath, localFileName)
	err = SaveFile(f.file, filePath)
	if err != nil {
		return fmt.Errorf("can not save file %v to disk", filePath)
	}
	err = ImporterFactory(filePath, input, cols).Import(fileTypeUpload)
	if err != nil {
		return err
	}
	defer func() {
		_ = DeleteFile(filePath) // delete file after import
	}()
	return nil
}

// Execute Run export file
func (f *fileExport) Execute(input interface{}) (outfile string, err error) {
	localFileName := fmt.Sprintf("%v_%v", time.Now().Unix(), "export_tmp")
	filePath := fmt.Sprintf("%v/%v", f.exportPath, localFileName)
	fileCloud, err := ExporterFactory(input, filePath).Export(f.fileType)
	if err != nil {
		return "", nil
	}
	return fileCloud, nil
}

// ExecuteByFileName for execute
func (f *fileExport) ExecuteByFileName(input interface{}, fileName string) (file string, err error) {
	localFileName := fmt.Sprintf("%v_%v", time.Now().Unix(), "export_tmp")
	filePath := fmt.Sprintf("%v/%v", f.exportPath, localFileName)
	fileCloud, err := ExporterFactory(input, filePath).ExportByFileName(f.fileType, fileName)
	if err != nil {
		return "", nil
	}
	return fileCloud, nil
}

// Build define function Build for import builder
func (f *fileImport) Build() FileImport {
	return f
}

// Build define function Build for export builder
func (f *fileExport) Build() FileExport {
	return f
}

// SetMaxSize set maximum file upload, default 8Mb
func (f *fileImport) SetMaxSize(maxSize int64) FileImport {
	f.maxSize = maxSize
	return f
}

// SetFilePath define folder template upload file, default get config IMPORT_FILE_DIR
func (f *fileImport) SetFilePath(filePath string) FileImport {
	f.filePath = filePath
	return f
}

// SetFilePath define folder template export file, default get config EXPORT_FILE_DIR
func (f *fileExport) SetFilePath(filePath string) FileExport {
	f.exportPath = filePath
	return f
}

// Import return
func (f *fileBuilder) SetData(file multipart.File, fh *multipart.FileHeader) FileBuilder {
	f.file = file
	f.fh = fh
	return f
}

func (f *fileBuilder) Export() FileBuilder {
	return f
}

// NewFileBuilder create instance file builder
func NewFileBuilder() FileBuilder {
	return &fileBuilder{}
}

func (f *fileBuilder) BuildImport() FileImport {
	return &fileImport{file: f.file,
		fh:       f.fh,
		maxSize:  MaxFileSize,
		filePath: cfgutil.Load("IMPORT_FILE_DIR"), // set default folder upload
	}
}

func (f *fileBuilder) BuildExport(exportType ExportType) FileExport {
	return &fileExport{
		fileType:   exportType,
		exportPath: cfgutil.Load("EXPORT_FILE_DIR"),
	}
}
