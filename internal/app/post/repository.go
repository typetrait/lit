package post

import (
	"context"
	"errors"

	"github.com/typetrait/lit/internal/domain/post"
)

var (
	ErrPostCreationFailed = errors.New("post creation failed")
	ErrPostUpdateFailed   = errors.New("post update failed")
	ErrPostNotFound       = errors.New("post not found")
)

type Repository interface {
	Create(ctx context.Context, post post.Post) (post.Post, error)
	Update(ctx context.Context, post post.Post) (post.Post, error)
	FindAll(ctx context.Context) ([]post.Post, error)
	FindAllPublished(ctx context.Context) ([]post.Post, error)
	FindByID(ctx context.Context, id int64) (post.Post, error)
	FindBySlug(ctx context.Context, slug string) (post.Post, error)
}
