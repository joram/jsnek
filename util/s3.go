package util

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func WriteToS3(bucket, key, content string) {
	s, err := session.NewSession(&aws.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Get file size and read the file content into a buffer
	var size = int64(len(content))
	buffer := make([]byte, size)
	reader := strings.NewReader(content)

	// Config settings: this is where you choose the bucket, filename, content-type etc.
	// of the file you're uploading.
	_, err = s3.New(s).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(bucket),
		Key:                  aws.String(key),
		ACL:                  aws.String("private"),
		Body:                 reader,
		ContentLength:        aws.Int64(size),
		ContentType:          aws.String(http.DetectContentType(buffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
	})
	if err != nil {
		log.Fatal(err)
	}
}

func GetFromS3(bucket, key string) []byte {
	s, err := session.NewSession(&aws.Config{})
	if err != nil {
		log.Fatal(err)
	}

	resp, err := s3.New(s).GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		log.Fatal(err)
	}

	s3objectBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return s3objectBytes
}
