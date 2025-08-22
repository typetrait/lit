package post

import (
	"time"

	"github.com/typetrait/lit/internal/domain/user"
)

type Post struct {
	ID        int64
	Title     string
	Slug      string
	Content   Content
	Author    user.User
	CreatedAt time.Time
}
