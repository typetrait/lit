package media

import (
	"context"

	"github.com/typetrait/lit/internal/domain/post"
)

type Repository interface {
	Create(ctx context.Context, media post.Media) (post.Media, error)
	FindByID(ctx context.Context, id int64) (post.Media, error)
}
