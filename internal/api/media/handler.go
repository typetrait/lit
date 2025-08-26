package media

import (
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/typetrait/lit/internal/app/media"
	domain "github.com/typetrait/lit/internal/domain/media"
)

const (
	MultipartFileFieldName = "file"
)

type uploadMedia interface {
	Upload(ctx context.Context, command media.UploadMediaCommand) (domain.Media, error)
}

type APIHandler struct {
	uploadMedia uploadMedia
}

func NewAPIHandler(uploadMedia uploadMedia) *APIHandler {
	return &APIHandler{
		uploadMedia: uploadMedia,
	}
}

func (h *APIHandler) Upload() echo.HandlerFunc {
	return func(c echo.Context) error {
		postIDParam := c.Param("id")
		if postIDParam == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "missing post id")
		}

		postID, err := strconv.ParseInt(postIDParam, 10, 64)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid post id")
		}

		file, ok := c.Request().MultipartForm.File[MultipartFileFieldName]
		if !ok {
			return echo.NewHTTPError(http.StatusBadRequest, "multipart file not found")
		}

		reader, err := file[0].Open()
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "failed to open file")
		}

		upload := media.Upload{
			PostID: postID,
			Reader: reader,
		}

		ctx := c.Request().Context()
		cmd := media.NewUploadMediaCommand(upload)
		uploadedMedia, err := h.uploadMedia.Upload(ctx, cmd)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, uploadMediaResponse{
			ID:     uploadedMedia.ID,
			PostID: postID,
		})
	}
}
