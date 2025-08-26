package post

import (
	"time"

	"github.com/typetrait/lit/internal/domain/user"
)

type Status string

const (
	StatusDraft     Status = "draft"
	StatusPublished Status = "published"
)

type Post struct {
	ID        int64
	Title     string
	Slug      string
	Content   Content
	Status    Status
	Author    user.User
	CreatedAt time.Time
}
