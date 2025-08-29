package media

import (
	"io"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/typetrait/lit/internal/app/media"
)

const (
	MultipartFileFieldName    = "file"
	MultipartAltFieldName     = "alt"
	MultipartCaptionFieldName = "caption"
)

type APIHandler struct {
	getMedia    getMedia
	uploadMedia uploadMedia
}

func NewAPIHandler(getMedia getMedia, uploadMedia uploadMedia) *APIHandler {
	return &APIHandler{
		getMedia:    getMedia,
		uploadMedia: uploadMedia,
	}
}

func (h *APIHandler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		postIDParam := c.Param("post_id")
		if postIDParam == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "missing post id")
		}

		postID, err := strconv.ParseInt(postIDParam, 10, 64)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid post id")
		}

		mediaIDParam := c.Param("media_id")
		if mediaIDParam == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "missing media id")
		}
		mediaID, err := strconv.ParseInt(mediaIDParam, 10, 64)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid media id")
		}

		ctx := c.Request().Context()
		result, err := h.getMedia.Get(ctx, media.NewGetMediaQuery(postID, mediaID))
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		contentBytes, err := io.ReadAll(result.Content)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.Blob(http.StatusOK, result.Media.Mime, contentBytes)
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

		err = c.Request().ParseMultipartForm(32 << 20)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "could not parse multipart form")
		}

		if c.Request().MultipartForm == nil {
			return echo.NewHTTPError(http.StatusBadRequest, "request must be valid multipart")
		}

		altField, ok := c.Request().MultipartForm.Value[MultipartAltFieldName]
		if !ok {
			return echo.NewHTTPError(http.StatusBadRequest, "missing multipart alt field")
		}
		alt := altField[0]

		var caption *string
		captionField, ok := c.Request().MultipartForm.Value[MultipartCaptionFieldName]
		if ok {
			caption = &captionField[0]
		}

		file, ok := c.Request().MultipartForm.File[MultipartFileFieldName]
		if !ok {
			return echo.NewHTTPError(http.StatusBadRequest, "no file to upload")
		}

		fileReader, err := file[0].Open()
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "failed to open file")
		}

		upload := media.NewUpload(postID, fileReader)

		ctx := c.Request().Context()
		cmd := media.NewUploadMediaCommand(upload, alt, caption)
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
