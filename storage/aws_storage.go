package storage

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type AwsStorage struct {
	client *s3.Client
	bucket string
}

func NewAwsStorage() *AwsStorage {
	return &AwsStorage{
		bucket: "metzger.fot.br",
	}
}

// init s3 client
func (s *AwsStorage) initClient() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	s.client = s3.NewFromConfig(cfg)
}

// save the results into a bucket
func (s *AwsStorage) SaveLighthouseResult(domain string) error {
	if s.client == nil {
		s.initClient()
	}

	// s.client.PutObject(context.TODO(), &s3.PutObjectInput{
	// 	Bucket: &s.bucket,
	// 	Key:    &domain,
	// })

	return nil
}

// store metadata from the lighthouse result
func (s *AwsStorage) StoreMetadata() error {
	return nil
}
