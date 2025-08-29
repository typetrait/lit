package media

import (
	"context"
	"fmt"
)

type GetMedia struct {
	storage         Storage
	mediaRepository Repository
}

func NewGetMedia(storage Storage, mediaRepository Repository) *GetMedia {
	return &GetMedia{
		storage:         storage,
		mediaRepository: mediaRepository,
	}
}

func (g *GetMedia) Get(ctx context.Context, query GetMediaQuery) (GetMediaQueryResult, error) {
	m, err := g.mediaRepository.FindByID(ctx, query.MediaID)
	if err != nil {
		return GetMediaQueryResult{}, fmt.Errorf("finding media: %w", err)
	}

	content, err := g.storage.Get(ctx, m.Key.String())
	if err != nil {
		return GetMediaQueryResult{}, fmt.Errorf("fetching media from storage: %w", err)
	}

	return NewGetMediaQueryResult(m, content), nil
}
