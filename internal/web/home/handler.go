package home

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/typetrait/lit/internal/domain/post"
	"github.com/typetrait/lit/internal/domain/user"
	"github.com/typetrait/lit/internal/web/rendering"
)

const (
	homeTemplate string = "home"
)

type getPosts interface {
	GetPosts(ctx context.Context) ([]post.Post, error)
}

type MockGetPosts struct {
}

func NewMockGetPosts() *MockGetPosts {
	return &MockGetPosts{}
}

func (mock *MockGetPosts) GetPosts(ctx context.Context) ([]post.Post, error) {
	author := user.User{
		ID:          0,
		Email:       "bruno@bcamargo.io",
		DisplayName: "Bruno",
		Roles:       []user.Role{},
		IsActive:    true,
		CreatedAt:   time.Date(2020, time.April, 14, 12, 0, 0, 0, time.UTC),
	}

	return []post.Post{
		{
			ID:    0,
			Title: "My first post",
			Slug:  "my-first-post",
			Content: post.Content{
				Format: post.FormatMarkdown,
				Source: "my awesome post",
			},
			Author:    author,
			CreatedAt: time.Date(2025, time.May, 2, 2, 10, 5, 0, time.UTC),
		},
		{
			ID:    1,
			Title: "Second post",
			Slug:  "second-post",
			Content: post.Content{
				Format: post.FormatMarkdown,
				Source: "yet another post out here",
			},
			Author:    author,
			CreatedAt: time.Date(2025, time.May, 3, 8, 2, 52, 0, time.UTC),
		},
	}, nil
}

type Handler struct {
	getPosts        getPosts
	contentRenderer *rendering.ContentRenderer
}

func NewHandler(getPosts getPosts, contentRenderer *rendering.ContentRenderer) *Handler {
	return &Handler{
		getPosts:        getPosts,
		contentRenderer: contentRenderer,
	}
}

func (h *Handler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		posts, err := h.getPosts.GetPosts(c.Request().Context())
		if err != nil {
			return err
		}

		var postViewModels []PostPreviewViewModel
		for _, p := range posts {
			renderedHTML, err := h.contentRenderer.ContentToHTML(p.Content)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, "failed to render post content")
			}

			postViewModels = append(postViewModels, PostPreviewViewModel{
				Title:       p.Title,
				HTMLContent: renderedHTML,
				Slug:        p.Slug,
			})
		}

		return c.Render(http.StatusOK, homeTemplate, ViewModel{
			Posts: postViewModels,
		})
	}
}
