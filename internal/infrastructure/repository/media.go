package repository

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

func (r *MediaRepository) Create(ctx context.Context, mediaToCreate post.Media) (post.Media, error) {
	mediaModel := model.FromDomainMedia(mediaToCreate)
	if err := r.db.WithContext(ctx).Create(&mediaModel).Error; err != nil {
		return post.Media{}, fmt.Errorf("creating media in repository: %w", err)
	}
	return mediaModel.ToDomainMedia(), nil
}

func (r *MediaRepository) FindByID(ctx context.Context, id int64) (post.Media, error) {
	var mediaModel model.Media
	if err := r.db.WithContext(ctx).First(&mediaModel, "id = ?", id).Error; err != nil {
		return post.Media{}, fmt.Errorf("finding media by ID: %w", err)
	}
	return mediaModel.ToDomainMedia(), nil
}
