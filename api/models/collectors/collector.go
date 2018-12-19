package models

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"reflect"

	"github.com/linkedin/goavro"

	"github.com/pkg/errors"
)

const (
	jsonExt         = ".json"
	jsonContentType = "application/json"
	// JSONType describes the enum for the json files
	JSONType       = "JSON"
	csvExt         = ".csv"
	csvContentType = "text/csv"
	// CSVType describes the enum for the csv files
	CSVType         = "CSV"
	avroExt         = ".avro"
	avroContentType = "binary/octet-stream"
	// AVROType describes the enum for the avro files
	AVROType       = "AVRO"
	txtExt         = ".txt"
	txtContentType = "text/plain"
	// TXTType describes the enum for the txt files
	TXTType = "TXT"
)

// Collectors is the object to encapsulate the various types of
// collectors that we offer
type Collectors struct {
	S3            *S3            `json:"s3"`
	GCloudStore   *GCloudStore   `json:"gcs"`
	AzureDataLake *AzureDataLake `json:"adl"`
	HTTP          *HTTP          `json:"http"`
	ApacheKafka   *ApacheKafka   `json:"kafka"`
}

// CollectorScheme contains the information about the collector files type
type CollectorScheme struct {
	FilePreview []FilePreview `json:"files_preview"`
	Collector   Collectors    `json:"collector"`
	FileType    string        `json:"file_type"`
}

// PreviewFile returns the list of files of a collector
func (c *Collectors) PreviewFile() (*CollectorScheme, error) {
	var filesPreview []FilePreview
	var fileType string
	if c.S3 != nil {
		files, err := c.S3.ListFiles()
		if err != nil {
			return nil, err
		}
		r, ct, err := c.S3.GetFirstValidObject(files)
		if err != nil {
			return nil, err
		}

		filesPreview, fileType, err = previewFile(r, ct)
		if err != nil {
			return nil, err
		}
	}

	result := &CollectorScheme{
		FilePreview: filesPreview,
		Collector:   *c,
		FileType:    fileType,
	}

	return result, nil
}

// IsRightExtension determines if the extension read in input is among the accepted ones
func IsRightExtension(ext string) bool {
	return ext == jsonExt || ext == csvExt || ext == avroExt || ext == txtExt
}

// DetermineFile will read a sequence of strings in input and will determine the
// type of file that the system will be then read
func previewFile(r io.Reader, ct string) ([]FilePreview, string, error) {
	switch ct {
	case jsonContentType:
		return handleJSONFile(r)
	case csvContentType:
		return handleCSVFile(r)
	case txtContentType:
		return handleTXTFile(r)
	case avroContentType:
		return handleAVROFile(r)
	default:
		return nil, "", errors.New("not available content type")
	}
}

// FilePreview is a struct that will be used to send back to the frontend
// the preview of the files' content of the collector
type FilePreview struct {
	Dimension string `json:"dimension"`
	Type      string `json:"type"`
	Example   string `json:"example"`
}

func handleJSONFile(r io.Reader) ([]FilePreview, string, error) {
	d, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, "", err
	}

	var result []FilePreview
	for _, line := range bytes.Split(d, []byte{'\n'}) {
		var v map[string]interface{}
		if err := json.Unmarshal(line, &v); err != nil {
			return nil, "", err
		}
		for k, v := range v {
			// TODO: test if string is datetime
			result = append(result, FilePreview{
				Dimension: fmt.Sprint(k),
				Example:   fmt.Sprint(v),
				Type:      fmt.Sprint(reflect.TypeOf(v)),
			})
		}
		break
	}

	return result, JSONType, nil
}

type field struct {
	Name string      `json:"name"`
	Type interface{} `json:"type"`
	Doc  string      `json:"doc"`
}

type avroSchema struct {
	Type   string  `json:"type"`
	Name   string  `json:"name"`
	Doc    string  `json:"doc"`
	Fields []field `json:"fields"`
}

func handleAVROFile(r io.Reader) ([]FilePreview, string, error) {
	ocfr, err := goavro.NewOCFReader(r)
	if err != nil {
		return nil, "", err
	}

	schema := ocfr.Codec().Schema()

	var s avroSchema
	if err := json.Unmarshal([]byte(schema), &s); err != nil {
		return nil, "", err
	}

	var result []FilePreview
	for _, field := range s.Fields {
		// TODO: test if string is datetime
		result = append(result, FilePreview{
			Dimension: field.Name,
			Example:   field.Doc,
			Type:      fmt.Sprint(field.Type),
		})
	}

	return result, AVROType, nil
}

func handleCSVFile(r io.Reader) ([]FilePreview, string, error) {
	reader := csv.NewReader(r)

	rows := []map[string]string{}
	var header []string

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, "", err
		}
		if header == nil {
			header = record
		} else {
			dict := map[string]string{}
			for i := range header {
				dict[header[i]] = record[i]
			}
			rows = append(rows, dict)
		}
	}

	// Empty file or just with header
	if len(rows) <= 0 {
		return nil, "", errors.New("empty csv file")
	}

	firstLine := rows[0]
	var result []FilePreview
	for key, value := range firstLine {
		// TODO: test if string is datetime
		result = append(result, FilePreview{
			Dimension: key,
			Example:   value,
			Type:      fmt.Sprint(reflect.TypeOf(value)),
		})
	}

	return result, CSVType, nil
}

func handleTXTFile(r io.Reader) ([]FilePreview, string, error) {
	return nil, "", nil
}
