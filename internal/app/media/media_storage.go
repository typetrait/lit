package media

import (
	"context"
	"io"
)

type Storage interface {
	Get(ctx context.Context, key string) (io.ReadCloser, error)
	Put(ctx context.Context, key string, readCloser io.ReadCloser) error
	Delete(ctx context.Context, key string) error
}
