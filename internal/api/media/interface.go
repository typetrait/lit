package media

import (
	"context"

	"github.com/typetrait/lit/internal/app/media"
	domain "github.com/typetrait/lit/internal/domain/post"
)

type uploadMedia interface {
	Upload(ctx context.Context, command media.UploadMediaCommand) (domain.Media, error)
}

type getMedia interface {
	Get(ctx context.Context, query media.GetMediaQuery) (media.GetMediaQueryResult, error)
}
