package s3

import (
	"context"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type BucketMediaStorage struct {
	s3         *s3.Client
	bucketName string
}

func NewMediaStorage(s3 *s3.Client, bucketName string) *BucketMediaStorage {
	return &BucketMediaStorage{
		s3:         s3,
		bucketName: bucketName,
	}
}

func (s *BucketMediaStorage) Get(ctx context.Context, key string) (io.ReadCloser, error) {
	out, err := s.s3.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(s.bucketName),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, fmt.Errorf("getting object from s3 bucket: %w", err)
	}
	return out.Body, nil
}

func (s *BucketMediaStorage) Put(ctx context.Context, key string, readCloser io.ReadCloser) error {
	_, err := s.s3.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(s.bucketName),
		Key:    aws.String(key),
		Body:   readCloser,
	})
	if err != nil {
		return fmt.Errorf("putting object into s3 bucket: %w", err)
	}
	return nil
}

func (s *BucketMediaStorage) Delete(ctx context.Context, key string) error {
	_, err := s.s3.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(s.bucketName),
		Key:    aws.String(key),
	})
	if err != nil {
		return fmt.Errorf("deleting object from s3 bucket: %w", err)
	}
	return nil
}
