package media

import (
	"context"

	"github.com/typetrait/lit/internal/domain/media"
)

type UploadMedia struct {
}

func NewUploadMedia() *UploadMedia {
	return &UploadMedia{}
}

func (upload *UploadMedia) Upload(ctx context.Context, cmd UploadMediaCommand) (media.Media, error) {
	panic("implement me")
}
