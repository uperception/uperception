package storage

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	util "github.com/leometzger/mmonitoring/pkg/util"
)

type AwsStorage struct {
	client *s3.Client
	bucket string
}

func NewAwsStorage(client *s3.Client, bucket string) *AwsStorage {
	return &AwsStorage{
		client: client,
		bucket: bucket,
	}
}

func (s *AwsStorage) SaveLighthouseResult(url string, content io.Reader) error {
	key := util.GetPathFromUrl(url)

	_, err := s.client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: &s.bucket,
		Key:    &key,
		Body:   content,
	})

	if err != nil {
		return err
	}

	return err
}

func (s *AwsStorage) StoreMetadata() error {
	return nil
}
