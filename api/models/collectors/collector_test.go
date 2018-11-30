package collectors

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/aws/aws-sdk-go/aws/credentials"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func uploadObject(c *s3.S3, bucket, key, contentType string, b io.ReadSeeker) error {
	// upload object
	p := s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(key),
		ACL:         aws.String("public-read"),
		Body:        b,
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

func TestS3GetFileJSON(t *testing.T) {
	bucketName := "bucket"
	region := "eu-west-1"

	c, err := setCollectorS3(region, bucketName)
	assert.Nil(t, err)

	err = uploadObject(c.S3.Client, bucketName, "hello.json", jsonContentType, strings.NewReader(`{"name":"hello"}`))
	assert.Nil(t, err)

	f, err := c.GetFile()
	assert.Nil(t, err)

	assert.Equal(t, `[{"dimension":"name","type":"string","example":"hello"}]`, f)
}

func TestS3GetFileAVRO(t *testing.T) {
	bucketName := "bucket"
	region := "eu-west-1"

	c, err := setCollectorS3(region, bucketName)
	assert.Nil(t, err)

	d, err := os.Open("../fixtures/test.avro")
	if err != nil {
		t.Fatal(err)
	}

	err = uploadObject(c.S3.Client, bucketName, "hello.avro", avroContentType, d)
	assert.Nil(t, err)

	f, err := c.GetFile()
	assert.Nil(t, err)

	assert.Equal(t, `[{"dimension":"name","type":"string","example":"hello"}]`, f)
}

func TestS3GetFileCSV(t *testing.T) {
	bucketName := "bucket"
	region := "eu-west-1"

	c, err := setCollectorS3(region, bucketName)
	assert.Nil(t, err)

	d, err := os.Open("../fixtures/test.csv")
	if err != nil {
		t.Fatal(err)
	}

	err = uploadObject(c.S3.Client, bucketName, "test.csv", csvContentType, d)
	assert.Nil(t, err)

	f, err := c.GetFile()
	assert.Nil(t, err)

	assert.Equal(t, `[{"dimension":"id","type":"string","example":"1"},{"dimension":"name","type":"string","example":"hello"},{"dimension":"created_at","type":"string","example":"2018-11-01T10:00:00Z"}]`, f)
}
