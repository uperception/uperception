package storage

import (
	"context"
	"fmt"
	"io"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type AwsStorage struct {
	client *s3.Client
	bucket string
}

func NewAwsStorage() *AwsStorage {
	return &AwsStorage{
		bucket: "mmonitoring",
	}
}

// init s3 client
func (s *AwsStorage) initClient() {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatal(err)
	}
	s.client = s3.NewFromConfig(cfg)
}

// save the results into a bucket
func (s *AwsStorage) SaveLighthouseResult(url string, content io.Reader) error {
	if s.client == nil {
		s.initClient()
	}

	key := GetS3KeyFromUrl(url)
	fmt.Println("inserting lighthouse", key)

	_, err := s.client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: &s.bucket,
		Key:    &key,
		Body:   content,
	})

	fmt.Println(err)

	return err
}

// store metadata from the lighthouse result
func (s *AwsStorage) StoreMetadata() error {
	return nil
}

// @TODO: Think a better way to build this path
func GetS3KeyFromUrl(url string) string {
	r := regexp.MustCompile(`https?:\/\/(?P<Domain>[a-zA-Z0-9.]+)(?P<Path>\/[a-zA-Z0-9\/]+)?`)
	result := r.FindStringSubmatch(url)

	if len(result) == 3 {
		domain := result[1]
		path := strings.ReplaceAll(result[2], "/", "-")
		datePath := time.Now().Format("2006/02/01/23")
		filename := strings.ReplaceAll(time.Now().Format("15:04"), ":", "-") + ".json"

		if path == "" {
			path = "root"
		}

		return "reports/" + domain + "/" + path + "/" + datePath + "/" + filename
	}

	return ""
}
