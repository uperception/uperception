package storage

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	util "github.com/leometzger/mmonitoring/pkg/util"
)

type AwsStorage struct {
	client        *s3.Client
	presignClient *s3.PresignClient
	bucket        string
}

func NewAwsStorage(client *s3.Client, bucket string, presignClient *s3.PresignClient) *AwsStorage {
	return &AwsStorage{
		client:        client,
		presignClient: presignClient,
		bucket:        bucket,
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

func (s *AwsStorage) AddAvatar(key string, avatar io.Reader, ext string) (string, error) {
	avatarKey := getAvatarKey(key) + ext
	_, err := s.client.PutObject(context.Background(), &s3.PutObjectInput{
		Bucket: &s.bucket,
		Key:    &avatarKey,
		Body:   avatar,
	})

	return avatarKey, err
}

func (s *AwsStorage) GetAvatarUrl(key string) (*SignedUrl, error) {
	avatarKey := getAvatarKey(key)
	result, err := s.presignClient.PresignGetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: &s.bucket,
		Key:    &avatarKey,
	})
	if err != nil {
		return nil, err
	}

	signedUrl := SignedUrl{
		Url:    result.URL,
		Header: result.SignedHeader,
	}

	return &signedUrl, nil
}

func (s *AwsStorage) RemoveAvatar(key string) error {
	avatarKey := getAvatarKey(key)
	_, err := s.client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: &s.bucket,
		Key:    &avatarKey,
	})

	return err
}

func getAvatarKey(key string) string {
	return "avatars/" + key
}
