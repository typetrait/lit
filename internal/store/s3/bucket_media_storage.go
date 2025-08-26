package s3

import (
	"context"

	"github.com/typetrait/lit/internal/app/media"
)

type BucketMediaStorage struct {
}

func NewMediaStorage() *BucketMediaStorage {
	return &BucketMediaStorage{}
}

func (s *BucketMediaStorage) Upload(ctx context.Context, media media.MediaUpload) error {
	panic("implement me")
}
