package fileutil

import (
	"bytes"
	"context"
	"errors"
	"io"
	"tng/common/logger"
	"mime/multipart"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// ValidateFileSize validate size file upload
func ValidateFileSize(w http.ResponseWriter, r *http.Request) error {
	if r.ContentLength > MaxFileSize {
		return errors.New("request too large")
	}
	r.Body = http.MaxBytesReader(w, r.Body, MaxFileSize)
	err := r.ParseMultipartForm(1024)
	if err != nil {
		return err
	}
	return nil
}

// Pointer Type Define
const (
	Int32Ptr  = "*int32"
	Int64Ptr  = "*int64"
	IntPtr    = "*int"
	StringPtr = "*string"
	BoolPtr   = "*bool"
)

// SaveFile store file user upload
// f, fh, err := this.GetFile("file") get multipart
func SaveFile(f multipart.File, localFilePath string) error {
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, f); err != nil {
		return err
	}
	tmpFile, err := os.Create(localFilePath)
	defer func() {
		if err := tmpFile.Close(); err != nil {
			logger.Errorf(context.Background(), "close file error: %v", err)
		}
	}()

	if err != nil {
		return err
	}
	_, err = buf.WriteTo(tmpFile)
	if err != nil {
		return err
	}
	return nil
}

// DeleteFile Delete file after import success
func DeleteFile(localFilePath string) error {
	return os.Remove(localFilePath)
}

// ParseValueString Assign value to struct
func ParseValueString(field reflect.Value, inputS string) error {
	if inputS == "" {
		return nil
	}
	switch field.Kind() {
	case reflect.String:
		field.SetString(inputS)
		return nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		inputS = strings.TrimSpace(inputS)
		int64Value, err := strconv.ParseInt(inputS, 10, 64)
		if err != nil {
			return err
		}
		field.SetInt(int64Value)

		return nil
	case reflect.Float64:
		inputS = strings.TrimSpace(inputS)
		float64Value, err := strconv.ParseFloat(inputS, 64)
		if err != nil {
			return err
		}
		field.SetFloat(float64Value)
		return nil
	case reflect.Bool:
		boolValue, err := strconv.ParseBool(inputS)
		if err != nil {
			return err
		}
		field.SetBool(boolValue)
		return nil
	case reflect.Ptr:
		switch reflect.TypeOf(field.Interface()).String() {
		case Int64Ptr, Int32Ptr, IntPtr:
			number, err := strconv.ParseInt(inputS, 10, 64)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(&number))
		case BoolPtr:
			boolValue, err := strconv.ParseBool(inputS)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(&boolValue))
		case StringPtr:
			field.Set(reflect.ValueOf(&inputS))
		}
		return nil
	case reflect.Struct: // time.Time field
		switch field.Interface().(type) {
		case time.Time:
			timeValue, err := time.Parse(TimeFormatInFile, inputS)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(timeValue))
			return nil
		default:
			return errors.New("Type " + field.Interface().(string) + " not implemented yet")
		}
	default:
		return errors.New("Type " + field.Kind().String() + " not implemented yet")
	}
}

// GetFieldName return array name of struct
func GetFieldName(obj interface{}) []string {
	listFields := make([]string, 0)
	tp := reflect.TypeOf(obj)
	for i := 0; i < tp.NumField(); i++ {
		listFields = append(listFields, tp.Field(i).Name)
	}
	return listFields
}

// GetHeaderTypeSupport validate file upload support
func GetHeaderTypeSupport(t string) int {
	switch t {
	case XlxsTypeHeader, XlsTypeHeader:
		return 1
	case CsvTypeHeader:
		return 2
	default:
		return 0
	}
}
