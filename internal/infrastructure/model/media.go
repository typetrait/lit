package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/typetrait/lit/internal/domain/post"
)

type Media struct {
	ID        int64     `gorm:"primary_key;auto_increment"`
	PostID    int64     `gorm:"not null"`
	Post      Post      `gorm:"foreignKey:PostID"`
	Key       uuid.UUID `gorm:"unique_index;not null"`
	Mime      string    `gorm:"not null"`
	Alt       string    `gorm:"not null"`
	Caption   *string
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

func (m *Media) ToDomainMedia() post.Media {
	return post.Media{
		ID:        m.ID,
		Post:      m.Post.ToDomainPost(),
		Key:       m.Key,
		Mime:      m.Mime,
		Alt:       m.Alt,
		Caption:   m.Caption,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.CreatedAt,
	}
}

func FromDomainMedia(media post.Media) Media {
	return Media{
		ID:        media.ID,
		PostID:    media.Post.ID,
		Post:      FromDomainPost(media.Post),
		Key:       media.Key,
		Mime:      media.Mime,
		Alt:       media.Alt,
		Caption:   media.Caption,
		CreatedAt: media.CreatedAt,
		UpdatedAt: media.UpdatedAt,
	}
}
