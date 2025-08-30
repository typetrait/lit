package infrastructure

import (
	"context"
	"fmt"

	"github.com/typetrait/lit/internal/domain/settings"
	"github.com/typetrait/lit/internal/infrastructure/model"
	"gorm.io/gorm"
)

type SettingsRepository struct {
	db *gorm.DB
}

func NewSettingsRepository(db *gorm.DB) *SettingsRepository {
	return &SettingsRepository{
		db: db,
	}
}

func (r *SettingsRepository) FindAll(ctx context.Context) (settings.Settings, error) {
	var settingsModel []model.Settings
	if err := r.db.WithContext(ctx).Find(&settingsModel).Error; err != nil {
		return settings.Settings{}, fmt.Errorf("finding settings: %w", err)
	}

	settingsMap := map[string]string{}
	additionalSettings := map[string]any{}
	for _, s := range settingsModel {
		settingsMap[s.Name] = s.Value
		if s.Name != settings.KeyBlogName && s.Name != settings.KeyBlogSubtitle && s.Name != settings.KeyBlogAbout {
			additionalSettings[s.Name] = s.Value
		}
	}
	return settings.Settings{
		BlogName:           settingsMap[settings.KeyBlogName],
		BlogSubtitle:       settingsMap[settings.KeyBlogSubtitle],
		BlogAbout:          settingsMap[settings.KeyBlogAbout],
		AdditionalSettings: additionalSettings,
	}, nil
}
