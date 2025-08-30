package media

import (
	"context"
	"io"
)

type Storage interface {
	Get(ctx context.Context, key string) (io.Reader, error)
	Put(ctx context.Context, key string, readCloser io.Reader) error
	Delete(ctx context.Context, key string) error
}
