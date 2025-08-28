package media

import (
	"context"

	"github.com/typetrait/lit/internal/app/media"
	domain "github.com/typetrait/lit/internal/domain/post"
)

type uploadMedia interface {
	Upload(ctx context.Context, command media.UploadMediaCommand) (domain.Media, error)
}
