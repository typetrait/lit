package post

import (
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/typetrait/lit/internal/app/post"
	domain "github.com/typetrait/lit/internal/domain/post"
	"github.com/typetrait/lit/internal/domain/user"
)

type createPost interface {
	Draft(ctx context.Context, draftPostCommand post.DraftPostCommand) (domain.Post, error)
	Publish(ctx context.Context, publishPostCommand post.PublishPostCommand) (domain.Post, error)
}

type APIHandler struct {
	createPost createPost
}

func NewAPIHandler(createPost createPost) *APIHandler {
	return &APIHandler{
		createPost: createPost,
	}
}

func (h *APIHandler) Draft() echo.HandlerFunc {
	return func(c echo.Context) error {
		draftPostCommand := post.DraftPostCommand{
			Author: user.User{
				ID: 1,
			}, // TODO: extract from session
		}

		ctx := c.Request().Context()
		draft, _ := h.createPost.Draft(ctx, draftPostCommand)

		return c.JSON(http.StatusCreated, DraftPostResponse{
			PostID: draft.ID,
		})
	}
}

func (h *APIHandler) Publish() echo.HandlerFunc {
	return func(c echo.Context) error {
		postIDParam := c.Param("id")
		if postIDParam == "" {
			return c.JSON(http.StatusBadRequest, "invalid id")
		}

		postID, err := strconv.ParseInt(postIDParam, 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "invalid post id")
		}

		var publishPostRequest PublishPostRequest
		if err := c.Bind(&publishPostRequest); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		publishPostCommand := post.PublishPostCommand{
			ID:            postID,
			Title:         publishPostRequest.Title,
			ContentFormat: publishPostRequest.ContentFormat,
			Content:       publishPostRequest.Content,
			Author: user.User{
				ID: 1,
			}, // TODO: extract from session
		}

		publishedPost, err := h.createPost.Publish(c.Request().Context(), publishPostCommand)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, PublishPostResponse{
			PostID: publishedPost.ID,
			Title:  publishedPost.Title,
			Slug:   publishedPost.Slug,
		})
	}
}
