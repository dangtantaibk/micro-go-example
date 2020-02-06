package fileutil

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"reflect"

	"github.com/unidoc/unioffice/spreadsheet"
)

// Importer define import interface
type Importer interface {
	Import(fileType int) error
}

type importer struct {
	file       string
	dest       interface{}
	listFields []string
}

// ImporterFactory create import factory
func ImporterFactory(file string, dest interface{}, listFields []string) Importer {
	return &importer{
		file:       file,
		dest:       dest,
		listFields: listFields,
	}
}

func (i *importer) Import(fileType int) error {
	switch fileType {
	case 1: // Import excel .xls, .xlsx
		return i.ImportExcel()
	case 2: // Import csv .csv
		return i.ImportCSV()
	default:
		return fmt.Errorf("import type %v not yet support", fileType)
	}
}

// ImportCSV Import file CSV
func (i *importer) ImportCSV() error {
	f, err := os.Open(i.file)
	if err != nil {
		return err
	}
	data, err := csv.NewReader(f).ReadAll()
	defer f.Close()
	if err != nil {
		return err
	}
	const ValidField bool = true
	listFieldsMap := make(map[string]bool)
	for _, fieldName := range i.listFields {
		listFieldsMap[fieldName] = ValidField
	}
	value := reflect.ValueOf(i.dest)
	if value.Kind() != reflect.Ptr {
		return errors.New("can not parse struct data")
	}
	if value.IsNil() {
		return errors.New("can not parse struct data")
	}
	direct := value.Elem()
	slice := direct.Type()
	if slice.Kind() != reflect.Slice {
		return errors.New("can not parse struct data")
	}
	basePtr := slice.Elem()
	if basePtr.Kind() != reflect.Ptr {
		return errors.New("can not parse struct data")
	}
	base := basePtr.Elem()
	if base.Kind() != reflect.Struct {
		return errors.New("can not parse struct data")
	}

	colMap := make(map[int]string)
	for rowIndex, row := range data {
		if rowIndex == 0 {
			for cellIndex, cell := range row {
				colMap[cellIndex] = cell
			}
		} else {
			isEmptyRow := true
			stPtr := reflect.New(base)
			for cellIndex, cell := range row {

				isEmptyRow = false

				if !listFieldsMap[colMap[cellIndex]] {
					continue
				}
				field := stPtr.Elem().FieldByName(colMap[cellIndex])
				if !field.IsValid() {
					continue
				}
				err = ParseValueString(field, cell)
				if err != nil {
					break
				}
			}
			if !isEmptyRow {
				direct.Set(reflect.Append(direct, stPtr))
			}
		}
	}
	return nil
}

// ImportExcel import file excel
func (i *importer) ImportExcel() error {
	// create field map
	const ValidField int = 1
	listFieldsMap := make(map[string]int)
	for _, fieldName := range i.listFields {
		listFieldsMap[fieldName] = ValidField
	}

	// Read file to memory
	w, err := spreadsheet.Open(i.file)
	if err != nil {
		return err
	}
	defer w.Close()

	// value is ref.value *[]*Struct
	value := reflect.ValueOf(i.dest)
	if value.Kind() != reflect.Ptr {
		return errors.New("can not parse struct data 1")
	}
	if value.IsNil() {
		return errors.New("can not parse struct data 2")
	}
	// direct is ref.value []*Struct
	direct := value.Elem()
	// slice is ref.type []*Struct
	slice := direct.Type()
	if slice.Kind() != reflect.Slice {
		return errors.New("can not parse struct data 3" + slice.Kind().String())
	}
	basePtr := slice.Elem()
	if basePtr.Kind() != reflect.Ptr {
		return errors.New("can not parse struct data 4")
	}
	// base is ref.type struct
	base := basePtr.Elem()
	if base.Kind() != reflect.Struct {
		return errors.New("can not parse struct data 5")
	}
	// Map column name to object field name: A => Id, B => Name
	colMap := make(map[string]string)
	for _, sheet := range w.Sheets() {
		for rowIndex, row := range sheet.Rows() {
			if rowIndex == 0 {
				for _, cell := range row.Cells() {
					colName, err := cell.Column()
					if err != nil {
						return err
					}
					colMap[colName] = cell.GetString()
				}
			} else {
				isEmptyRow := true
				stPtr := reflect.New(base)
				for _, cell := range row.Cells() {
					if cell.GetString() == "" {
						continue
					}
					isEmptyRow = false
					colName, err := cell.Column()
					if err != nil {
						return err
					}
					if listFieldsMap[colMap[colName]] != ValidField {
						continue
					}
					field := stPtr.Elem().FieldByName(colMap[colName])
					if !field.IsValid() {
						continue
					}
					err = ParseValueString(field, cell.GetString())
					if err != nil {
						break
					}
				}
				if !isEmptyRow {
					direct.Set(reflect.Append(direct, stPtr))
				}
			}
		}
	}
	return nil
}
