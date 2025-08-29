package media

import (
	"io"

	"github.com/typetrait/lit/internal/domain/post"
)

type GetMediaQuery struct {
	PostID  int64
	MediaID int64
}

func NewGetMediaQuery(postID int64, mediaID int64) GetMediaQuery {
	return GetMediaQuery{
		PostID:  postID,
		MediaID: mediaID,
	}
}

type GetMediaQueryResult struct {
	Media   post.Media
	Content io.Reader
}

func NewGetMediaQueryResult(media post.Media, content io.Reader) GetMediaQueryResult {
	return GetMediaQueryResult{
		Media:   media,
		Content: content,
	}
}
