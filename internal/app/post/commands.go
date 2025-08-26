package post

import (
	"github.com/typetrait/lit/internal/domain/user"
)

const (
	contentFormatMarkdown = "markdown"
)

type DraftPostCommand struct {
	Author user.User
}

type PublishPostCommand struct {
	ID            int64
	Title         string
	ContentFormat string
	Content       string
	Author        user.User
}
