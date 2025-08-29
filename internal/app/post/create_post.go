package post

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/typetrait/lit/internal/domain/post"
)

const (
	UntitledDraft = "Untitled Draft"
)

type CreatePost struct {
	postRepository Repository
}

func NewCreatePost(postRepository Repository) *CreatePost {
	return &CreatePost{
		postRepository: postRepository,
	}
}

func (cp *CreatePost) Draft(ctx context.Context, draftPostCommand DraftPostCommand) (post.Post, error) {
	title := fmt.Sprintf("%s %s", UntitledDraft, time.Now().Format("20060102150405.000000000"))
	draftToCreate := post.Post{
		Title: title,
		Slug: cp.defaultSlugStrategy(
			title,
		),
		Content: post.Content{
			Format: post.FormatMarkdown,
			Source: "",
		},
		Status:    post.StatusDraft,
		Author:    draftPostCommand.Author,
		CreatedAt: time.Now(),
	}

	createdDraft, err := cp.postRepository.Create(ctx, draftToCreate)
	if err != nil {
		return post.Post{}, errors.Join(ErrPostCreationFailed, err)
	}

	return createdDraft, nil
}

func (cp *CreatePost) Publish(ctx context.Context, publishPostCommand PublishPostCommand) (post.Post, error) {
	existingPost, err := cp.postRepository.FindByID(ctx, publishPostCommand.ID)
	if err != nil {
		return post.Post{}, fmt.Errorf("finding draft post: %w", err)
	}

	contentType := post.FormatMarkdown
	switch strings.ToLower(publishPostCommand.ContentFormat) {
	case contentFormatMarkdown:
		contentType = post.FormatMarkdown
	default:
		return post.Post{}, post.ErrInvalidContentFormat
	}

	existingPost.Title = publishPostCommand.Title
	existingPost.Slug = cp.defaultSlugStrategy(publishPostCommand.Title)
	existingPost.Content = post.Content{
		Format: contentType,
		Source: publishPostCommand.Content,
	}
	existingPost.Status = post.StatusPublished
	existingPost.Author = publishPostCommand.Author

	publishedPost, err := cp.postRepository.Update(ctx, existingPost)
	if err != nil {
		return post.Post{}, errors.Join(ErrPostUpdateFailed, err)
	}

	return publishedPost, nil
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
