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
	csvExt          = ".csv"
	csvContentType  = "text/csv"
	avroExt         = ".avro"
	avroContentType = "binary/octet-stream"
	txtExt          = ".txt"
	txtContentType  = "text/plain"
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

// GetFile returns the list of files of a collector
func (c *Collectors) GetFile() (string, error) {
	// var files []string
	if c.S3 != nil {
		files, err := c.S3.ListFiles()
		if err != nil {
			return "", err
		}
		r, ct, err := c.S3.GetFirstValidObject(files)
		if err != nil {
			return "", err
		}
		return previewFile(r, ct)
	}

	if c.GCloudStore != nil {
		files, err := c.GCloudStore.ListFiles()
		if err != nil {
			return "", err
		}
		r, ct, err := c.GCloudStore.GetFirstValidObject(files)
		if err != nil {
			return "", err
		}
		return previewFile(r, ct)
	}

	// if c.AzureDataLake != nil {
	// 	files, err := c.AzureDataLake.ListFiles()
	// }

	// if c.HTTP != nil {
	// 	files, err := c.HTTP.ListFiles()
	// }

	// if c.ApacheKafka != nil {
	// 	files, err := c.ApacheKafka.ListFiles()
	// }

	return "", errors.New("Collector was not set it up correctly")
}

// IsRightExtension determines if the extension read in input is among the accepted ones
func IsRightExtension(ext string) bool {
	return ext == jsonExt || ext == csvExt || ext == avroExt || ext == txtExt
}

// DetermineFile will read a sequence of strings in input and will determine the
// type of file that the system will be then read
func previewFile(r io.Reader, ct string) (string, error) {
	switch ct {
	case jsonContentType:
		return handleJSONFile(r)
	case csvContentType:
		return handleCSVFile(r)
	case txtContentType:
		break
	case avroContentType:
		return handleAVROFile(r)
	default:
		break
	}
	return "", nil
}

// FilePreview is a struct that will be used to send back to the frontend
// the preview of the files' content of the collector
type FilePreview struct {
	Dimension string `json:"dimension"`
	Type      string `json:"type"`
	Example   string `json:"example"`
}

func handleJSONFile(r io.Reader) (string, error) {
	d, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}

	var result []FilePreview
	for _, line := range bytes.Split(d, []byte{'\n'}) {
		var v map[string]interface{}
		if err := json.Unmarshal(line, &v); err != nil {
			return "", err
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

	out, err := json.Marshal(result)
	if err != nil {
		return "", err
	}

	return string(out), nil
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

func handleAVROFile(r io.Reader) (string, error) {
	ocfr, err := goavro.NewOCFReader(r)
	if err != nil {
		return "", err
	}

	schema := ocfr.Codec().Schema()

	var s avroSchema
	if err := json.Unmarshal([]byte(schema), &s); err != nil {
		return "", err
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

	out, err := json.Marshal(result)
	if err != nil {
		return "", err
	}

	return string(out), nil
}

func handleCSVFile(r io.Reader) (string, error) {
	reader := csv.NewReader(r)

	rows := []map[string]string{}
	var header []string

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
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
		return "", errors.New("empty csv file")
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

	out, err := json.Marshal(result)
	if err != nil {
		return "", err
	}

	return string(out), nil
}
