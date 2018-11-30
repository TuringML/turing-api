package collectors

import (
	"io"
	"path/filepath"

	"github.com/pkg/errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// S3 is the AWS S3 object that can be used to get/store files
type S3 struct {
	Client          *s3.S3
	Bucket          string `json:"bucket"`
	Prefix          string `json:"prefix"`
	AccessKeyID     string `json:"accessKeyId"`
	AccessKeySecret string `json:"accessKeySecret"`
	Region          string `json:"region"`
}

// NewClient creates a new S3 client with the given credentials
func (s *S3) NewClient() error {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(s.Region),
		Credentials: credentials.NewStaticCredentials(s.AccessKeyID, s.AccessKeySecret, ""),
	})
	if err != nil {
		return err
	}

	s.Client = s3.New(sess)
	return nil
}

// ListFiles returns the list of files in S3 bucket
func (s *S3) ListFiles() ([]string, error) {
	var files []string
	err := s.Client.ListObjectsPages(&s3.ListObjectsInput{
		Bucket: aws.String(s.Bucket),
		Prefix: aws.String(s.Prefix),
	}, func(p *s3.ListObjectsOutput, last bool) (shouldContinue bool) {
		for _, obj := range p.Contents {
			extension := filepath.Ext(*obj.Key) // a file
			if extension != "" && IsRightExtension(extension) {
				files = append(files, *obj.Key)
			}
		}
		return !last
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}

// GetFirstValidObject returns the first valid object from the list of files
// passed as input, the content type and eventually an error
func (s *S3) GetFirstValidObject(files []string) (io.Reader, string, error) {
	for _, file := range files {
		return s.GetObject(file)
	}
	return nil, "", errors.New("no valid files to read")
}

// GetObject will return the body of the selected object
func (s *S3) GetObject(f string) (io.Reader, string, error) {
	o, err := s.Client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(f),
	})

	if err != nil {
		return nil, "", err
	}

	return o.Body, *o.ContentType, nil
}
