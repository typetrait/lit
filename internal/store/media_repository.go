package store

import (
	"context"
	"fmt"

	"github.com/typetrait/lit/internal/domain/media"
	"github.com/typetrait/lit/internal/store/model"
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

func (mr *MediaRepository) Create(ctx context.Context, mediaToCreate media.Media) (media.Media, error) {
	mediaModel := model.FromDomainMedia(mediaToCreate)
	if err := mr.db.WithContext(ctx).Create(&mediaModel).Error; err != nil {
		return media.Media{}, fmt.Errorf("creating media in repository: %w", err)
	}
	return mediaModel.ToDomainMedia(), nil
}
