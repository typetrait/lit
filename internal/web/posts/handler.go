package posts

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/typetrait/lit/internal/domain/post"
	"github.com/typetrait/lit/internal/web/rendering"
)

const (
	postTemplate  string = "post"
	postsTemplate string = "posts"
)

type getPost interface {
	GetPostBySlug(ctx context.Context, slug string) (post.Post, error)
}

type getPosts interface {
	GetPosts(ctx context.Context) ([]post.Post, error)
}

type Handler struct {
	getPost         getPost
	getPosts        getPosts
	contentRenderer *rendering.ContentRenderer
}

func NewHandler(getPost getPost, getPosts getPosts, contentRenderer *rendering.ContentRenderer) *Handler {
	return &Handler{
		getPost:         getPost,
		getPosts:        getPosts,
		contentRenderer: contentRenderer,
	}
}

func (h *Handler) View() echo.HandlerFunc {
	return func(c echo.Context) error {
		slug := c.Param("slug")
		if slug == "" {
			return c.HTML(http.StatusBadRequest, "missing slug parameter")
		}

		p, err := h.getPost.GetPostBySlug(c.Request().Context(), slug)
		if err != nil {
			// TODO: render 404 placeholder page
			return c.HTML(http.StatusNotFound, "post not found")
		}

		renderedHTML, err := h.contentRenderer.ContentToHTML(p.Content)
		if err != nil {
			return c.HTML(http.StatusInternalServerError, "failed to render post content")
		}

		postViewModel := PostViewModel{
			Title:       p.Title,
			HTMLContent: renderedHTML,
			Slug:        p.Slug,
			DatePosted:  p.CreatedAt.String(),
		}

		return c.Render(http.StatusOK, postTemplate, postViewModel)
	}
}

func (h *Handler) List() echo.HandlerFunc {
	return func(c echo.Context) error {
		posts, err := h.getPosts.GetPosts(c.Request().Context())
		if err != nil {
			return c.HTML(http.StatusInternalServerError, "failed to list posts")
		}

		h.sortPostsByDate(posts)

		var postViewModels []PostViewModel
		for _, p := range posts {
			renderedHTML, err := h.contentRenderer.ContentToHTML(p.Content)
			if err != nil {
				return c.HTML(http.StatusInternalServerError, "failed to render post content")
			}

			postViewModels = append(postViewModels, PostViewModel{
				Title:       p.Title,
				HTMLContent: renderedHTML,
				Slug:        p.Slug,
				DatePosted:  p.CreatedAt.Format(time.DateOnly),
			})
		}

		postsViewModel := ViewModel{
			Posts: postViewModels,
		}
		return c.Render(http.StatusOK, postsTemplate, postsViewModel)
	}
}

func (h *Handler) sortPostsByDate(posts []post.Post) {
	slices.SortFunc(posts, func(a, b post.Post) int {
		if a.CreatedAt.Equal(b.CreatedAt) {
			return 0
		}
		if a.CreatedAt.After(b.CreatedAt) {
			return -1
		}
		return 1
	})
}
