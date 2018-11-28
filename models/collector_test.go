package models

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/aws/aws-sdk-go/aws/credentials"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func uploadObject(c *s3.S3, contentType, val string) error {
	// upload object
	p := s3.PutObjectInput{
		Bucket:      aws.String("bucket"),
		Key:         aws.String("hello.json"),
		ACL:         aws.String("public-read"),
		Body:        strings.NewReader(val),
		ContentType: aws.String(contentType),
	}

	_, err := c.PutObject(&p)
	return err
}

func mockS3Client(region, bucketName string) (*s3.S3, error) {
	sess := session.Must(session.NewSession(&aws.Config{
		Endpoint:         aws.String("http://192.168.99.100:4572"),
		Region:           aws.String(region),
		S3ForcePathStyle: aws.Bool(true),
		DisableSSL:       aws.Bool(true),
		Credentials:      credentials.NewStaticCredentials("foo", "var", ""),
	}))

	// new session
	c := s3.New(sess, &aws.Config{})

	// create bucket
	c.CreateBucket(&s3.CreateBucketInput{Bucket: aws.String(bucketName)})

	return c, nil
}

func setCollectorS3(region, bucketName string) (*Collectors, error) {
	c := &Collectors{
		S3: &S3{
			AccessKeyID:     "test",
			AccessKeySecret: "test",
			Region:          region,
			Bucket:          bucketName,
			Prefix:          "",
		},
	}
	s3Client, err := mockS3Client(region, bucketName)
	if err != nil {
		return nil, err
	}

	// set up mock client
	c.S3.Client = s3Client
	return c, nil
}

func TestS3GetFile(t *testing.T) {
	bucketName := "bucket"
	region := "eu-west-1"

	c, err := setCollectorS3(region, bucketName)
	assert.Nil(t, err)

	err = uploadObject(c.S3.Client, jsonContentType, `{"name":"hello"}`)
	assert.Nil(t, err)

	f, err := c.GetFile()
	assert.Nil(t, err)

	assert.Equal(t, `[{"dimension":"name","type":"string","example":"hello"}]`, f)
}
