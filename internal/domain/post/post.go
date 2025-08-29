package post

import (
	"errors"
	"time"

	"github.com/typetrait/lit/internal/domain/user"
)

type Status string

const (
	StatusDraft     Status = "draft"
	StatusPublished Status = "published"
)

var (
	ErrInvalidContentFormat = errors.New("invalid content format")
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

func (p Post) IsDraft() bool {
	return p.Status == StatusDraft
}

func (p Post) IsPublished() bool {
	return p.Status == StatusPublished
}
