package storage

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"regexp"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
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
	key := GetS3KeyFromUrl(url)
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

// Return the path pattern a s3 key path to store the lighthouse result
func GetS3KeyFromUrl(url string) string {
	r := regexp.MustCompile(`https?:\/\/(?P<Domain>[a-zA-Z0-9.]+)(?P<Path>\/[a-zA-Z0-9\/]+)?`)
	result := r.FindStringSubmatch(url)

	hasher := sha1.New()
	hasher.Write([]byte(url))
	id := hex.EncodeToString(hasher.Sum(nil))

	if len(result) == 3 {
		domain := result[1]
		pathing := []string{
			"reports",
			domain,
			nowAsPath(),
			id + ".json",
		}

		return strings.Join(pathing, "/")
	}

	return ""
}

func nowAsPath() string {
	datePath := time.Now().Format("2006/02/01")
	timePath := strings.ReplaceAll(time.Now().Format("15:04"), ":", "/")
	return datePath + timePath
}
