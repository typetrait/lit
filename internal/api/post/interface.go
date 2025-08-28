package post

import (
	"context"

	"github.com/typetrait/lit/internal/app/post"
	domain "github.com/typetrait/lit/internal/domain/post"
)

type createPost interface {
	Draft(ctx context.Context, draftPostCommand post.DraftPostCommand) (domain.Post, error)
	Publish(ctx context.Context, publishPostCommand post.PublishPostCommand) (domain.Post, error)
}
