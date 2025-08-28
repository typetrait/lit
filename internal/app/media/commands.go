package media

import "io"

type Upload struct {
	PostID int64
	Reader io.ReadCloser
}

func NewUpload(postID int64, reader io.ReadCloser) Upload {
	return Upload{
		PostID: postID,
		Reader: reader,
	}
}

type UploadMediaCommand struct {
	Upload  Upload
	Alt     string
	Caption *string
}

func NewUploadMediaCommand(upload Upload, alt string, caption *string) UploadMediaCommand {
	return UploadMediaCommand{
		Upload:  upload,
		Alt:     alt,
		Caption: caption,
	}
}
