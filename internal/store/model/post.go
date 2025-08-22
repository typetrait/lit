package model

import (
	"time"

	"github.com/typetrait/lit/internal/domain/post"
)

type Post struct {
	ID            int64     `gorm:"primaryKey"`
	Title         string    `gorm:"unique;not null"`
	Slug          string    `gorm:"unique;not null"`
	ContentFormat uint8     `gorm:"not null"`
	ContentBody   string    `gorm:"type:text;not null"`
	AuthorID      int64     `gorm:"not null"`
	Author        User      `gorm:"foreignKey:AuthorID"`
	CreatedAt     time.Time `gorm:"not null"`
}

func (p *Post) ToDomainPost() post.Post {
	return post.Post{
		ID:    p.ID,
		Title: p.Title,
		Slug:  p.Slug,
		Content: post.Content{
			Format: post.ContentFormat(p.ContentFormat),
			Source: p.ContentBody,
		},
		Author:    p.Author.ToDomainUser(),
		CreatedAt: p.CreatedAt,
	}
}

func FromDomainPost(post post.Post) Post {
	return Post{
		ID:            post.ID,
		Title:         post.Title,
		Slug:          post.Slug,
		ContentFormat: uint8(post.Content.Format),
		ContentBody:   post.Content.Source,
		AuthorID:      post.Author.ID,
		Author:        FromDomainUser(post.Author),
		CreatedAt:     post.CreatedAt,
	}
}
