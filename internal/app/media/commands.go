package media

import "io"

type Upload struct {
	PostID int64
	Reader io.ReadCloser
}

type UploadMediaCommand struct {
	upload Upload
}

func NewUploadMediaCommand(upload Upload) UploadMediaCommand {
	return UploadMediaCommand{
		upload: upload,
	}
}
