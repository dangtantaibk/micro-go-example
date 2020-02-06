package fileutil

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"time"

	"github.com/unidoc/unioffice/spreadsheet"

	_fileUpload "tng/common/adapters/fileupload"
	"tng/common/logger"
	"tng/common/utils/cfgutil"
)

// Exporter define export interface
type Exporter interface {
	Export(exportType ExportType) (string, error)
	ExportByFileName(exportType ExportType, fileName string) (string, error)
}

type exporter struct {
	input    interface{}
	filePath string
}

// ExporterFactory create exporter factory
func ExporterFactory(input interface{}, filePath string) Exporter {
	return &exporter{
		input:    input,
		filePath: filePath,
	}
}

// Export execute export. Default export xlsx file
func (e *exporter) Export(exportType ExportType) (string, error) {
	switch exportType {
	case ExportCSVFile:
		err := e.ExportCSV()
		if err != nil {
			return "", err
		}
	default:
		exportType = ExportXLSFile
		err := e.ExportExcel()
		if err != nil {
			return "", err
		}
	}
	ggStorageURL := cfgutil.Load("GCP_SERVICE_BUCKET_URL_PREFIX")
	ggBucketName := cfgutil.Load("GCP_SERVICE_PUBLIC_BUCKET_NAME")
	ggFolderName := cfgutil.Load("GCP_SERVICE_BUCKET_FILE_EXPORT_PREFIX")
	ggKeyString := cfgutil.Load("GCP_SERVICE_PUBLIC_BUCKET_ACCOUNT_KEY")
	ggKeyFile := cfgutil.Load("GCP_SERVICE_PUBLIC_BUCKET_ACCOUNT_KEY_FILE")
	var fileUploadAdapter _fileUpload.Adapter
	var ggCredential []byte
	var err error
	if ggKeyFile != "" {
		ggCredential, err = ioutil.ReadFile(ggKeyFile)
		if err != nil {
			logger.Log("Failed to read google storage key file: ", err)
			os.Exit(1)
		}
	} else if ggKeyString != "" {
		ggCredential = []byte(ggKeyString)
	}
	if ggCredential != nil {
		fileUploadAdapter, err = _fileUpload.NewGoogleStorageAdapter(ggStorageURL, ggFolderName, ggCredential, ggBucketName)
	}
	if err != nil {
		logger.Log("Failed to create instance of google storage adapter: ", err)
		os.Exit(1)
	}
	f, err := os.Open(e.filePath)
	if err != nil {
		return "", err
	}
	return fileUploadAdapter.Upload(f, fmt.Sprintf("%v", exportType))
}

func (e *exporter) ExportByFileName(exportType ExportType, fileName string) (string, error) {
	switch exportType {
	case ExportCSVFile:
		err := e.ExportCSV()
		if err != nil {
			return "", err
		}
	default:
		exportType = ExportXLSFile
		err := e.ExportExcel()
		if err != nil {
			return "", err
		}
	}
	ggStorageURL := cfgutil.Load("GCP_SERVICE_BUCKET_URL_PREFIX")
	ggBucketName := cfgutil.Load("GCP_SERVICE_PUBLIC_BUCKET_NAME")
	ggFolderName := cfgutil.Load("GCP_SERVICE_BUCKET_FILE_EXPORT_PREFIX")
	ggKeyString := cfgutil.Load("GCP_SERVICE_PUBLIC_BUCKET_ACCOUNT_KEY")
	ggKeyFile := cfgutil.Load("GCP_SERVICE_PUBLIC_BUCKET_ACCOUNT_KEY_FILE")
	var fileUploadAdapter _fileUpload.Adapter
	var ggCredential []byte
	var err error
	if ggKeyFile != "" {
		ggCredential, err = ioutil.ReadFile(ggKeyFile)
		if err != nil {
			logger.Log("Failed to read google storage key file: ", err)
			os.Exit(1)
		}
	} else if ggKeyString != "" {
		ggCredential = []byte(ggKeyString)
	}
	if ggCredential != nil {
		fileUploadAdapter, err = _fileUpload.NewGoogleStorageAdapter(ggStorageURL, ggFolderName, ggCredential, ggBucketName)
	}
	if err != nil {
		logger.Log("Failed to create instance of google storage adapter: ", err)
		os.Exit(1)
	}
	f, err := os.Open(e.filePath)
	if err != nil {
		return "", err
	}
	fullFileName := fileName + "." + string(exportType)
	return fileUploadAdapter.UploadWithByFileName(f, fullFileName)
}

// ExportCSV Export file csv
func (e *exporter) ExportCSV() error {
	vietnamZone, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	f, err := os.Create(e.filePath)
	if err != nil {
		return err
	}
	tp := reflect.TypeOf(e.input).Elem()
	items := reflect.ValueOf(e.input)
	var data [][]string
	var header []string
	for i := 0; i < tp.NumField(); i++ {
		header = append(header, tp.Field(i).Name)
	}
	data = append(data, header)
	for i := 0; i < items.Len(); i++ {
		item := reflect.Indirect(items.Index(i))
		var row []string
		for j := 0; j < len(data[0]); j++ {
			field := item.FieldByName(data[0][j])
			if !field.IsValid() {
				continue
			}
			switch field.Kind() {
			case reflect.String:
				row = append(row, field.String())
			case reflect.Int64, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32:
				row = append(row, fmt.Sprintf("%v", field.Int()))
			case reflect.Float64:
				row = append(row, fmt.Sprintf("%.f", field.Float()))
			case reflect.Bool:
				row = append(row, fmt.Sprintf("%v", field.Bool()))
			case reflect.Ptr:
				value := field.Elem()
				if value.IsValid() {
					row = append(row, fmt.Sprintf("%v", value.Int()))
				} else {
					row = append(row, "")
				}
			case reflect.Struct: // time.Time field
				switch t := field.Interface().(type) {
				case time.Time:
					if t.IsZero() {
						row = append(row, "")
					} else {
						row = append(row, t.In(vietnamZone).Format(TimeFormatInFile))
					}
				default:
					row = append(row, "")
				}
			}

		}
		data = append(data, row)
	}
	err = csv.NewWriter(f).WriteAll(data)
	defer f.Close()
	if err != nil {
		return err
	}
	return nil
}

// ExportExcel Export data to excel file
func (e *exporter) ExportExcel() error {
	ss := spreadsheet.New()
	sheet := ss.AddSheet()
	sheet.Name()

	colMap := make(map[int]string)

	numCol := 0
	headerRow := sheet.AddRow()
	tp := reflect.TypeOf(e.input).Elem()
	for i := 0; i < tp.NumField(); i++ {
		cell := headerRow.AddCell()
		cell.SetString(tp.Field(i).Name)
		colMap[i] = tp.Field(i).Name
		numCol++
	}
	vietnamZone, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	items := reflect.ValueOf(e.input)
	for i := 0; i < items.Len(); i++ {
		row := sheet.AddRow()
		item := reflect.Indirect(items.Index(i))
		for j := 0; j < numCol; j++ {
			cell := row.AddCell()
			field := item.FieldByName(colMap[j])
			if !field.IsValid() {
				continue
			}
			switch field.Kind() {
			case reflect.String:
				cell.SetString(field.String())
			case reflect.Int64, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32:
				cell.SetString(fmt.Sprintf("%v", field.Int()))
			case reflect.Float64:
				cell.SetString(fmt.Sprintf("%.f", field.Float()))
			case reflect.Bool:
				cell.SetString(fmt.Sprintf("%v", field.Bool()))
			case reflect.Struct: // time.Time field
				switch t := field.Interface().(type) {
				case time.Time:
					if t.IsZero() {
						cell.SetString("")
					} else {

						cell.SetString(t.In(vietnamZone).Format(TimeFormatInFile))
					}
				default:
					cell.SetString("")
				}

			default:
				cell.SetString("")
			}
		}
	}
	// Save to disk
	if err := ss.Validate(); err != nil {
		return err
	}

	return ss.SaveToFile(e.filePath)
}
