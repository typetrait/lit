package infrastructure

import (
	"context"
	"fmt"

	"github.com/typetrait/lit/internal/domain/post"
	"github.com/typetrait/lit/internal/infrastructure/model"
	"gorm.io/gorm"
)

type MediaRepository struct {
	db *gorm.DB
}

func NewMediaRepository(db *gorm.DB) *MediaRepository {
	return &MediaRepository{
		db: db,
	}
}

func (mr *MediaRepository) Create(ctx context.Context, mediaToCreate post.Media) (post.Media, error) {
	mediaModel := model.FromDomainMedia(mediaToCreate)
	if err := mr.db.WithContext(ctx).Create(&mediaModel).Error; err != nil {
		return post.Media{}, fmt.Errorf("creating media in repository: %w", err)
	}
	return mediaModel.ToDomainMedia(), nil
}
