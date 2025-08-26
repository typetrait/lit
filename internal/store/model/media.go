package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/typetrait/lit/internal/domain/media"
)

type Media struct {
	ID        int64     `gorm:"primary_key;auto_increment"`
	Key       uuid.UUID `gorm:"unique_index;not null"`
	Mime      string    `gorm:"not null"`
	Alt       string    `gorm:"not null"`
	Caption   string
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

func (m *Media) ToDomainMedia() media.Media {
	return media.Media{
		ID:        m.ID,
		Key:       m.Key,
		Mime:      m.Mime,
		Alt:       m.Alt,
		Caption:   m.Caption,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.CreatedAt,
	}
}

func FromDomainMedia(media media.Media) Media {
	return Media{
		ID:        media.ID,
		Key:       media.Key,
		Mime:      media.Mime,
		Alt:       media.Alt,
		Caption:   media.Caption,
		CreatedAt: media.CreatedAt,
		UpdatedAt: media.UpdatedAt,
	}
}
