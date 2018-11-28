package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"reflect"

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

	// if c.GCloudStore != nil {
	// 	files, err := c.GCloudStore.ListFiles()
	// }

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
		return handleJSONfile(r)
	case csvContentType:
		break
	case txtContentType:
		break
	case avroContentType:
		break
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

func handleJSONfile(r io.Reader) (string, error) {
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
