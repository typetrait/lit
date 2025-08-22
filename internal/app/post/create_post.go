package post

import (
	"context"
	"errors"
	"regexp"
	"strings"
	"time"

	"github.com/typetrait/lit/internal/domain/post"
)

type CreatePost struct {
	postRepository Repository
}

func NewCreatePost(postRepository Repository) *CreatePost {
	return &CreatePost{
		postRepository: postRepository,
	}
}

func (cp *CreatePost) Create(ctx context.Context, createPostCommand CreatePostCommand) (post.Post, error) {
	postToCreate := post.Post{
		Title: createPostCommand.Title,
		Slug:  cp.defaultSlugStrategy(createPostCommand.Title),
		Content: post.Content{
			Format: post.FormatMarkdown,
			Source: createPostCommand.Content,
		},
		Author:    createPostCommand.Author,
		CreatedAt: time.Now(),
	}

	createdPost, err := cp.postRepository.Create(ctx, postToCreate)
	if err != nil {
		return post.Post{}, errors.Join(ErrPostCreationFailed, err)
	}

	return createdPost, nil
}

func (cp *CreatePost) defaultSlugStrategy(title string) string {
	return strings.ReplaceAll(
		cp.removeSpecialCharacters(
			strings.ToLower(title),
		),
		" ",
		"-",
	)
}

func (cp *CreatePost) removeSpecialCharacters(input string) string {
	re, _ := regexp.Compile("[^a-zA-Z0-9 ]+")
	return re.ReplaceAllString(input, "")
}
