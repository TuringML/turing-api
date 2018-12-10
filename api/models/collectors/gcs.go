package collectors

import (
	"context"
	"io"
	"log"

	"cloud.google.com/go/storage"
	"github.com/pkg/errors"
	"google.golang.org/api/option"
)

// GCloudStore is the object that can be used to get/store files
type GCloudStore struct {
	Client            string
	ProjectID         string `json:"projectId"`
	ServiceAccountKey string `json:"serviceAccountKey"`
	Bucket            string `json:"bucket"`
	Prefix            string `json:"prefix"`
}

// NewClient creates a new instance of a Google Cloud Storage client
func (g *GCloudStore) NewClient() {
	ctx := context.Background()

	// Creates a client.
	client, err := storage.NewClient(ctx, option.WithCredentials(nil))
	if err != nil {
		log.Fatal(err)
	}
	client.Bucket("").Object("sdsds").NewWriter(ctx)
}

// ListFiles returns the list of files in GCloudStorage bucket
func (g *GCloudStore) ListFiles() ([]string, error) {
	return nil, nil
}

// GetFirstValidObject returns the first valid object from the list of files
// passed as input, the content type and eventually an error
func (g *GCloudStore) GetFirstValidObject(files []string) (io.Reader, string, error) {
	for _, file := range files {
		return g.GetObject(file)
	}
	return nil, "", errors.New("no valid files to read")
}

// GetObject will return the body of the selected object
func (g *GCloudStore) GetObject(f string) (io.Reader, string, error) {
	// o, err :=

	// if err != nil {
	// 	return nil, "", err
	// }

	// return o.Body, *o.ContentType, nil
	return nil, "", nil
}
