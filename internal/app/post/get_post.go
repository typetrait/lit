package post

import (
	"context"
	"errors"

	"github.com/typetrait/lit/internal/domain/post"
)

type GetPost struct {
	postRepository Repository
}

func NewGetPost(postRepository Repository) *GetPost {
	return &GetPost{
		postRepository: postRepository,
	}
}

func (gp *GetPost) GetPostBySlug(ctx context.Context, slug string) (post.Post, error) {
	p, err := gp.postRepository.FindBySlug(ctx, slug)
	if err != nil {
		if errors.Is(err, ErrPostNotFound) {
			return post.Post{}, ErrPostNotFound
		}
	}
	return p, nil
}
