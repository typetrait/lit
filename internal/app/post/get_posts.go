package post

import (
	"context"
	"fmt"

	"github.com/typetrait/lit/internal/domain/post"
)

type GetPosts struct {
	postRepository Repository
}

func NewGetPosts(postRepository Repository) *GetPosts {
	return &GetPosts{
		postRepository: postRepository,
	}
}

func (gp *GetPosts) GetPosts(ctx context.Context) ([]post.Post, error) {
	posts, err := gp.postRepository.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting posts: %w", err)
	}
	return posts, nil
}
